// Cross-language conformance suite: this file and conformance.test.js (the JS side) both read the SAME fixtures/*.dsl
// source files and the SAME fixtures/*.expected.json golden files,
// generated once from the JS reference implementation via
// generate-golden.mjs (see that file's header comment for the mechanism).
//
// This half is what actually proves the Go port matches: canonical.go
// converts the Go AST to the same plain-JSON shape the golden was
// recorded in, and reflect.DeepEqual (via require's ObjectsAreEqual, here
// hand-rolled to avoid a new dependency) does the comparison. A
// field-name or shape mismatch between the JS and Go parsers shows up
// here as a normal failing test with a real diff, not a green run that
// only proves "the Go parser didn't crash."
package conformance

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"testing"

	"github.com/dnoegel/rasterklang-dsl/pkg/dsl"
)

type golden struct {
	OK      bool   `json:"ok"`
	AST     any    `json:"ast"`
	Line    *int   `json:"line"`
	Message string `json:"message"`
}

var errLinePattern = regexp.MustCompile(`^line (\d+):`)

func fixturesDir(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return filepath.Join(dir, "fixtures")
}

func TestConformance(t *testing.T) {
	dir := fixturesDir(t)
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	var dslFiles []string
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".dsl" {
			dslFiles = append(dslFiles, e.Name())
		}
	}
	sort.Strings(dslFiles)
	if len(dslFiles) < 15 {
		t.Fatalf("expected at least 15 fixtures, found %d", len(dslFiles))
	}

	for _, name := range dslFiles {
		name := name
		t.Run(name, func(t *testing.T) {
			src, err := os.ReadFile(filepath.Join(dir, name))
			if err != nil {
				t.Fatal(err)
			}
			goldenPath := filepath.Join(dir, name[:len(name)-len(".dsl")]+".expected.json")
			rawGolden, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatal(err)
			}
			var g golden
			if err := json.Unmarshal(rawGolden, &g); err != nil {
				t.Fatalf("invalid golden JSON: %v", err)
			}

			tokens, tokErr := dsl.Tokenize(string(src))
			var program *dsl.Program
			var parseErr error
			if tokErr == nil {
				program, parseErr = dsl.ParseProgram(tokens)
			}
			finalErr := tokErr
			if finalErr == nil {
				finalErr = parseErr
			}

			if g.OK {
				if finalErr != nil {
					t.Fatalf("expected a successful parse, got error: %v", finalErr)
				}
				got := canonProgram(program)
				// Round-trip both sides through JSON so map[string]any's
				// float64 numbers and the golden's json.Unmarshal'd
				// interface{} numbers compare identically -- avoids any
				// int/float64 mismatch between how canonical.go and
				// encoding/json each represent a plain integer.
				gotJSON, err := json.Marshal(got)
				if err != nil {
					t.Fatal(err)
				}
				var gotVal, wantVal any
				if err := json.Unmarshal(gotJSON, &gotVal); err != nil {
					t.Fatal(err)
				}
				wantJSON, err := json.Marshal(g.AST)
				if err != nil {
					t.Fatal(err)
				}
				if err := json.Unmarshal(wantJSON, &wantVal); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(gotVal, wantVal) {
					t.Fatalf("AST mismatch for %s\n--- got ---\n%s\n--- want ---\n%s", name, mustIndent(gotJSON), mustIndent(wantJSON))
				}
				return
			}

			// Error fixture: only the 1-based line number is asserted for
			// cross-language equivalence -- error MESSAGE text is not
			// required to be byte-identical between Go and JS, just present
			// and line-numbered.
			if finalErr == nil {
				t.Fatalf("expected a parse error (golden line %v), got a successful parse", g.Line)
			}
			m := errLinePattern.FindStringSubmatch(finalErr.Error())
			if m == nil {
				t.Fatalf("error message doesn't start with 'line N:': %q", finalErr.Error())
			}
			gotLine, _ := strconv.Atoi(m[1])
			if g.Line == nil {
				t.Fatalf("golden has no line number to compare against (JS message: %q)", g.Message)
			}
			if gotLine != *g.Line {
				t.Fatalf("error line mismatch for %s: got line %d, want line %d (Go message: %q, JS message: %q)", name, gotLine, *g.Line, finalErr.Error(), g.Message)
			}
		})
	}
}

func mustIndent(raw []byte) string {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return string(raw)
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return string(raw)
	}
	return string(out)
}
