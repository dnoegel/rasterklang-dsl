// Go half of the mini-notation LENGTH conformance suite (the JS half is
// miniLength.test.js). Both read the same mini-length.expected.json golden,
// generated from the JS implementation by generate-golden.mjs.
//
// Same argument as the parser conformance test one directory up: the beat
// arithmetic behind a declared pattern length (ADR 0001's D1) runs on both
// sides of the pipeline - JS checks the length a song declares, Go
// generators WRITE it - so the two implementations
// agreeing is a correctness property, not a nicety.
package conformance

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/dnoegel/rasterklang-dsl/pkg/dsl"
)

type miniLengthGolden struct {
	Source       string  `json:"source"`
	OK           bool    `json:"ok"`
	ContentBeats float64 `json:"contentBeats"`
	Candidates   []struct {
		Beats float64 `json:"beats"`
		What  string  `json:"what"`
	} `json:"candidates"`
	Expanded string `json:"expanded"`
}

func TestMiniLengthConformance(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(dir, "mini-length.expected.json"))
	if err != nil {
		t.Fatal(err)
	}
	var cases []miniLengthGolden
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	if len(cases) < 15 {
		t.Fatalf("expected at least 15 mini-length golden cases, found %d", len(cases))
	}

	for _, c := range cases {
		t.Run(c.Source, func(t *testing.T) {
			expanded, err := dsl.ExpandHolds(c.Source)
			if !c.OK {
				// ExpandHolds alone doesn't scan (and so can't fail) a
				// hold-free string - see its own "return source unchanged"
				// fast path - so a case invalid for a reason other than
				// hold placement (e.g. ADR 0003 D1's mixed /N-and-@n check)
				// only surfaces once something that always scans runs.
				if err == nil {
					if _, err2 := dsl.MiniContentBeats(expanded); err2 == nil {
						t.Fatalf("ExpandHolds/MiniContentBeats(%q) = %q, want an error (the JS side rejects it)", c.Source, expanded)
					}
				}
				return
			}
			if err != nil {
				t.Fatalf("ExpandHolds(%q) failed: %v", c.Source, err)
			}
			if expanded != c.Expanded {
				t.Fatalf("ExpandHolds(%q) = %q, golden %q", c.Source, expanded, c.Expanded)
			}
			beats, err := dsl.MiniContentBeats(expanded)
			if err != nil {
				t.Fatalf("MiniContentBeats(%q) failed: %v", expanded, err)
			}
			if math.Abs(beats-c.ContentBeats) > 1e-9 {
				t.Fatalf("MiniContentBeats(%q) = %v, golden %v", expanded, beats, c.ContentBeats)
			}
			got, err := dsl.MiniDeclaredLengthCandidates(expanded)
			if err != nil {
				t.Fatalf("MiniDeclaredLengthCandidates(%q) failed: %v", expanded, err)
			}
			want := make([]dsl.MiniLengthCandidate, 0, len(c.Candidates))
			for _, gc := range c.Candidates {
				want = append(want, dsl.MiniLengthCandidate{Beats: gc.Beats, What: gc.What})
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("MiniDeclaredLengthCandidates(%q) = %+v, golden %+v", expanded, got, want)
			}
		})
	}
}
