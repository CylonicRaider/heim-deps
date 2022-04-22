#!/bin/bash
abspath() { cd $(dirname $1); echo $(pwd)/$(basename $1); }

usage="USAGE: $0 (update|update-go|update-js|compact-js) <heim-dir>"

get_emoji() {
  # Newer versions of twemoji no longer ship the images themselves on npm, and
  # recommend hunting for them in the GitHub repository instead.
  local version="$(egrep '^ *"version":' node_modules/twemoji/package.json | sed -r 's/.*"([^"]+)",?$/\1/')"

  rm -rf node_modules/.resources/emoji-svg tmp
  mkdir -p node_modules/.resources tmp

  git -C tmp clone --depth 1 -b gh-pages https://github.com/twitter/twemoji

  mv "tmp/twemoji/v/$version/svg" node_modules/.resources/emoji-svg

  rm -rf tmp
}

update_js_deps() {
  cp $HEIMDIR/client/package.json ./

  npm install

  get_emoji

  rm package.json
}

compact_js_deps() {
  cp $HEIMDIR/client/package.json ./

  # a few hacks to reduce footprint...

  # remove tests
  find -name test -type d -print0 | xargs -0 rm -r

  # A single global dedupe should suffice.
  #for d in node_modules/*; do pushd $d; npm dedupe; popd; done
  npm dedupe

  rm package.json package.json.original
}

print_js_versions() {
  set +x
  echo "node `node -v`; npm `npm -v`"
}

update_go_deps() {
  rm -rf godeps
  mkdir -p godeps/src/euphoria.io
  cp -r $HEIMDIR/ godeps/src/euphoria.io/heim

  GOPATH=`pwd`/godeps go get -d -t -x ./godeps/src/euphoria.io/heim/... github.com/coreos/etcd

  # pin go-etcd to v0.4.6 (because the newest version gives compile errors)
  git --git-dir=godeps/src/github.com/coreos/go-etcd/.git \
      --work-tree=godeps/src/github.com/coreos/go-etcd \
      checkout 6aa2da5

  # save the commit hash of deps but remove the git metadata so we don't create submodules.
  find godeps -depth -name .git -type d -execdir sh -c "git rev-parse HEAD > git-HEAD && rm -rf .git" \;

  # and likewise for hg repos
  find godeps -depth -name .hg -type d -execdir sh -c "hg parent -T '{node}\n' > hg-HEAD && rm -rf .hg" \;

  rm -rf godeps/src/euphoria.io/heim
}

print_go_versions() {
  set +x
  go version
}

if [[ "$1" = "" || "$2" = "" ]]; then
  echo $usage
  exit 1
fi

SRCDIR=$(dirname `abspath $0`)
HEIMDIR=$(abspath $2)

cd $SRCDIR
set -x

case $1 in
  update-go)
    update_go_deps
    print_go_versions
    date
    ;;
  update-js)
    update_js_deps
    print_js_versions
    date
    ;;
  compact-js)
    compact_js_deps
    print_js_versions
    date
    ;;
  update)
    update_go_deps
    update_js_deps
    print_go_versions
    print_js_versions
    date
    ;;
  *)
    echo $usage
    exit 1
esac
