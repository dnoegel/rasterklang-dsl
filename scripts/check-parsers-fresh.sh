#!/usr/bin/env bash
# Staleness check for the checked-in generated parsers (src/dsl/generated/,
# pkg/dsl/generated/): regenerates them into a scratch directory and diffs
# against what's actually committed -- same "committed generated file must
# match its generator" pattern as test/dsl/generatedSpec.test.js, but the
# output needs Java+ANTLR to reproduce, so it's an opt-in script rather than
# part of `npm test`/`go test` (which should not require Java).
#
# Run this after any change to grammar/RasterklangDsl.g4 (and in CI, if
# available) to catch an edited .g4 whose parsers were not regenerated.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SCRATCH="$(mktemp -d)"
trap 'rm -rf "$SCRATCH"' EXIT

cp -r "$ROOT" "$SCRATCH/rasterklang-dsl"
(cd "$SCRATCH/rasterklang-dsl" && ANTLR_JAR="${ANTLR_JAR:-/tmp/antlr/antlr4.jar}" bash scripts/generate-parsers.sh >/dev/null)

if diff -rq "$ROOT/src/dsl/generated" "$SCRATCH/rasterklang-dsl/src/dsl/generated" \
   && diff -rq "$ROOT/pkg/dsl/generated" "$SCRATCH/rasterklang-dsl/pkg/dsl/generated"; then
  echo "OK: checked-in generated parsers match grammar/RasterklangDsl.g4"
  exit 0
fi

echo "STALE: committed src/dsl/generated/ and/or pkg/dsl/generated/ do not match" >&2
echo "what grammar/RasterklangDsl.g4 currently generates. Run" >&2
echo "  bash scripts/generate-parsers.sh" >&2
echo "and commit the result." >&2
exit 1
