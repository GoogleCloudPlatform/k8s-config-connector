#!/usr/bin/env python3
# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""
gather_release_notes.py

Gathers all pull requests and their release notes between two git tags or releases.
Generates a markdown table report (or JSON / CSV) listing every PR merged
and its extracted release note (or NONE).

Kinds listed in BLOCKED_KINDS in dev/tasks/build-release-bundle (as of the ending
tag) are excluded from the release bundle and must not be release noted. The report
flags PRs that mention them, and --check-notes verifies that a drafted release notes
file does not mention any of them.
"""

import argparse
import csv
import json
import os
import re
import shutil
import signal
import subprocess
import sys

# Safely handle pipe termination (e.g. piping into head or less)
if hasattr(signal, "SIGPIPE"):
    signal.signal(signal.SIGPIPE, signal.SIG_DFL)


def get_repo_root():
    """Returns the absolute path to the root of the git repository."""
    try:
        return subprocess.check_output(
            ["git", "rev-parse", "--show-toplevel"], text=True
        ).strip()
    except subprocess.CalledProcessError:
        return os.getcwd()


def resolve_git_ref(ref: str) -> str:
    """Resolves a tag or version string to a valid git reference (e.g. 1.154.0 -> v1.154.0)."""
    # Check if exact ref exists
    res = subprocess.run(
        ["git", "rev-parse", "--verify", ref],
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )
    if res.returncode == 0:
        return ref

    # Try prepending 'v' if missing
    if not ref.startswith("v"):
        v_ref = f"v{ref}"
        res = subprocess.run(
            ["git", "rev-parse", "--verify", v_ref],
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
        )
        if res.returncode == 0:
            return v_ref

    # Try stripping 'v' if present
    if ref.startswith("v"):
        no_v_ref = ref[1:]
        res = subprocess.run(
            ["git", "rev-parse", "--verify", no_v_ref],
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
        )
        if res.returncode == 0:
            return no_v_ref

    return ref


def get_commit_sha(ref: str) -> str:
    """Returns the commit SHA for a given ref."""
    try:
        return subprocess.check_output(
            ["git", "rev-parse", "--short", ref], text=True
        ).strip()
    except subprocess.CalledProcessError:
        return ref


def extract_release_note(text: str) -> str:
    """Extracts the release note content from a PR body or commit message."""
    if not text:
        return "NONE"

    # Match code block tagged with release-note, release_note, or releasenote
    pattern = re.compile(
        r"```(?:release-note|release_note|releasenote)\s*?\r?\n(.*?)\r?\n\s*```",
        re.DOTALL | re.IGNORECASE,
    )
    match = pattern.search(text)
    if not match:
        return "NONE"

    raw_note = match.group(1).strip()

    # Strip html comments inside the note
    raw_note = re.sub(r"<!--.*?-->", "", raw_note, flags=re.DOTALL).strip()

    # Remove template instructions if present
    if "Enter your extended release note" in raw_note:
        return "NONE"

    # Normalize empty or NONE notes
    clean_upper = raw_note.upper().strip("`'\" \t\r\n.")
    if not clean_upper or clean_upper in ("NONE", "N/A", "NA", "NO", "NONE."):
        return "NONE"

    return raw_note


# Kinds listed in BLOCKED_KINDS in this script are excluded from the release bundle,
# so they must not be release noted.
RELEASE_BUNDLE_SCRIPT = "dev/tasks/build-release-bundle"


def parse_blocked_kinds(script_text: str) -> list[str]:
    """Parses the BLOCKED_KINDS bash array out of the release bundle script."""
    # Drop bash comments first so they cannot confuse the array parsing.
    uncommented = re.sub(r"(^|\s)#.*$", r"\1", script_text, flags=re.MULTILINE)
    match = re.search(r"\bBLOCKED_KINDS=\(([^)]*)\)", uncommented)
    if not match:
        return []
    return [token.strip("\"'") for token in match.group(1).split() if token.strip("\"'")]


def load_blocked_kinds(ref: str) -> list[str]:
    """Returns BLOCKED_KINDS as of the given git ref.

    Raises subprocess.CalledProcessError if the release bundle script cannot be read at that ref.
    """
    script_text = subprocess.check_output(
        ["git", "show", f"{ref}:{RELEASE_BUNDLE_SCRIPT}"],
        text=True,
        stderr=subprocess.DEVNULL,
    )
    return parse_blocked_kinds(script_text)


def make_blocked_kinds_matcher(blocked_kinds: list[str]):
    """Returns a function listing the blocked kinds that a text mentions.

    Kinds are matched case-insensitively as whole words, so for a blocked kind `FooBar`,
    both `FooBar` and `foobar-minimal` match, but `FooBarBaz` does not.
    """
    canonical = {kind.lower(): kind for kind in blocked_kinds}
    if not canonical:
        return lambda text: []
    pattern = re.compile(
        r"(?<![A-Za-z0-9])(?:" + "|".join(map(re.escape, canonical.values())) + r")(?![A-Za-z0-9])",
        re.IGNORECASE,
    )

    def find_blocked_kinds(text: str) -> list[str]:
        return sorted({canonical[m.group(0).lower()] for m in pattern.finditer(text or "")})

    return find_blocked_kinds


def check_notes_file(path: str, find_blocked_kinds, num_blocked_kinds: int, ref: str) -> int:
    """Reports each line of a drafted release notes file that mentions a blocked kind.

    Returns the process exit code: 0 if no blocked kind is mentioned, 1 otherwise.
    """
    with open(path, encoding="utf-8") as f:
        lines = f.read().splitlines()

    violations = []
    for lineno, line in enumerate(lines, start=1):
        kinds = find_blocked_kinds(line)
        if kinds:
            violations.append(f"  {path}:{lineno}: {', '.join(kinds)}: {line.strip()}")

    if not violations:
        print(f"OK: {path} mentions none of the {num_blocked_kinds} kinds blocked at '{ref}'.")
        return 0

    print(
        f"ERROR: {path} mentions kinds listed in BLOCKED_KINDS in {RELEASE_BUNDLE_SCRIPT} at '{ref}'.\n"
        "Their CRDs are excluded from the release bundle; remove these entries from the release notes:",
        file=sys.stderr,
    )
    print("\n".join(violations), file=sys.stderr)
    return 1


def fetch_commits(from_ref: str, to_ref: str) -> list[dict]:
    """Fetches first-parent commits between from_ref and to_ref."""
    git_range = f"{from_ref}..{to_ref}"
    try:
        # Format: SHA %x1f Author %x1f Subject %x1f Body %x1e
        output = subprocess.check_output(
            ["git", "log", "--first-parent", "--format=%H%x1f%an%x1f%s%x1f%b%x1e", git_range],
            text=True,
        )
    except subprocess.CalledProcessError as e:
        print(f"Error executing git log for range '{git_range}': {e}", file=sys.stderr)
        sys.exit(1)

    commits = []
    pr_num_pattern = re.compile(r"\(#(\d+)\)")

    for entry in output.split("\x1e"):
        entry = entry.strip("\n")
        if not entry:
            continue
        parts = entry.split("\x1f")
        if len(parts) < 4:
            continue

        sha, author, subject, body = parts[0], parts[1], parts[2], parts[3]

        pr_match = pr_num_pattern.search(subject)
        pr_num = int(pr_match.group(1)) if pr_match else None

        commits.append({
            "sha": sha,
            "short_sha": sha[:10],
            "author": author,
            "subject": subject,
            "body": body,
            "pr_num": pr_num,
        })

    return commits


def fetch_prs_graphql(repo: str, pr_numbers: list[int]) -> dict[int, dict]:
    """Batches GitHub GraphQL queries using the gh CLI to retrieve accurate PR details."""
    if not pr_numbers or not shutil.which("gh"):
        return {}

    owner, repo_name = repo.split("/", 1) if "/" in repo else ("GoogleCloudPlatform", repo)
    results = {}
    chunk_size = 50

    for i in range(0, len(pr_numbers), chunk_size):
        chunk = pr_numbers[i : i + chunk_size]
        subqueries = [
            f"pr_{num}: issueOrPullRequest(number: {num}) {{ "
            f"... on PullRequest {{ number title body url author {{ login }} }} }}"
            for num in chunk
        ]
        query = (
            f"query {{ repository(owner: \"{owner}\", name: \"{repo_name}\") {{ "
            f"{' '.join(subqueries)} }} }}"
        )

        try:
            res = subprocess.check_output(
                ["gh", "api", "graphql", "-f", f"query={query}"],
                text=True,
                stderr=subprocess.DEVNULL,
            )
            data = json.loads(res)
            repo_data = data.get("data", {}).get("repository", {})
            for key, pr_info in repo_data.items():
                if pr_info and "number" in pr_info:
                    results[pr_info["number"]] = {
                        "number": pr_info["number"],
                        "title": pr_info.get("title") or "",
                        "body": pr_info.get("body") or "",
                        "url": pr_info.get("url") or "",
                        "author": pr_info.get("author", {}).get("login") if pr_info.get("author") else "",
                    }
        except Exception:
            # Fall back gracefully to git commit logs if GitHub API is unreachable
            pass

    return results


def escape_markdown_table_cell(text: str) -> str:
    """Escapes pipes and normalizes newlines for Markdown table cells."""
    if not text:
        return ""
    # Replace pipe characters to prevent breaking table columns
    text = text.replace("|", "\\|")
    # Replace newlines with <br> for multi-line support inside Markdown tables
    lines = [line.strip() for line in text.splitlines() if line.strip()]
    return "<br>".join(lines)


def generate_markdown_report(
    from_ref: str,
    to_ref: str,
    from_sha: str,
    to_sha: str,
    records: list[dict],
    repo_url: str,
    num_blocked_kinds: int,
    only_notable: bool = False,
) -> str:
    """Generates a structured Markdown report with summary stats and tables."""
    notable = [r for r in records if r["release_note"] != "NONE"]
    none_records = [r for r in records if r["release_note"] == "NONE"]
    blocked = [r for r in (notable if only_notable else records) if r["blocked_kinds"]]

    lines = []
    lines.append(f"# Release Notes Report: `{from_ref}` → `{to_ref}`\n")
    lines.append(f"- **From Tag/Ref:** `{from_ref}` (`{from_sha}`)")
    lines.append(f"- **To Tag/Ref:** `{to_ref}` (`{to_sha}`)")
    lines.append(f"- **Total PRs Merged:** {len(records)}")
    lines.append(f"- **PRs with Release Notes:** {len(notable)}")
    lines.append(f"- **PRs with `NONE` Release Note:** {len(none_records)}")
    lines.append(
        f"- **PRs Mentioning Blocked Kinds:** {len(blocked)} "
        f"({num_blocked_kinds} kinds blocked at `{to_ref}`)\n"
    )

    # Listed first so that drafters see it before any release note content.
    if blocked:
        lines.append("## ⛔ PRs Mentioning Blocked Kinds\n")
        lines.append(
            f"These kinds are listed in `BLOCKED_KINDS` in `{RELEASE_BUNDLE_SCRIPT}` at `{to_ref}`, "
            "so their CRDs are excluded from the release bundle. Do NOT release note them.\n"
        )
        lines.append("| PR | Title | Blocked Kinds |")
        lines.append("| :--- | :--- | :--- |")
        for r in blocked:
            pr_link = (
                f"[{r['pr_label']}]({r['url']})" if r.get("url") else r["pr_label"]
            )
            title = escape_markdown_table_cell(r["title"])
            kinds = ", ".join(f"`{k}`" for k in r["blocked_kinds"])
            lines.append(f"| {pr_link} | {title} | {kinds} |")
        lines.append("")

    # Section 1: Notable Release Notes Table
    if notable:
        lines.append("## 🌟 PRs with Release Notes\n")
        lines.append("| PR | Title | Author | Release Note |")
        lines.append("| :--- | :--- | :--- | :--- |")
        for r in notable:
            pr_link = (
                f"[{r['pr_label']}]({r['url']})" if r.get("url") else r["pr_label"]
            )
            title = escape_markdown_table_cell(r["title"])
            author = f"@{r['author']}" if r.get("author") else ""
            note = escape_markdown_table_cell(r["release_note"])
            lines.append(f"| {pr_link} | {title} | {author} | {note} |")
        lines.append("")

    if only_notable:
        return "\n".join(lines)

    # Section 2: Complete Table of all PRs
    lines.append("## 📋 Complete PR & Release Note List\n")
    lines.append("| PR | Title | Author | Release Note |")
    lines.append("| :--- | :--- | :--- | :--- |")
    for r in records:
        pr_link = (
            f"[{r['pr_label']}]({r['url']})" if r.get("url") else r["pr_label"]
        )
        title = escape_markdown_table_cell(r["title"])
        author = f"@{r['author']}" if r.get("author") else ""
        note = escape_markdown_table_cell(r["release_note"])
        lines.append(f"| {pr_link} | {title} | {author} | {note} |")

    lines.append("")
    return "\n".join(lines)


def main():
    parser = argparse.ArgumentParser(
        description="Gather all PRs and release notes between 2 git tags or releases."
    )
    parser.add_argument("from_tag", help="Starting tag, release, or git ref (e.g., v1.154.0 or 1.154.0)")
    parser.add_argument("to_tag", help="Ending tag, release, or git ref (e.g., v1.155.0 or 1.155.0)")
    parser.add_argument(
        "-o", "--output", help="Optional output file path to write report to (defaults to stdout)"
    )
    parser.add_argument(
        "--format",
        choices=["markdown", "json", "csv"],
        default="markdown",
        help="Output format (default: markdown)",
    )
    parser.add_argument(
        "--only-notable",
        action="store_true",
        help="Only output PRs that have a non-NONE release note",
    )
    parser.add_argument(
        "--offline",
        "--git-only",
        action="store_true",
        help="Only use git logs without making GitHub API requests",
    )
    parser.add_argument(
        "--repo",
        default="GoogleCloudPlatform/k8s-config-connector",
        help="GitHub repository owner/name (default: GoogleCloudPlatform/k8s-config-connector)",
    )
    parser.add_argument(
        "--check-notes",
        metavar="FILE",
        help=(
            "Instead of gathering, check that a drafted release notes file does not mention any kind "
            f"listed in BLOCKED_KINDS in {RELEASE_BUNDLE_SCRIPT} at to_tag (exits 1 if it does)"
        ),
    )

    args = parser.parse_args()

    from_ref = resolve_git_ref(args.from_tag)
    to_ref = resolve_git_ref(args.to_tag)

    # The release bundle for to_ref is built with the BLOCKED_KINDS list as of to_ref.
    try:
        blocked_kinds = load_blocked_kinds(to_ref)
    except subprocess.CalledProcessError:
        message = f"could not read {RELEASE_BUNDLE_SCRIPT} at '{to_ref}'"
        if args.check_notes:
            sys.exit(f"Error: {message}; cannot check the release notes for blocked kinds.")
        print(f"Warning: {message}; PRs mentioning blocked kinds will not be flagged.", file=sys.stderr)
        blocked_kinds = []
    find_blocked_kinds = make_blocked_kinds_matcher(blocked_kinds)

    if args.check_notes:
        sys.exit(check_notes_file(args.check_notes, find_blocked_kinds, len(blocked_kinds), to_ref))

    from_sha = get_commit_sha(from_ref)
    to_sha = get_commit_sha(to_ref)

    commits = fetch_commits(from_ref, to_ref)
    if not commits:
        print(f"No commits found between '{from_ref}' and '{to_ref}'.", file=sys.stderr)
        sys.exit(0)

    # Fetch PR details from GitHub API if available and not offline
    pr_nums = [c["pr_num"] for c in commits if c["pr_num"] is not None]
    gh_prs = {}
    if not args.offline and shutil.which("gh"):
        gh_prs = fetch_prs_graphql(args.repo, pr_nums)

    repo_base_url = f"https://github.com/{args.repo}"

    records = []
    for c in commits:
        pr_num = c["pr_num"]
        gh_info = gh_prs.get(pr_num) if pr_num else None

        title = gh_info["title"] if gh_info and gh_info.get("title") else c["subject"]
        body = gh_info["body"] if gh_info and gh_info.get("body") else c["body"]
        author = gh_info["author"] if gh_info and gh_info.get("author") else c["author"]
        pr_url = (
            gh_info["url"]
            if gh_info and gh_info.get("url")
            else (f"{repo_base_url}/pull/{pr_num}" if pr_num else "")
        )
        pr_label = f"#{pr_num}" if pr_num else c["short_sha"]

        release_note = extract_release_note(body)

        records.append({
            "pr_number": pr_num,
            "pr_label": pr_label,
            "sha": c["sha"],
            "short_sha": c["short_sha"],
            "title": title,
            "author": author,
            "url": pr_url,
            "release_note": release_note,
            "blocked_kinds": find_blocked_kinds(f"{title}\n{release_note}"),
        })

    # Output formatting
    if args.format == "json":
        output_content = json.dumps(
            {
                "from_ref": from_ref,
                "to_ref": to_ref,
                "from_sha": from_sha,
                "to_sha": to_sha,
                "total_prs": len(records),
                "blocked_kinds": blocked_kinds,
                "prs": [r for r in records if not args.only_notable or r["release_note"] != "NONE"],
            },
            indent=2,
        )
    elif args.format == "csv":
        import io

        out_io = io.StringIO()
        writer = csv.writer(out_io)
        writer.writerow(["PR", "Title", "Author", "URL", "Release Note", "Blocked Kinds"])
        for r in records:
            if args.only_notable and r["release_note"] == "NONE":
                continue
            writer.writerow([
                r["pr_label"], r["title"], r["author"], r["url"], r["release_note"],
                ", ".join(r["blocked_kinds"]),
            ])
        output_content = out_io.getvalue()
    else:
        output_content = generate_markdown_report(
            from_ref=from_ref,
            to_ref=to_ref,
            from_sha=from_sha,
            to_sha=to_sha,
            records=records,
            repo_url=repo_base_url,
            num_blocked_kinds=len(blocked_kinds),
            only_notable=args.only_notable,
        )

    if args.output:
        with open(args.output, "w", encoding="utf-8") as f:
            f.write(output_content)
        print(f"Report written to {args.output}")
    else:
        print(output_content)


if __name__ == "__main__":
    main()
