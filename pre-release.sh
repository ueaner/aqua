#!/usr/bin/env bash

git fetch origin --tags
latest_tag="$(git describe --tags --abbrev=0)"
git checkout -b "${latest_tag}-a" "${latest_tag}"
git cherry-pick 6b8b06096174c3eecb00bf5291b10622d0404966
git tag -d "${latest_tag}"
git tag "${latest_tag}"
git push --force ueaner refs/tags/"${latest_tag}"

git checkout main
git branch -D "${latest_tag}-a"
