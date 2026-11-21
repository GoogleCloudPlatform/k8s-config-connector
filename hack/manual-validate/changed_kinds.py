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

"""Change scout: which Kinds changed on master in a window, and which are worth verifying.

Reads git only (no network, no GCP). For every first-parent commit in the window it maps the
changed files to Kinds and classifies each Kind's change:

  new        the Kind's direct controller first appeared in the window
  fields     its *_types.go gained, lost or changed fields (listed by JSON name)
  behaviour  its controller, hand-written mapper, identity or reference code changed
  types      no controller serves the Kind yet (e.g. a prepopulate step), or its default
             controller is not the direct one, so the changed direct code does not run: not selected
  mechanical every change came from a commit touching many Kinds at once (a refactor, a bulk
             regeneration): not selected, except a sample per such commit
  tests      only its fixtures, fuzzer or tests changed: not selected
  generated  only generated code changed: not selected

Usage:
  hack/manual-validate/changed_kinds.py --since "7 days ago"
  hack/manual-validate/changed_kinds.py --base <ref> [--head origin/master] --format json
  hack/manual-validate/changed_kinds.py --since "1 day ago" --format items   # fan-out item lines
  hack/manual-validate/changed_kinds.py --base <sha> --format title          # a batch issue's title
  hack/manual-validate/changed_kinds.py --base <sha> --format issue          # a batch issue's body

--format issue writes a complete fan-out parent: the spec (window, child labels), the task (one
plain issue per Kind, with what changed and the instructions), the items, the scout report, and a
marker with the window's commits (BATCH_MARKER) that the next batch
starts from. .agents/verify-changes.md opens it on its schedule.
"""

import argparse
import json
import os
import re
import subprocess
import sys
from collections import OrderedDict

SELECTED_CLASSES = ("new", "fields", "behaviour")
CLASS_ORDER = ("new", "fields", "behaviour", "types", "mechanical", "tests", "generated")

DIRECT = "pkg/controller/direct/"
FIXTURES = "pkg/test/resourcefixture/testdata/"
RELEASE_BUNDLE = "dev/tasks/build-release-bundle"
STATIC_CONFIG = "pkg/controller/resourceconfig/static_config.go"
STATIC_RE = re.compile(r'Kind:\s*"(\w+)"\}:\s*\{DefaultController:\s*k8s\.ReconcilerType(\w+),\s*'
                       r'SupportedControllers:\s*\[\]k8s\.ReconcilerType\{([^}]*)\}')
SCRIPTS = "hack/manual-validate/"

PR_RE = re.compile(r"\(#(\d+)\)\s*$")
STRUCT_RE = re.compile(r"^type\s+(\w+)\s+struct\b", re.M)
FIELD_RE = re.compile(r'^([+-])\s+(\w+)\s+[^/\n]*?`[^`]*json:"([^",]+)')


def git(*args, repo="."):
    return subprocess.run(["git", "-C", repo] + list(args), check=True,
                          capture_output=True, text=True).stdout


# ---------------------------------------------------------------------------
# Index of Kinds at a commit
# ---------------------------------------------------------------------------

def kind_index(repo, ref):
    """lowercase kind -> Kind, from apis/**/<kind>_types.go files declaring `type <Kind> struct`."""
    index = {}
    files = [f for f in git("ls-tree", "-r", "--name-only", ref, "--", "apis/", repo=repo).splitlines()
             if f.endswith("_types.go")]
    for f in files:
        key = os.path.basename(f)[: -len("_types.go")]
        if key in index:
            continue
        body = git("show", f"{ref}:{f}", repo=repo)
        for name in STRUCT_RE.findall(body):
            if name.lower() == key:
                index[key] = name
                break
    return index


def blocked_kinds(repo, ref):
    try:
        body = git("show", f"{ref}:{RELEASE_BUNDLE}", repo=repo)
    except subprocess.CalledProcessError:
        return set()
    m = re.search(r"BLOCKED_KINDS=\((.*?)\)", body, re.S)
    if not m:
        return set()
    return {w.strip("\"'") for w in m.group(1).split() if not w.startswith("#")}


def default_controllers(repo, ref):
    """Kind -> default controller (direct, terraform, dcl, ...) from static_config.go at ref."""
    try:
        body = git("show", f"{ref}:{STATIC_CONFIG}", repo=repo)
    except subprocess.CalledProcessError:
        return {}
    out = {}
    for kind, default, supported in STATIC_RE.findall(body):
        default = default.lower()
        direct_opt_in = default != "direct" and "ReconcilerTypeDirect" in supported
        out[kind] = default + (" (direct opt-in)" if direct_opt_in else "")
    return out


def paths_at(repo, ref, prefix):
    return git("ls-tree", "-r", "--name-only", ref, "--", prefix, repo=repo).splitlines()


# ---------------------------------------------------------------------------
# File -> (kind key, role)
# ---------------------------------------------------------------------------

def is_generated(path):
    b = os.path.basename(path)
    return ".generated." in b or b.startswith("zz_generated") or b.endswith("_generated_test.go")


def classify_file(path, index):
    """Returns (kind_key or None, role). Roles: controller, mapper, identity, types, tests, generated."""
    base = os.path.basename(path)
    stem = base.split(".", 1)[0]
    prefix = stem.split("_", 1)[0]

    if path.startswith(FIXTURES):
        parts = path[len(FIXTURES):].split("/")
        for p in parts:
            if p in index:
                return p, "tests"
        return None, "tests"

    if path.startswith(DIRECT):
        key = prefix if prefix in index else None
        if key is None:
            parent = os.path.basename(os.path.dirname(path))
            key = parent if parent in index else None
        if key is None or not path.endswith(".go"):
            return key, "generated" if is_generated(path) else "other"
        if base.endswith("_test.go") or "_fuzzer" in base:
            return key, "tests"
        if is_generated(path):
            return key, "generated"
        if "_mapper" in base or "_mapping" in base:
            return key, "mapper"
        return key, "controller"

    if path.startswith("apis/") and path.endswith(".go"):
        key = prefix if prefix in index else None
        if key is None:
            return None, "other"
        if base.endswith("_test.go"):
            return key, "tests"
        if is_generated(path):
            return key, "generated"
        if base.endswith("_types.go"):
            return key, "types"
        if "_identity" in base or "_reference" in base:
            return key, "identity"
        return key, "other"

    return None, "other"


def has_controller(paths, key):
    for p in paths:
        k = os.path.basename(p).split(".", 1)[0].split("_", 1)[0]
        parent = os.path.basename(os.path.dirname(p))
        if (k == key and p.endswith("_controller.go")) or (parent == key and p.endswith(".go")
                                                          and not p.endswith("_test.go")):
            return True
    return False


# ---------------------------------------------------------------------------
# Commits in the window
# ---------------------------------------------------------------------------

def resolve_range(repo, base, head, since):
    head_sha = git("rev-parse", head, repo=repo).strip()
    if base:
        return git("rev-parse", base, repo=repo).strip(), head_sha
    out = git("rev-list", "-1", "--first-parent", f"--before={since}", head_sha, repo=repo).strip()
    if not out:
        sys.exit(f"no commit on {head} before {since!r}")
    return out, head_sha


def commits_in(repo, base, head):
    """[(sha, subject, [(status, path)])] oldest first, first-parent only."""
    out = git("log", "--first-parent", "--reverse", "--format=\x01%H\t%s", "--name-status",
              f"{base}..{head}", repo=repo)
    commits = []
    for chunk in out.split("\x01")[1:]:
        lines = [l for l in chunk.splitlines() if l.strip()]
        sha, _, subject = lines[0].partition("\t")
        files = []
        for l in lines[1:]:
            parts = l.split("\t")
            status = parts[0][0]
            files.append((status, parts[-1]))
        commits.append((sha, subject, files))
    return commits


TYPE_RE = re.compile(r"^type\s+(\w+)\s+struct\s*\{")
LINE_FIELD_RE = re.compile(r'^\s+(\w+)\s+[^/\n]*?`[^`]*json:"([^",]+)')


def status_fields(repo, ref, path, kind):
    """JSON names that only occur in status-like structs (or the Kind/KindList wrappers) of path at ref."""
    try:
        body = git("show", f"{ref}:{path}", repo=repo)
    except subprocess.CalledProcessError:
        return set()
    where, current = {}, None
    for line in body.splitlines():
        m = TYPE_RE.match(line)
        if m:
            current = m.group(1)
            continue
        if line.startswith("}"):
            current = None
            continue
        f = LINE_FIELD_RE.match(line)
        if f and current:
            where.setdefault(f.group(2), set()).add(current)

    def statusish(struct):
        return (struct in (kind, kind + "List") or struct.endswith("Status")
                or struct.endswith("ObservedState"))
    return {name for name, structs in where.items() if all(statusish(s) for s in structs)}


def field_changes(repo, base, head, path, kind):
    """Spec JSON field names added / removed / changed in a *_types.go between base and head."""
    try:
        diff = git("diff", "-U0", f"{base}..{head}", "--", path, repo=repo)
    except subprocess.CalledProcessError:
        return {}
    skip = status_fields(repo, head, path, kind) | status_fields(repo, base, path, kind)
    plus, minus = {}, {}
    for line in diff.splitlines():
        if line.startswith(("+++", "---")):
            continue
        m = FIELD_RE.match(line)
        if m and m.group(3) not in skip:
            (plus if m.group(1) == "+" else minus)[m.group(3)] = line[1:].strip()
    added = sorted(set(plus) - set(minus))
    removed = sorted(set(minus) - set(plus))
    changed = sorted(n for n in set(plus) & set(minus) if plus[n] != minus[n])
    return {k: v for k, v in (("added", added), ("removed", removed), ("changed", changed)) if v}


# ---------------------------------------------------------------------------
# Scout
# ---------------------------------------------------------------------------

def scout(repo, base, head, mechanical_threshold, sample):
    index = kind_index(repo, head)
    base_direct = paths_at(repo, base, DIRECT)
    head_direct = paths_at(repo, head, DIRECT)
    head_paths = set(paths_at(repo, head, FIXTURES)) | set(paths_at(repo, head, SCRIPTS))
    fixture_keys = {p for path in head_paths if path.startswith(FIXTURES)
                    for p in path[len(FIXTURES):].split("/")[:-1]}
    blocked = blocked_kinds(repo, head)
    controllers = default_controllers(repo, head)

    kinds = OrderedDict()
    mechanical_commits = []
    commits = commits_in(repo, base, head)
    for sha, subject, files in commits:
        touched = OrderedDict()
        for status, path in files:
            key, role = classify_file(path, index)
            if key is None or role == "other":
                continue
            touched.setdefault(key, []).append((status, path, role))
        code_kinds = [k for k, fs in touched.items()
                      if any(r in ("controller", "mapper", "identity", "types") for _, _, r in fs)]
        mechanical = len(code_kinds) >= mechanical_threshold
        if mechanical:
            mechanical_commits.append({"sha": sha[:10], "subject": subject, "kinds": len(code_kinds),
                                       "pr": pr_of(subject)})
        for key, fs in touched.items():
            k = kinds.setdefault(key, {"kind": index[key], "key": key, "roles": set(),
                                       "mechanical_roles": set(), "prs": [], "commits": [],
                                       "mechanical_from": [], "types_files": set(),
                                       "added_controller": False, "added_types": False})
            for status, path, role in fs:
                (k["mechanical_roles"] if mechanical else k["roles"]).add(role)
                if role == "types":
                    k["types_files"].add(path)
                if role == "controller" and status == "A":
                    k["added_controller"] = True
                if role == "types" and status == "A":
                    k["added_types"] = True
            pr = pr_of(subject)
            if mechanical:
                k["mechanical_from"].append(sha[:10])
            else:
                k["commits"].append(sha[:10])
                if pr and pr not in k["prs"]:
                    k["prs"].append(pr)

    results = []
    for key, k in kinds.items():
        roles, why, fields = k["roles"], [], {}
        is_new = k["added_controller"] and not has_controller(base_direct, key)
        controller = controllers.get(k["kind"], "none")
        controlled = (controller == "direct" or "direct opt-in" in controller) and has_controller(head_direct, key)
        if "types" in roles and not k["added_types"]:
            for f in sorted(k["types_files"]):
                for kind_of, names in field_changes(repo, base, head, f, k["kind"]).items():
                    fields.setdefault(kind_of, []).extend(names)
        if not controlled and roles & {"types", "identity", "mapper", "controller"}:
            cls = "types"
            why.append("types added" if k["added_types"] else "direct code changed")
            why.append("no controller yet" if controller == "none" else f"default controller is {controller}")
        elif is_new:
            cls = "new"
            why.append("direct controller added in the window")
        elif fields:
            cls = "fields"
            why.append("fields " + "; ".join(f"{a}: {', '.join(n)}" for a, n in fields.items()))
        elif roles & {"controller", "mapper", "identity", "types"}:
            cls = "behaviour"
            why.append("changed: " + ", ".join(sorted(roles & {"controller", "mapper", "identity", "types"})))
        elif k["mechanical_roles"] & {"controller", "mapper", "identity", "types"}:
            cls = "mechanical"
            why.append(f"only in bulk commits: {', '.join(k['mechanical_from'][:3])}")
        elif "tests" in roles | k["mechanical_roles"]:
            cls = "tests"
            why.append("fixtures/fuzzer/tests only")
        else:
            cls = "generated"
            why.append("generated code only")
        results.append({
            "kind": k["kind"], "class": cls, "why": "; ".join(why), "fields": fields,
            "prs": k["prs"], "commits": k["commits"], "mechanical_from": k["mechanical_from"],
            "has_fixture": key in fixture_keys,
            "has_script": f"{SCRIPTS}{key}/validate.sh" in head_paths,
            "released": k["kind"] not in blocked,
            "controller": controller,
            "selected": cls in SELECTED_CLASSES,
        })

    # A sample of each bulk commit's Kinds, so a refactor is not verified nowhere.
    if sample > 0:
        for mc in mechanical_commits:
            picked = 0
            for r in sorted(results, key=lambda r: r["kind"]):
                if picked >= sample:
                    break
                if r["class"] == "mechanical" and not r["selected"] and mc["sha"] in r["mechanical_from"] \
                        and r["has_fixture"]:
                    r["selected"] = True
                    r["why"] += f"; sampled from bulk commit {mc['sha']}"
                    picked += 1

    results.sort(key=lambda r: (CLASS_ORDER.index(r["class"]), r["kind"]))
    return {"base": base[:10], "head": head[:10], "base_sha": base, "head_sha": head, "commits": len(commits),
            "mechanical_commits": mechanical_commits, "kinds": results}


def pr_of(subject):
    m = PR_RE.search(subject)
    return int(m.group(1)) if m else None


# ---------------------------------------------------------------------------
# Output
# ---------------------------------------------------------------------------

def render_md(report, window):
    sel = [r for r in report["kinds"] if r["selected"]]
    rest = [r for r in report["kinds"] if not r["selected"]]
    out = [f"## Kinds changed {window} ({report['base']}..{report['head']})", "",
           f"{report['commits']} commits; {len(report['kinds'])} Kinds touched; **{len(sel)} selected** to verify.", ""]
    counts = OrderedDict((c, 0) for c in CLASS_ORDER)
    for r in report["kinds"]:
        counts[r["class"]] += 1
    out.append(" · ".join(f"{c}: {n}" for c, n in counts.items() if n))
    out += ["", "| Kind | Class | Why | PRs | Controller | Fixture | Script | In release |", "|---|---|---|---|---|---|---|---|"]
    for r in sel:
        prs = " ".join(f"#{p}" for p in r["prs"]) or "—"
        out.append(f"| {r['kind']} | {r['class']} | {r['why']} | {prs} | {r['controller']} | {yn(r['has_fixture'])} | "
                   f"{yn(r['has_script'])} | {yn(r['released'])} |")
    if report["mechanical_commits"]:
        out += ["", "Bulk commits (changes not attributed to single Kinds):", ""]
        for mc in report["mechanical_commits"]:
            pr = f" #{mc['pr']}" if mc["pr"] else ""
            out.append(f"- {mc['sha']}{pr} - {mc['kinds']} Kinds - {mc['subject']}")
    if rest:
        out += ["", "Not selected:", ""]
        for c in CLASS_ORDER:
            names = [r["kind"] for r in rest if r["class"] == c]
            if names:
                out.append(f"- {c} ({len(names)}): {', '.join(names)}")
    return "\n".join(out) + "\n"


# The item line goes into each child's task text: no quotes, backslashes or braces (template syntax).
ITEM_LINE_DROP = re.compile(r"[^A-Za-z0-9 *:;,.()#_/-]")
BATCH_LABEL = "validate-kind/batch"
# Each child's label: the validate-kind rule (.agents/watch/validate-kind.yaml) runs on it.
CHILD_LABEL = "overseer/validate-kind"
# A child whose script already covers its change makes no commit; with this marker in its body
# the factory closes it as completed (rather than leaving it open), so the fan-out goes on.
NO_CHANGE_MARKER = "<!-- factory:no-change ok -->"

# The batch's ## Task: what each child (one plain issue per Kind) does. __BASE__/__HEAD__ are the
# window; {{.item.name}} and {{.item.line}} are filled in by the fan-out for each child. No `## `
# heading inside it: the fan-out reads the parent's `## ` headings as its sections.
TASK_TEMPLATE = """Validate the {{.item.name}} controller against real GCP, the way a person would, and leave a
reusable script. Follow `.gemini/skills/manual-validate-controller/SKILL.md`; the reference script
is `hack/manual-validate/bigqueryrowaccesspolicy/validate.sh`.

**What changed** (master __BASE_SHORT__..__HEAD_SHORT__): {{.item.line}}

1. **Read the change.** `git log --oneline __BASE__..__HEAD__ -- apis pkg/controller/direct`
   and `git diff __BASE__..__HEAD__` on {{.item.name}}'s files (its `*_types.go`, controller,
   mapper, identity). Note every spec field added, removed or changed, and any changed behaviour
   (create, update, delete, how a field is sent).
2. **Write or extend the script** `hack/manual-validate/<kind lowercase>/validate.sh`:
   - It does not exist: write it from the Kind's `-maximal` (or `-full`) fixture under
     `pkg/test/resourcefixture/testdata/`, as the skill describes, including `mv_enable_apis`
     for every API it needs.
   - It exists: extend it so the change is checked - a new or changed field is set in the
     manifest the script applies (a fixture variant under
     `hack/manual-validate/<kind lowercase>/fixtures/` if the test fixture does not set it),
     compared to the value read from GCP after create, and changed in the update phase if it is
     mutable. A removed field: check the script no longer sets it and the resource still works.
   - **The tests are cumulative: add, never replace.** Keep every existing check and every
     fixture variant as it is. Change or remove one only when the API itself removed or changed
     what it checks, and say which and why in your summary.
   - Never edit test fixtures, golden files, `_http.log`, the controller or the types.
3. **If the script already covers this change** (nothing to add), make no commit.

# Review Instructions

Review `hack/manual-validate/<kind lowercase>/validate.sh` with
`.gemini/skills/manual-validate-controller/SKILL.md`: every check must compare the applied object
with a value read from GCP outside KCC (REST, gcloud/bq, audit logs), never KCC status; the change
named in the PR description must be covered; cleanup must leak-check. The tests are cumulative:
compare with the script and fixtures/ on master - no existing check or fixture variant may be
removed or weakened, unless the API itself removed or changed what it checks and the PR description
says so. The receipts in the PR description must cover the default fixture and every variant under
fixtures/. The PR must not touch the controller, the types,
or any test fixture or golden file.

__MARKER__"""
BATCH_MARKER = "<!-- manual-validate-batch base={base} head={head} -->"
BATCH_MARKER_RE = re.compile(r"<!-- manual-validate-batch base=([0-9a-f]+) head=([0-9a-f]+) -->")


def item_line(r):
    """A fan-out checklist line: the bold name is the item's key; the rest says what changed.
    It is put into each child's task text, so it keeps only plain characters (ITEM_LINE_DROP) and
    a bounded length."""
    prs = " ".join(f"#{p}" for p in r["prs"])
    text = f"**{r['kind']}** - {r['class']}: {r['why']}" + (f" ({prs})" if prs else "")
    text = ITEM_LINE_DROP.sub("", text)
    return text if len(text) <= 400 else text[:397] + "..."


def render_title(report):
    return f"Verify changed Kinds {report['base']}..{report['head']}"


def render_issue(report, window, auto_merge=True):
    sel = [r for r in report["kinds"] if r["selected"]]
    base, head = report["base_sha"], report["head_sha"]
    # overseer/auto-merge opts each child's PR in to merge-validate-kind (graded good, green).
    child_labels = ", ".join([CHILD_LABEL] + (["overseer/auto-merge"] if auto_merge else []))
    items = "\n".join(f"- [ ] {item_line(r)}" for r in sel)
    task = (TASK_TEMPLATE.replace("__BASE_SHORT__", report["base"]).replace("__HEAD_SHORT__", report["head"])
            .replace("__BASE__", base).replace("__HEAD__", head).replace("__MARKER__", NO_CHANGE_MARKER))
    # No headings inside the report: the fan-out reads the body's ## sections as the spec.
    scout = re.sub(r"^#+ (.*)$", r"**\1**", render_md(report, window), flags=re.M)
    return f"""Human-style validation, against real GCP, of every Kind whose controller changed on master in
{report['base']}..{report['head']} ({report['commits']} commits). One child per Kind runs the
plain issue (labelled `overseer/validate-kind`): its agent writes or extends `hack/manual-validate/<kind>/validate.sh`
for what changed, runs it, and reports in its PR. The batch closes once every child is closed.

Opened by `.agents/verify-changes.md` from `hack/manual-validate/changed_kinds.py`.

<details><summary>Change scout report</summary>

{scout}
</details>

{BATCH_MARKER.format(base=base, head=head)}

## Fan-out
```yaml
title: "Verify {{{{.item.name}}}} ({report['base']}..{report['head']})"
labels: [{child_labels}]
group: 1
create: lazy
window: {{start: 2, max: 3}}
checkpoints: []
```

## Task
{task}

## Items
{items}
"""


def yn(b):
    return "yes" if b else "no"


def main(argv=None):
    p = argparse.ArgumentParser(description=__doc__.split("\n")[0])
    p.add_argument("--repo", default=".")
    p.add_argument("--head", default="origin/master")
    g = p.add_mutually_exclusive_group()
    g.add_argument("--base", help="start of the window (exclusive)")
    g.add_argument("--since", default="7 days ago", help='git date, e.g. "1 day ago" or 2026-11-13')
    p.add_argument("--mechanical-threshold", type=int, default=8,
                   help="a commit changing code of this many Kinds or more is a bulk change")
    p.add_argument("--sample", type=int, default=1, help="Kinds to sample from each bulk commit")
    p.add_argument("--format", choices=("md", "json", "items", "title", "issue"), default="md")
    p.add_argument("--no-auto-merge", action="store_true",
                   help="for --format issue: do not label the children overseer/auto-merge")
    a = p.parse_args(argv)

    base, head = resolve_range(a.repo, a.base, a.head, a.since)
    report = scout(a.repo, base, head, a.mechanical_threshold, a.sample)
    window = f"since {a.base}" if a.base else f"since {a.since}"
    if a.format == "json":
        json.dump(report, sys.stdout, indent=2)
        print()
    elif a.format == "items":
        for r in report["kinds"]:
            if r["selected"]:
                print(f"- [ ] {item_line(r)}")
    elif a.format == "title":
        print(render_title(report))
    elif a.format == "issue":
        sys.stdout.write(render_issue(report, window, auto_merge=not a.no_auto_merge))
    else:
        sys.stdout.write(render_md(report, window))


if __name__ == "__main__":
    main()
