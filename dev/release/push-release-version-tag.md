This is the instruction for propose a new release for config connector.

Please use the VERSION, GIT_COMMIT and REPO_PATH environment variables. If not set, find the correct new version that should be created from `git tag`. Find the git commit hash that corresponds to the propose release PR that was merged into the release branch, which is `release-${VERSION}`. The repo path should be the full path to the root repo of kubernetes config connector.

1. Please print out the value for VERSION, GIT_COMMIT and REPO_PATH variable and make sure that they are correct.
2. Run the ../tasks/push-tag-github.sh script. Check for the dry run results, if the results are good, continue to push the release version tag. Otherwise abort.