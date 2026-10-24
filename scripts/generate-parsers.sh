#!/usr/bin/env bash
# Regenerates the JS and Go ANTLR parsers from grammar/RasterklangDsl.g4,
# the single source of truth for the grammar. Run this after any change to
# the .g4 file, then commit both the .g4 diff and the regenerated output
# under src/dsl/generated/ and pkg/dsl/generated/ (generated code is checked
# in so nobody needs Java/ANTLR installed just to build).
#
# Requires: java (17+) and an antlr4 jar. Set ANTLR_JAR to its path, or
# this defaults to /tmp/antlr/antlr4.jar.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ANTLR_JAR="${ANTLR_JAR:-/tmp/antlr/antlr4.jar}"

if [ ! -f "$ANTLR_JAR" ]; then
  echo "ANTLR jar not found at $ANTLR_JAR (set ANTLR_JAR=/path/to/antlr-4.13.2-complete.jar)" >&2
  exit 1
fi

# Run from grammar/ and pass a bare relative filename: ANTLR embeds
# whatever path it's invoked with into every generated file's "Code
# generated from ..." header comment. An absolute path there means the
# header (and nothing else) differs every time someone regenerates on a
# different machine/checkout location -- a spurious diff on every
# regeneration that has nothing to do with the grammar actually changing.
cd "$ROOT/grammar"

echo "Generating JS parser -> src/dsl/generated/"
rm -rf "$ROOT/src/dsl/generated"
mkdir -p "$ROOT/src/dsl/generated"
java -jar "$ANTLR_JAR" -Dlanguage=JavaScript -visitor -o "$ROOT/src/dsl/generated" "RasterklangDsl.g4"

echo "Generating Go parser -> pkg/dsl/generated/"
rm -rf "$ROOT/pkg/dsl/generated"
mkdir -p "$ROOT/pkg/dsl/generated"
java -jar "$ANTLR_JAR" -Dlanguage=Go -visitor -package generated -o "$ROOT/pkg/dsl/generated" "RasterklangDsl.g4"

echo "Done. Review the diff, then run the JS and Go test suites."
