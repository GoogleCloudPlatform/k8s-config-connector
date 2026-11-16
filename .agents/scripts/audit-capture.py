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

"""Deterministic evidence capture for a KCC greenfield controller PR.

Reads the search plan an agent wrote (audit-query.yaml), widens it so it
cannot hide anything, and writes what it found:

  - the PR's E2E fixtures and the RPCs their _http.log files record;
  - Cloud Audit Log entries in the recording project for the plan's
    window and searches, plus a search per mutating gRPC method the logs
    record (derived here, not by the agent);
  - a cross-check: each mutation the logs record against the audit
    entries for it.

It judges nothing: the analysis is the agent's, from this file. It never
writes to GCP or GitHub.

usage: audit-capture.py --query audit-query.yaml --out audit-capture.md
                        [--json audit-capture.json] [--repo .]
                        [--base refs/factory/pr/base] [--fixtures DIR ...]
"""
import argparse
import datetime as dt
import json
import os
import re
import subprocess
import sys
from collections import Counter, defaultdict

import yaml

FIXTURE_ROOT = "pkg/test/resourcefixture/testdata/"
MAX_WINDOW = dt.timedelta(days=14)
COMMIT_MARGIN = dt.timedelta(hours=1)
MAX_SEARCHES = 40
LIMIT = 1000
SAFE = re.compile(r"^[A-Za-z0-9._/\-]{3,200}$")
REQ = re.compile(r"^(GET|POST|PATCH|PUT|DELETE|GRPC) (\S+)")
HTTP_STATUS = re.compile(r"^(\d{3}) [A-Za-z]")
GRPC_ERR = re.compile(r"^error: rpc error: code = (\w+)")
CODES = ["OK", "Canceled", "Unknown", "InvalidArgument", "DeadlineExceeded", "NotFound",
         "AlreadyExists", "PermissionDenied", "ResourceExhausted", "FailedPrecondition",
         "Aborted", "OutOfRange", "Unimplemented", "Internal", "Unavailable", "DataLoss",
         "Unauthenticated"]


def ts(s):
    return dt.datetime.fromisoformat(str(s).replace("Z", "+00:00")).astimezone(dt.timezone.utc)


def iso(t):
    return t.strftime("%Y-%m-%dT%H:%M:%SZ")


def git(repo, *args):
    return subprocess.run(["git", "-C", repo, *args], check=True, capture_output=True, text=True).stdout


def fixtures_of(repo, base):
    names = git(repo, "diff", "--name-only", f"{base}...HEAD", "--", FIXTURE_ROOT).split()
    dirs = sorted({os.path.dirname(n) for n in names})
    return [d for d in dirs if os.path.isdir(os.path.join(repo, d))]


def commit_dates(repo, base, paths):
    if not paths:
        return []
    out = git(repo, "log", "--format=%aI", f"{base}..HEAD", "--", *paths).split()
    return sorted(ts(x) for x in out)


def parse_log(path):
    """The RPCs of an _http.log: phase, verb, target, outcome."""
    calls, phase, cur = [], "", None
    for line in open(path, encoding="utf-8", errors="replace"):
        line = line.rstrip("\n")
        if line.startswith("# Phase:"):
            phase = line.split(":", 1)[1].strip()
            continue
        m = REQ.match(line)
        if m and (cur is None or cur.get("outcome")):
            cur = {"phase": phase, "verb": m.group(1), "target": m.group(2), "outcome": ""}
            calls.append(cur)
            continue
        if cur is None or cur["outcome"]:
            continue
        if cur["verb"] == "GRPC":
            if line == "OK":
                cur["outcome"] = "OK"
            elif GRPC_ERR.match(line):
                cur["outcome"] = GRPC_ERR.match(line).group(1)
        else:
            m = HTTP_STATUS.match(line)
            if m:
                cur["outcome"] = m.group(1)
    return calls


def is_mutation(c):
    if c["verb"] in ("POST", "PATCH", "PUT", "DELETE"):
        return True
    if c["verb"] != "GRPC":
        return False
    method = c["target"].rsplit("/", 1)[-1]
    return not c["target"].startswith("/google.longrunning.") and not method.startswith(("Get", "List"))


def grpc_method_name(target):
    return target.strip("/").replace("/", ".")


def rest_path(url):
    """host and the last path segments of a REST call, without the query."""
    host, _, path = url.split("?", 1)[0].split("://", 1)[-1].partition("/")
    return host.split(":")[0] + " .../" + "/".join(path.split("/")[-3:])


def ok(outcome):
    return outcome in ("OK",) or (outcome.isdigit() and outcome.startswith("2"))


def query_audit(project, start, end, clause):
    flt = (f'logName:"cloudaudit.googleapis.com" AND timestamp>="{iso(start)}" '
           f'AND timestamp<="{iso(end)}" AND {clause}')
    cmd = ["gcloud", "logging", "read", flt, f"--project={project}", f"--limit={LIMIT}", "--format=json"]
    p = subprocess.run(cmd, capture_output=True, text=True)
    if p.returncode != 0:
        return flt, None, p.stderr.strip().splitlines()[-1:] or ["gcloud failed"]
    out = []
    for e in json.loads(p.stdout or "[]"):
        pp = e.get("protoPayload", {})
        st = pp.get("status") or {}
        out.append({
            "timestamp": e.get("timestamp", ""),
            "log": e.get("logName", "").rsplit("%2F", 1)[-1],
            "service": pp.get("serviceName", ""),
            "method": pp.get("methodName", ""),
            "resource": pp.get("resourceName", ""),
            "code": CODES[st.get("code", 0)] if isinstance(st.get("code", 0), int) and st.get("code", 0) < len(CODES) else str(st.get("code")),
            "message": (st.get("message") or "")[:200],
            "principal": (pp.get("authenticationInfo") or {}).get("principalEmail", ""),
        })
    return flt, out, None


def query_markers(project, start, end, test_key):
    flt = (f'logName="projects/{project}/logs/kcc-e2e-actuation-audit" AND timestamp>="{iso(start)}" '
           f'AND timestamp<="{iso(end)}" AND jsonPayload.testKey="{test_key}"')
    cmd = ["gcloud", "logging", "read", flt, f"--project={project}", "--limit=200", "--format=json"]
    p = subprocess.run(cmd, capture_output=True, text=True)
    if p.returncode != 0:
        return flt, None, p.stderr.strip().splitlines()[-1:] or ["gcloud failed"]
    out = []
    for e in json.loads(p.stdout or "[]"):
        j = e.get("jsonPayload", {})
        out.append({"timestamp": e.get("timestamp", ""), "event": j.get("event", ""), "gvk": j.get("gvk", ""),
                    "resourceName": j.get("resourceName", ""), "probedUrls": j.get("probedUrls") or []})
    return flt, out, None


def test_key(fixture_dir):
    d = fixture_dir.replace(os.sep, "/").rstrip("/")
    i = d.find(FIXTURE_ROOT)
    return d[i + len(FIXTURE_ROOT):] if i >= 0 else d


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--query", required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--json")
    ap.add_argument("--repo", default=".")
    ap.add_argument("--base", default="refs/factory/pr/base")
    ap.add_argument("--fixtures", nargs="*", help="fixture dirs (default: those the PR changes)")
    ap.add_argument("--no-git", action="store_true", help="skip git (local tests)")
    a = ap.parse_args()

    q = yaml.safe_load(open(a.query)) or {}
    notes = []
    project = str(q.get("project") or os.environ.get("GCP_PROJECT_ID") or "")
    if not SAFE.match(project):
        sys.exit(f"audit-query: project {project!r} is not a project id")

    fixtures = a.fixtures if a.fixtures else ([] if a.no_git else fixtures_of(a.repo, a.base))
    logs = {}
    for d in fixtures:
        full = d if os.path.isabs(d) else os.path.join(a.repo, d)
        files = sorted(os.listdir(full)) if os.path.isdir(full) else []
        hl = os.path.join(full, "_http.log")
        logs[d] = {"files": files, "calls": parse_log(hl) if os.path.exists(hl) else None}

    stale = None
    if not a.no_git and fixtures:
        logs_paths = [os.path.join(d, "_http.log") for d in fixtures]
        rec = commit_dates(a.repo, a.base, logs_paths)
        ctl = commit_dates(a.repo, a.base, ["pkg/controller/direct/"])
        if rec and ctl and ctl[-1] > rec[-1]:
            stale = (rec[-1], ctl[-1])

    # The window: the plan's, widened to cover every commit that touched a
    # fixture, never longer than MAX_WINDOW.
    start, end = ts(q["window"]["start"]), ts(q["window"]["end"])
    dates = [] if a.no_git else commit_dates(a.repo, a.base, fixtures)
    if dates:
        if dates[0] - COMMIT_MARGIN < start:
            notes.append(f"window start moved from {iso(start)} to cover the first fixture commit ({iso(dates[0])})")
            start = dates[0] - COMMIT_MARGIN
        if dates[-1] + COMMIT_MARGIN > end:
            notes.append(f"window end moved from {iso(end)} to cover the last fixture commit ({iso(dates[-1])})")
            end = dates[-1] + COMMIT_MARGIN
    if end <= start:
        sys.exit("audit-query: window end is not after its start")
    if end - start > MAX_WINDOW:
        notes.append(f"window cut to {MAX_WINDOW.days} days, ending {iso(end)}")
        start = end - MAX_WINDOW

    # Searches: the plan's, plus one per mutating gRPC method the logs record.
    searches = []
    for s in q.get("searches") or []:
        if not isinstance(s, dict):
            continue
        for key, field in (("method", "protoPayload.methodName"), ("resource", "protoPayload.resourceName")):
            v = str(s.get(key) or "")
            if v and SAFE.match(v):
                searches.append({"by": "plan", "why": str(s.get("why", ""))[:120], "clause": f'{field}:"{v}"', "key": f"{key}:{v}"})
            elif v:
                notes.append(f"search {key} {v!r} skipped: not a plain name")
    derived = sorted({grpc_method_name(c["target"]) for l in logs.values() if l["calls"] for c in l["calls"] if c["verb"] == "GRPC" and is_mutation(c)})
    for m in derived:
        searches.append({"by": "derived", "why": "a mutating gRPC call in an _http.log", "clause": f'protoPayload.methodName="{m}"', "key": f"method={m}"})
    seen, uniq = set(), []
    for s in searches:
        if s["clause"] not in seen:
            seen.add(s["clause"])
            uniq.append(s)
    if len(uniq) > MAX_SEARCHES:
        notes.append(f"{len(uniq) - MAX_SEARCHES} searches dropped (more than {MAX_SEARCHES})")
        uniq = uniq[:MAX_SEARCHES]

    results = []
    for s in uniq:
        flt, entries, err = query_audit(project, start, end, s["clause"])
        results.append({**s, "filter": flt, "entries": entries, "error": err})

    markers = {}
    for d in logs:
        markers[d] = query_markers(project, start, end, test_key(d))

    # Write it down.
    o = []
    o.append("# Evidence capture (deterministic)\n")
    o.append(f"- **Project:** `{project}`")
    o.append(f"- **Window:** {iso(start)} .. {iso(end)} (UTC)")
    if dates:
        o.append(f"- **Fixture commits (author dates):** {len(dates)}, first {iso(dates[0])}, last {iso(dates[-1])}")
    if stale:
        o.append(f"- **Controller changed after the last recording:** last `_http.log` commit {iso(stale[0])}, "
                 f"last `pkg/controller/direct/` commit {iso(stale[1])}")
    for n in notes:
        o.append(f"- **Note:** {n}")
    o.append("")
    o.append("## Fixtures\n")
    for d, l in logs.items():
        name = d.rstrip("/").rsplit("/", 1)[-1]
        has = lambda f: "yes" if f in l["files"] else "**no**"
        o.append(f"### `{name}`\n")
        o.append(f"- `create.yaml` {has('create.yaml')}, `update.yaml` {has('update.yaml')}, `_http.log` {has('_http.log')}, "
                 f"`_http_report.md` {has('_http_report.md')}, `_audit_probe.log` {has('_audit_probe.log')}")
        if l["calls"] is None:
            o.append("- no `_http.log`\n")
            continue
        calls = l["calls"]
        bad = [c for c in calls if c["outcome"] and not ok(c["outcome"]) and c["outcome"] != "404"]
        o.append(f"- `_http.log`: {len(calls)} calls, {sum(map(is_mutation, calls))} mutations, phases {sorted({c['phase'] for c in calls if c['phase']}) or 'none marked'}")
        o.append(f"- calls not OK (404 aside): {len(bad)}")
        o.append("\n| phase | call | outcome |\n|---|---|---|")
        for c in calls:
            if is_mutation(c) or (c["outcome"] and not ok(c["outcome"])):
                tgt = c["target"] if c["verb"] == "GRPC" else c["target"].split("?", 1)[0]
                o.append(f"| {c['phase'] or '-'} | `{c['verb']} {tgt[-110:]}` | {c['outcome'] or '?'} |")
        o.append("")

    o.append("## Audit log searches\n")
    for r in results:
        o.append(f"### `{r['key']}` ({r['by']}{': ' + r['why'] if r['why'] else ''})\n")
        o.append(f"`{r['filter']}`\n")
        if r["error"]:
            o.append(f"**search failed:** {r['error'][0]}\n")
            continue
        es = r["entries"]
        o.append(f"- entries: {len(es)}{' (limit reached)' if len(es) >= LIMIT else ''}")
        if not es:
            o.append("")
            continue
        o.append(f"- first {min(e['timestamp'] for e in es)}, last {max(e['timestamp'] for e in es)}")
        o.append(f"- principals: {dict(Counter(e['principal'] for e in es))}")
        o.append("\n| method | outcome | count |\n|---|---|---|")
        for (m, c), n in sorted(Counter((e["method"], e["code"]) for e in es).items()):
            o.append(f"| `{m}` | {c} | {n} |")
        msgs = Counter(e["message"] for e in es if e["message"])
        if msgs:
            o.append("\nerror messages:")
            for m, n in msgs.most_common(5):
                o.append(f"- {n}x {m}")
        res = Counter(e["resource"] for e in es)
        o.append("\nresources (most frequent first):")
        for rname, n in res.most_common(8):
            o.append(f"- {n}x `{rname}`")
        o.append("")

    o.append("## Audit probe markers (log kcc-e2e-actuation-audit)\n")
    o.append("The test harness writes one per fixture after its out-of-band probe, from the recording run itself.\n")
    for d, (flt, ms, err) in markers.items():
        o.append(f"### `{test_key(d)}`\n")
        if err:
            o.append(f"**search failed:** {err[0]}\n")
            continue
        o.append(f"- markers: {len(ms)}" + (f", first {min(m['timestamp'] for m in ms)}, last {max(m['timestamp'] for m in ms)}" if ms else ""))
        urls = sorted({u for m in ms for u in m["probedUrls"]})
        for u in urls[:6]:
            o.append(f"- probed `{u[-150:]}`")
        o.append("")

    o.append("## Cross-check: recorded gRPC mutations against audit entries\n")
    o.append("| method | in _http.log (outcomes) | in audit logs (outcomes) |\n|---|---|---|")
    by_method = defaultdict(Counter)
    for r in results:
        for e in r["entries"] or []:
            by_method[e["method"]][e["code"]] += 1
    rec = defaultdict(Counter)
    for l in logs.values():
        for c in l["calls"] or []:
            if c["verb"] == "GRPC" and is_mutation(c):
                rec[grpc_method_name(c["target"])][c["outcome"] or "?"] += 1
    for m in sorted(rec):
        audit = by_method.get(m)
        o.append(f"| `{m}` | {dict(rec[m])} | {dict(audit) if audit else '**none**'} |")
    rest = sorted({(c['verb'], rest_path(c['target'])) for l in logs.values() for c in l['calls'] or [] if c['verb'] != 'GRPC' and is_mutation(c)})
    if rest:
        o.append("\nREST mutations recorded (no fixed audit methodName: the plan's searches cover them):")
        for v, pth in rest:
            o.append(f"- `{v} {pth}`")
    o.append("")
    with open(a.out, "w") as f:
        f.write("\n".join(o))
    if a.json:
        with open(a.json, "w") as f:
            json.dump({"project": project, "window": [iso(start), iso(end)], "notes": notes,
                       "fixtures": logs, "searches": results,
                       "markers": {test_key(d): v[1] for d, v in markers.items()}}, f, indent=1)
    print(f"wrote {a.out}: {len(logs)} fixtures, {len(results)} searches, "
          f"{sum(len(r['entries'] or []) for r in results)} audit entries")


if __name__ == "__main__":
    main()
