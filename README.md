# heim-deps

This repository contains production dependencies for Euphoria. It is intended
to serve as an audit log for third-party software used by our servers and to
preserve a known-working state of the dependencies.

## Setting up dependencies for development

Heim's `docker-compose.yml` is set up to use dependencies stored in
`_deps/godeps` and `_deps/node_modules`. This repository is included in the
Heim repository as a Git submodule named `_deps` (prefixed with an underscore
so `go test` will not traverse it). To ensure that the submodule is up to
date, run:

    git submodule update --init

## Adding new dependencies / updating

To fetch dependencies, run:

    ./deps.sh update ./path/to/heim/repo

This will download Go and Node.js dependencies into this repository. When the
process is completed, the versions of `go`, `node`, and `npm` will be printed
along with the date. It is helpful to include these lines in your commit
message when checking in new dependencies.

## Large files

Due to GitHub restrictions, a few large binaries are stored in a chunked form.
To restore them to their original condition (e.g. before running tests), run

    node_modules/.bin/frankenstein stitch

and to expand them into the chunked form, run

    node_modules/.bin/frankenstein rip

Depending on your disk speed, you might wish to add `--verbose` switches,
which will enable status reports.

**Note** that overly large files *must not* by present in *any* commit;
otherwise, uploading to GitHub will fail.
