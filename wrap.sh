#!/usr/bin/env bash
set -euo pipefail

CONTAINING_DIR=$(unset CDPATH && cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
readonly CONTAINING_DIR

cd "${CONTAINING_DIR}"
exec go run ./cmd "$@"
