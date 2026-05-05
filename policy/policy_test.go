package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 83, Capacity: 78, Latency: 16, Risk: 14, Weight: 8}
	if got := Score(signal); got != 174 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 75, Capacity: 72, Latency: 13, Risk: 9, Weight: 13}
	if got := Score(signal); got != 216 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 73, Capacity: 99, Latency: 9, Risk: 17, Weight: 8}
	if got := Score(signal); got != 181 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
