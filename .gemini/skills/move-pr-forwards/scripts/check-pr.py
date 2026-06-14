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

import argparse
import json
import re
import subprocess
import sys


def run_command(cmd):
  try:
    res = subprocess.run(
        cmd, check=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True
    )
    return res.stdout
  except subprocess.CalledProcessError as e:
    print(f"Error running command: {' '.join(cmd)}", file=sys.stderr)
    print(e.stderr, file=sys.stderr)
    sys.exit(1)


def get_pr_metadata(pr, repo):
  cmd = [
      "gh",
      "pr",
      "view",
      str(pr),
      "--repo",
      repo,
      "--json",
      "title,state,mergeStateStatus,mergeable,reviewDecision,comments,latestReviews,baseRefName,headRefName",
  ]
  output = run_command(cmd)
  return json.loads(output)


def get_pr_checks(pr, repo):
  cmd = [
      "gh",
      "pr",
      "checks",
      str(pr),
      "--repo",
      repo,
      "--json",
      "name,state,bucket,link",
  ]
  try:
    # gh pr checks returns non-zero if checks are failing, so we can't use check=True directly without handling it
    res = subprocess.run(
        cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True
    )
    if not res.stdout.strip():
      return []
    return json.loads(res.stdout)
  except Exception as e:
    print(f"Error fetching PR checks: {e}", file=sys.stderr)
    return []


def get_job_log(job_id, repo):
  cmd = ["gh", "run", "view", "--log", "--job", str(job_id), "--repo", repo]
  try:
    res = subprocess.run(
        cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True
    )
    return res.stdout
  except Exception as e:
    print(f"Error fetching job log for {job_id}: {e}", file=sys.stderr)
    return ""


def main():
  parser = argparse.ArgumentParser(
      description="Check PR status, comments, and analyze failing CI logs."
  )
  parser.add_argument(
      "--pr", required=True, help="PR number or URL (e.g., 8426)"
  )
  parser.add_argument(
      "--repo",
      default="GoogleCloudPlatform/k8s-config-connector",
      help="GitHub repository",
  )
  args = parser.parse_args()

  pr_num = args.pr.split("/")[-1] if "/" in args.pr else args.pr

  print(f"=== Analyzing PR {pr_num} in {args.repo} ===\n")

  # 1. Fetch PR Metadata
  meta = get_pr_metadata(pr_num, args.repo)
  print("--- PR SUMMARY ---")
  print(f"Title:              {meta.get('title')}")
  print(f"State:              {meta.get('state')}")
  print(f"Merge State Status: {meta.get('mergeStateStatus')}")
  print(f"Mergeable:          {meta.get('mergeable')}")
  print(f"Review Decision:    {meta.get('reviewDecision')}")
  print(f"Head Branch:        {meta.get('headRefName')}")
  print(f"Base Branch:        {meta.get('baseRefName')}")
  print("------------------\n")

  # Check rebasing / conflict status
  if meta.get("mergeable") == "CONFLICTING":
    print(
        "⚠️  WARNING: PR has merge conflicts. Rebase and conflict resolution required."
    )
    print("To rebase locally:")
    print(f"  gh pr checkout {pr_num}")
    print(f"  git fetch origin {meta.get('baseRefName')}")
    print(f"  git rebase origin/{meta.get('baseRefName')}")
    print(f"  git push --force-with-lease origin {meta.get('headRefName')}\n")
  elif meta.get("mergeStateStatus") in ["BEHIND", "DIRTY"]:
    print(
        f"⚠️  NOTE: PR merge state is {meta.get('mergeStateStatus')}. A rebase on top of {meta.get('baseRefName')} may be needed."
    )
    print("To rebase locally:")
    print(f"  gh pr checkout {pr_num}")
    print(f"  git fetch origin {meta.get('baseRefName')}")
    print(f"  git rebase origin/{meta.get('baseRefName')}")
    print(f"  git push --force-with-lease origin {meta.get('headRefName')}\n")

  # 2. Analyze Comments and Reviews
  approvers = {
      "acpana",
      "barney-s",
      "cheftako",
      "fedebongio",
      "justinsb",
      "ldanielmadariaga",
      "maqiuyujoyce",
      "xiaoweim",
  }

  print("--- 🗣️  APPROVER FEEDBACK ---")
  approver_comments_found = False

  comments = meta.get("comments", [])
  reviews = meta.get("latestReviews", [])

  for c in comments:
    author = c.get("author", {}).get("login", "")
    if author in approvers:
      approver_comments_found = True
      print(f"\n[Comment by {author} at {c.get('createdAt')}]:")
      print(c.get("body", "").strip())
      print("-" * 40)

  for r in reviews:
    author = r.get("author", {}).get("login", "")
    if author in approvers:
      approver_comments_found = True
      print(
          f"\n[Review ({r.get('state')}) by {author} at {r.get('createdAt')}]:"
      )
      print(r.get("body", "").strip())
      print("-" * 40)

  if not approver_comments_found:
    print("No comments or reviews from KCC OWNERS approvers found.")
  print("-----------------------------\n")

  # Print recent non-approver comments for context
  print("--- 🤖 RECENT AUTOMATED / OTHER COMMENTS ---")
  other_comments = [
      c
      for c in comments
      if c.get("author", {}).get("login", "") not in approvers
  ]
  for c in other_comments[-3:]:  # show last 3
    author = c.get("author", {}).get("login", "")
    print(f"\n[{author} at {c.get('createdAt')}]:")
    # Truncate long automated comments if necessary
    body = c.get("body", "").strip()
    if len(body) > 300:
      body = body[:300] + " ... [truncated]"
    print(body)
    print("-" * 40)
  print("--------------------------------------------\n")

  # 3. Analyze CI Checks
  print("--- ❌ CI CHECKS & LOG ANALYSIS ---")
  checks = get_pr_checks(pr_num, args.repo)
  failing_checks = [
      c
      for c in checks
      if c.get("bucket") == "fail" or c.get("state") in ["FAILURE", "ERROR"]
  ]

  if not failing_checks:
    print("✅ No failing CI checks detected.")
  else:
    for fc in failing_checks:
      name = fc.get("name", "Unknown")
      link = fc.get("link", "")
      print(f"\nFailing Check: {name}")
      print(f"Link: {link}")

      # Extract job id from link
      job_match = re.search(r"/job/(\d+)", link)
      if not job_match:
        print(
            "Could not extract job ID from link. Unable to fetch job log directly."
        )
        continue

      job_id = job_match.group(1)
      print(f"Fetching log for Job ID {job_id}...")
      log = get_job_log(job_id, args.repo)

      if not log:
        print("⚠️ Could not retrieve log for this job.")
        continue

      lines = log.splitlines()

      # Priority 1: hint_for_agent
      hints = []
      for line in lines:
        if "<hint_for_agent>" in line and "</hint_for_agent>" in line:
          match = re.search(r"<hint_for_agent>(.*?)</hint_for_agent>", line)
          if match:
            hints.append(match.group(1).strip())
        elif "<hint_for_agent>" in line:
          hints.append(line.strip())

      if hints:
        print("\n🎯 STRONG HINT(S) FOR AGENT FOUND:")
        for h in hints:
          print(f"  👉 {h}")
      else:
        # Priority 2: FAIL:
        fails = [
            l.strip()
            for l in lines
            if "FAIL:" in l and "UNKNOWN STEP" not in l
        ]
        if not fails:
          fails = [l.strip() for l in lines if "FAIL:" in l]

        if fails:
          print("\n📝 FAILURES FOUND IN LOG:")
          for f in fails[:10]:  # limit to first 10 to avoid spam
            print(f"  - {f}")
          if len(fails) > 10:
            print(f"  ... and {len(fails) - 10} more FAIL lines.")
        else:
          # Priority 3: ERROR
          errors = [l.strip() for l in lines if "ERROR" in l]
          if errors:
            print("\n⚠️ ERRORS FOUND IN LOG:")
            for e in errors[:10]:
              print(f"  - {e}")
            if len(errors) > 10:
              print(f"  ... and {len(errors) - 10} more ERROR lines.")
          else:
            print("\n🔍 No explicit <hint_for_agent>, FAIL:, or ERROR lines found in log.")

  print("\n-----------------------------------")
  print("\n=== SUMMARY OF RECOMMENDED AGENT ACTIONS ===")
  print(
      "1. If a STRONG HINT FOR AGENT was found above, run the suggested command (e.g. 'dev/ci/presubmits/...'), verify changes, commit, and push."
  )
  print(
      "2. If KCC APPROVER FEEDBACK was found, review their comments, make the requested code changes, commit, and push."
  )
  print(
      "3. If the PR has merge conflicts or is behind/dirty, follow the rebase instructions above to rebase and force-push."
  )
  print("============================================")


if __name__ == "__main__":
  main()
