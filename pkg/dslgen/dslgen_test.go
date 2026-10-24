package dslgen

import (
	"strings"
	"testing"
)

// TestFoldRepeatingGroups is an equivalence oracle against the JavaScript
// importer's foldRepeatingGroups(): expected strings were captured by
// running the JS function on each case, not hand-derived. "The exact
// sequence that broke the first version" guards the equal-weight-chunk
// safety requirement both implementations must preserve.
func TestFoldRepeatingGroups(t *testing.T) {
	cases := []struct {
		label string
		flat  string
		want  string
	}{
		{"a bar repeated six times, then a variation twice",
			"e2 e2 e3 e2 e2 e2 e3 e2 e2 e2 e3 e2 e2 e2 e3 e2 e2 e2 e3 e2 e2 e2 e3 e2 e2 d2 d3 d2 e2 d2 d3 d2",
			"[e2 e2 e3 e2]!6 [e2 d2 d3 d2]!2"},
		{"weighted figure repeated, then a weighted variation",
			"a3@2 a3 a3@2 a3 a3 a3 a3@2 a3 a3@2 a3 a3 a3 a3@2 a3 a3@2 a3 a3 a3 g3@2 g3 g3@2 g3 g3 g3",
			"[a3@2 a3 a3@2 a3 a3 a3]!3 [g3@2 g3 g3@2 g3 g3 g3]"},
		{"the exact sequence that broke the first version",
			"g3@2 g3 g3@2 g3 g3 g3",
			"g3@2 g3 g3@2 g3!3"},
		{"two halves, each a repeated figure",
			"a3 a3 a4 a3 a3 a3 a4 a3 g3 g3 g4 g3 g3 g3 g4 g3",
			"[a3 a3 a4 a3]!2 [g3 g3 g4 g3]!2"},
		{"one long figure played twice",
			"a4 e5 c5 e5 g5 e5 d5 b4 a4 e5 c5 e5 g5 e5 d5 b4",
			"[a4 e5 c5 e5 g5 e5 d5 b4]!2"},
		{"no repetition at all (mixed weights)",
			"d5@3 e5@3 a4@8 b4@2 c5@4 a4@2 g4@8 e4@2",
			"d5@3 e5@3 a4@8 b4@2 c5@4 a4@2 g4@8 e4@2"},
		{"adjacent identical weighted atoms",
			"a3@4 d4@2 e4@4 e4@4 c4@2 a3@4 g3@2 a3@8 g3@2",
			"a3@4 d4@2 e4@4!2 c4@2 a3@4 g3@2 a3@8 g3@2"},
		{"single token", "c4", "c4"},
		{"all identical, weight 1", "c4 c4 c4 c4", "c4!4"},
		{"all identical, weighted", "c4@2 c4@2 c4@2 c4@2", "c4@2!4"},
		{"rests mixed in", "~ c4 ~ c4 ~ c4 ~ c4", "[~ c4]!4"},
		{"instrument names with weights",
			"Kick Bass@3 Snare Bass@3 Kick Bass@3 Snare Bass@3",
			"[Kick Bass@3 Snare Bass@3]!2"},
	}
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			got := strings.Join(FoldRepeatingGroups(strings.Split(c.flat, " ")), " ")
			if got != c.want {
				t.Errorf("FoldRepeatingGroups(%q) = %q, want %q", c.flat, got, c.want)
			}
		})
	}
}

// TestFormatNumericTable is an equivalence oracle against the JavaScript
// importer's formatNumericTable(), same discipline as above: expected
// strings are what the JS function produces.
func TestFormatNumericTable(t *testing.T) {
	floorSweep := func() []int {
		var out []int
		v := 1088
		for i := 0; i < 128; i++ {
			out = append(out, v)
			v -= 48
			if v < 0 {
				v = 0
			}
		}
		return out
	}()
	ceilSweep := func() []int {
		var out []int
		v := 2175
		for i := 0; i < 68; i++ {
			out = append(out, v)
			v += 127
			if v > 4095 {
				v = 4095
			}
		}
		return out
	}()
	const104x2048 := func() []int {
		out := make([]int, 104)
		for i := range out {
			out[i] = 2048
		}
		return out
	}()

	cases := []struct {
		label  string
		values []int
		want   string
	}{
		{"single value", []int{2048}, "2048"},
		{"short constant run (below the ramp threshold)", []int{0, 0, 0}, "0!3"},
		{"short run then a different value", []int{0, 0, 0, 1}, "0!3 1"},
		{"the gate-table shape that regressed", []int{1, 1, 1, 0, 0}, "1!3 0 0"},
		{"the arp shape that regressed", []int{0, 3, 3, 7, 7}, "0 3 3 7 7"},
		{"long constant run", const104x2048, "2048!104"},
		{"pure ascending ramp", []int{10, 20, 30, 40, 50}, "ramp 10..50 step 10 over 5"},
		{"pure descending ramp", []int{50, 40, 30, 20, 10}, "ramp 50..10 step -10 over 5"},
		{"ramp saturating at a floor", floorSweep, "ramp 1088..0 step -48 over 128"},
		{"ramp saturating at a ceiling", ceilSweep, "ramp 2175..4095 step 127 over 68"},
		{"alternating values (no run at all)", []int{0, 1, 0, 1, 0, 1}, "0 1 0 1 0 1"},
	}
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			got := FormatNumericTable(c.values)
			if got != c.want {
				t.Errorf("FormatNumericTable(%v) = %q, want %q", c.values, got, c.want)
			}
		})
	}
}

// TestFoldPeriodicRuns exercises the phase-alignment correctness property
// foldPeriodicRuns/FoldPeriodicRuns documents: a period-2 alternating run
// that begins mid-phase must fold to the unit that, when queried at the
// run's OWN absolute start index mod period, reproduces the run's first
// element - not simply "whatever order the run's first two elements happen
// to appear in".
func TestFoldPeriodicRuns(t *testing.T) {
	// A run of "b a b a b a" starting at absolute index 7 (odd) must fold to
	// "<a b>", because index 7 selects unit element 7%2=1, which must be "b"
	// (the run's own first element) -- so unit[0]="a", unit[1]="b".
	prefix := []string{"x", "x", "x", "x", "x", "x", "x"} // indices 0..6
	run := []string{"b", "a", "b", "a", "b", "a"}         // indices 7..12
	slots := append(append([]string{}, prefix...), run...)

	got := FoldPeriodicRuns(slots)
	joined := strings.Join(got, " ")
	if !strings.Contains(joined, "<a b>!6") {
		t.Fatalf("FoldPeriodicRuns(%v) = %q, want a unit containing \"<a b>!6\" (phase-locked to absolute index 7)", slots, joined)
	}

	// Sanity: the fold must be semantically equivalent -- re-expanding the
	// nested alternation at absolute cycle indices 7..12 (i.e. element
	// (cycleIndex-7) of the "<a b>!6" run, sampled as unit[cycleIndex%2])
	// must reproduce the original run exactly.
	unit := []string{"a", "b"}
	for i, want := range run {
		cycleIndex := 7 + i
		got := unit[cycleIndex%2]
		if got != want {
			t.Fatalf("phase check failed at cycle %d: unit[%d%%2]=%q, want %q", cycleIndex, cycleIndex, got, want)
		}
	}
}

func TestAtomWeight(t *testing.T) {
	cases := map[string]float64{"c4": 1, "c4@3": 3, "c4@2.5": 2.5}
	for atom, want := range cases {
		if got := AtomWeight(atom); got != want {
			t.Errorf("AtomWeight(%q) = %v, want %v", atom, got, want)
		}
	}
}

func TestCollapseRepeats(t *testing.T) {
	got := strings.Join(CollapseRepeats([]string{"0", "0", "0", "2", "2", "0", "0", "0", "0"}), " ")
	want := "0!3 2!2 0!4"
	if got != want {
		t.Errorf("CollapseRepeats = %q, want %q", got, want)
	}
	// weight-suffixed tokens are never folded, even when repeated.
	got2 := strings.Join(CollapseRepeats([]string{"c4@2", "c4@2", "c4@2"}), " ")
	want2 := "c4@2 c4@2 c4@2"
	if got2 != want2 {
		t.Errorf("CollapseRepeats(weighted) = %q, want %q", got2, want2)
	}
}

func TestFoldSingleRuns(t *testing.T) {
	got := strings.Join(FoldSingleRuns([]string{"a", "a", "a", "b", "c", "c"}), " ")
	want := "a!3 b c!2"
	if got != want {
		t.Errorf("FoldSingleRuns = %q, want %q", got, want)
	}
}

func TestRenderAlternation(t *testing.T) {
	if got := RenderAlternation([]string{"x", "x", "x"}); got != "x" {
		t.Errorf("all-identical slots = %q, want bare %q", got, "x")
	}
	// RenderAlternation always wraps its FoldPeriodicRuns result in an outer
	// "<...>" (this IS the top-level alternation, distinct from the nested
	// "<a b>!6" FoldPeriodicRuns itself produces for a periodic run).
	got := RenderAlternation([]string{"a", "b", "a", "b", "a", "b"})
	if got != "<<a b>!6>" {
		t.Errorf("RenderAlternation alternating = %q, want %q", got, "<<a b>!6>")
	}
}
