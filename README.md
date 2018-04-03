# heim-deps

This repository contains production dependencies for Euphoria. It's intended to
serve as an audit log for third-party software used by our servers.

## Setting up deps for development

Heim's docker-compose.yml is set up to use deps stored in `deps/godeps` and
`deps/node_modules`. This repository is included in the heim repo as a git
submodule named `_deps` (prefixed with an underscore so `go test` will not
traverse it). To ensure that the submodule is up to date, run:

    git submodule update --init

## Adding new deps / updating

To fetch deps, run:

    ./deps.sh update ./path/to/heim/repo

This will download golang and nodejs deps into the repository. When the process
is completed, the versions of go and node/npm will be printed along with the
date. It's helpful to include these lines in your commit message when checking
in new deps.

## Large files

Due to GitHub restrictions, a few large binaries are stored in a chunked form.
To restore them to their original condition (e.g. before running tests), run

    node_modules/.bin/frankenstein stitch

and to expand them into the chunked form, run

    node_modules/.bin/frankenstein rip

Depending on your disk speed, you might wish to add `--verbose` switches,
which will enable status reports.

**Note** that overly large files MUST NOT by present in ANY commit; otherwise
uploading to GitHub will FAIL.
