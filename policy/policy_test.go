package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 83, Capacity: 78, Latency: 16, Risk: 14, Weight: 8}, wantScore: 174, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 75, Capacity: 72, Latency: 13, Risk: 9, Weight: 13}, wantScore: 216, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 73, Capacity: 99, Latency: 9, Risk: 17, Weight: 8}, wantScore: 181, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
