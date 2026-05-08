#!/bin/bash

abspath() { cd "$(dirname "$1")" && echo "$(pwd)/$(basename "$1")"; }

usage="USAGE: $0 (update|update-go|update-js|compact|print-info) <heim-dir>"

get_emoji() {
  # Newer versions of twemoji no longer ship the images themselves on npm, and
  # recommend hunting for them in the GitHub repository instead.
  local version="$(egrep '^ *"version":' node_modules/@twemoji/api/package.json | sed -r 's/.*"([^"]+)",?$/\1/')"

  rm -rf node_modules/.resources/emoji-svg tmp
  mkdir -p node_modules/.resources tmp

  git -C tmp clone --depth 1 -b gh-pages https://github.com/jdecked/twemoji

  mv "tmp/twemoji/v/$version/svg" node_modules/.resources/emoji-svg

  rm -rf tmp
}

update_js_deps() {
  cp "$HEIMDIR/client/package.json" ./
  rm -rf node_modules package-lock.json

  npm install

  get_emoji

  rm package.json
}

compact_deps() {
  # a few hacks to reduce footprint...

  # remove tests
  find -name test -type d -print0 | xargs -r0 rm -r

  # remove github cruft
  find -name .github -type d -print0 | xargs -r0 rm -r
}

print_js_versions() {
  set +x
  echo "node $(node -v); npm $(npm -v)"
}

update_go_deps() {
  rm -rf godeps
  mkdir godeps

  (cd "$HEIMDIR/server" && go mod tidy -v -x && go mod vendor -v -o "$SRCDIR/godeps")
}

print_go_versions() {
  set +x
  go version
}

print_all_versions() {
  set +x
  print_go_versions
  print_js_versions
}

if [ "$1" = "" ] || [ "$2" = "" ]; then
  echo "$usage"
  exit 1
fi

SRCDIR="$(dirname "$(abspath $0)")"
HEIMDIR="$(abspath "$2")"

cd "$SRCDIR"
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
  update)
    update_go_deps
    update_js_deps
    print_all_versions
    date
    ;;
  compact)
    compact_deps
    print_all_versions
    date
    ;;
  print-info)
    print_all_versions
    date
    ;;
  *)
    echo "$usage"
    exit 1
esac
