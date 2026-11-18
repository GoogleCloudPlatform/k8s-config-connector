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

"""Unit tests for dev/tasks/gather_release_notes.py.

Run: python3 -B dev/tasks/gather_release_notes_test.py -v
"""

import io
import os
import sys
import tempfile
import unittest
import urllib.error

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

import gather_release_notes as grn


class _FakeHTTPResponse:
    def __init__(self, body: str, status: int = 200):
        self._data = body.encode("utf-8")
        self.status = status

    def getcode(self) -> int:
        return self.status

    def read(self, amt: int | None = None) -> bytes:
        return self._data if amt is None else self._data[:amt]

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc, tb):
        return False


class GatherReleaseNotesTest(unittest.TestCase):

    def test_parse_and_match_blocked_kinds(self):
        script = """
        # BLOCKED_KINDS=(CommentedOutKind)
        BLOCKED_KINDS=(
          FooAlpha
          "BarBeta" # trailing comment
        )
        """
        kinds = grn.parse_blocked_kinds(script)
        self.assertEqual(kinds, ["FooAlpha", "BarBeta"])
        matcher = grn.make_blocked_kinds_matcher(kinds)
        self.assertEqual(matcher("Added `FooAlpha` and `barbeta-sample`."), ["BarBeta", "FooAlpha"])
        self.assertEqual(matcher("Added `FooAlphaExtra`."), [])

    def test_parse_resource_doc_slugs(self):
        ls_tree = "\n".join([
            "scripts/generate-google3-docs/resource-reference/generated/resource-docs/sql/sqlinstance.md",
            "scripts/generate-google3-docs/resource-reference/generated/resource-docs/compute/computeurlmap.md",
            "scripts/generate-google3-docs/resource-reference/_toc.yaml",
        ])
        self.assertEqual(
            grn.parse_resource_doc_slugs(ls_tree),
            {"sql/sqlinstance", "compute/computeurlmap"},
        )

    def test_extract_urls_from_markdown_handles_parens_code_and_comments(self):
        md = """<!--
* [`IgnoredInComment`](https://cloud.google.com/config-connector/docs/reference/resource-docs/foo/bar)
-->
* [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
* Distrusts compromised CAs ([golang/go#82345](https://github.com/golang/go/issues/82345)).
* Supports [Wiki](https://en.wikipedia.org/wiki/Kubernetes_(software)) and `https://%s-aiplatform.googleapis.com`.
* Bare link (see https://cloud.google.com/composer/docs).
```yaml
url: https://ignored-inside-fenced-code.example.com/404
```
[ref-id]: https://docs.cloud.google.com/logging/docs
"""
        extracted = grn.extract_urls_from_markdown(md)
        self.assertEqual(
            extracted,
            [
                (
                    4,
                    "https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance",
                ),
                (5, "https://github.com/golang/go/issues/82345"),
                (6, "https://en.wikipedia.org/wiki/Kubernetes_(software)"),
                (7, "https://cloud.google.com/composer/docs"),
                (11, "https://docs.cloud.google.com/logging/docs"),
            ],
        )

    def test_check_resource_doc_url(self):
        valid_slugs = {"sql/sqlinstance", "firebasehosting/firebasehostingsite_beta"}

        is_res, err = grn.check_resource_doc_url(
            "https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance#spec",
            valid_slugs,
        )
        self.assertTrue(is_res)
        self.assertIsNone(err)

        is_res, err = grn.check_resource_doc_url(
            "https://docs.cloud.google.com/config-connector/docs/reference/resource-docs/discoveryengine/discoveryenginedatastore",
            valid_slugs,
        )
        self.assertTrue(is_res)
        self.assertIsNotNone(err)
        self.assertIn("discoveryengine/discoveryenginedatastore.md", err)

        is_res, err = grn.check_resource_doc_url(
            "https://cloud.google.com/vertex-ai/docs/training/create-custom-job",
            valid_slugs,
        )
        self.assertFalse(is_res)
        self.assertIsNone(err)

    def test_detect_soft_404(self):
        # cloud.google.com soft-404 structure (script contains dummy <title> before real 404 title)
        cloud_soft_404 = """
        <html><head>
        <script>const tpl = '<title></title>';</script>
        <title>404. Page Not Found</title>
        </head><body><h1>Sorry, we can't find that page</h1></body></html>
        """
        self.assertIsNotNone(grn.detect_soft_404(cloud_soft_404))

        # Real documentation page mentioning 404 / notFoundPage in content must NOT trigger false positive
        valid_doc = """
        <html><head>
        <title>Troubleshooting 404 errors | Cloud CDN | Google Cloud Documentation</title>
        </head><body>
        <h1>Troubleshooting 404 errors</h1>
        <p>Configure notFoundPage for 404 Page Not Found responses.</p>
        </body></html>
        """
        self.assertIsNone(grn.detect_soft_404(valid_doc))

    def test_check_remote_url_handles_http_errors_and_soft_404(self):
        def fake_urlopen(req, timeout=15.0):
            url = req.full_url
            if "missing" in url:
                raise urllib.error.HTTPError(url, 404, "Not Found", hdrs=None, fp=None)
            if "soft404" in url:
                return _FakeHTTPResponse(
                    "<html><head><title>404. Page Not Found</title></head></html>", 200
                )
            return _FakeHTTPResponse(
                "<html><head><title>Cloud Composer overview | Google Cloud</title></head></html>",
                200,
            )

        self.assertIsNone(
            grn.check_remote_url("https://cloud.google.com/composer/docs#section", urlopen=fake_urlopen)
        )
        err_404 = grn.check_remote_url(
            "https://cloud.google.com/vertex-ai/docs/missing", urlopen=fake_urlopen
        )
        self.assertIn("HTTP 404", err_404)

        err_soft = grn.check_remote_url(
            "https://cloud.google.com/config-delivery/soft404", urlopen=fake_urlopen
        )
        self.assertIn("soft-404", err_soft)

    def test_check_notes_file_catches_blocked_kinds_alpha_resource_docs_and_bad_urls(self):
        matcher = grn.make_blocked_kinds_matcher(["SecretBlockedKind"])
        valid_slugs = {"sql/sqlinstance", "compute/computeurlmap"}

        def fake_checker(url: str) -> str | None:
            if "create-custom-jobs" in url:
                return "HTTP 404 (Not Found)"
            if "config-delivery/docs" in url:
                return "HTTP 200 soft-404 (title: '404. Page Not Found')"
            return None

        with tempfile.TemporaryDirectory() as tmpdir:
            bad_path = os.path.join(tmpdir, "release-1.161.md")
            with open(bad_path, "w", encoding="utf-8") as f:
                f.write(
                    "\n".join([
                        "* `AIPlatformCustomJob`: [Custom jobs](https://cloud.google.com/vertex-ai/docs/training/create-custom-jobs)",
                        "* `ConfigDeliveryRelease`: [Config Delivery](https://cloud.google.com/config-delivery/docs)",
                        "* [`DiscoveryEngineDataStore`](https://cloud.google.com/config-connector/docs/reference/resource-docs/discoveryengine/discoveryenginedatastore)",
                        "* [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)",
                        "* Added `SecretBlockedKind`.",
                    ])
                )

            stderr_buf = io.StringIO()
            orig_stderr = sys.stderr
            try:
                sys.stderr = stderr_buf
                rc = grn.check_notes_file(
                    bad_path,
                    matcher,
                    1,
                    "v1.161.0",
                    valid_resource_doc_slugs=valid_slugs,
                    check_remote_urls=True,
                    url_checker=fake_checker,
                )
            finally:
                sys.stderr = orig_stderr

            self.assertEqual(rc, 1)
            err_out = stderr_buf.getvalue()
            self.assertIn("SecretBlockedKind", err_out)
            self.assertIn("discoveryengine/discoveryenginedatastore", err_out)
            self.assertIn("create-custom-jobs", err_out)
            self.assertIn("config-delivery/docs", err_out)
            self.assertNotIn("sql/sqlinstance", err_out)

            good_path = os.path.join(tmpdir, "release-clean.md")
            with open(good_path, "w", encoding="utf-8") as f:
                f.write(
                    "\n".join([
                        "* `AIPlatformCustomJob`: [Custom jobs](https://cloud.google.com/vertex-ai/docs/training/create-custom-job)",
                        "* `DiscoveryEngineDataStore`",
                        "* [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)",
                    ])
                )

            rc_good = grn.check_notes_file(
                good_path,
                matcher,
                1,
                "v1.161.0",
                valid_resource_doc_slugs=valid_slugs,
                check_remote_urls=True,
                url_checker=fake_checker,
            )
            self.assertEqual(rc_good, 0)


if __name__ == "__main__":
    unittest.main()
