#!/bin/sh

set -e -x

# Creates an integer version number from the semantic version format
# May be changed when we decide to fully use semantic versions for releases
export integer_version=`cut -d "." -f1 go-url-preview-version/number`
cp -r go-url-preview promote/go-url-preview
echo ${integer_version} > promote/integer_version

pushd promote/go-url-preview
  git config --global user.email ${GITHUB_EMAIL}
  git config --global user.name ${GITHUB_USER}
  git config --global push.default simple

  echo "${integer_version}.0.0" > ci/version
  git add ci/version

  git commit -m ":airplane: New final release v${integer_version}" -m "[ci skip]"
popd
