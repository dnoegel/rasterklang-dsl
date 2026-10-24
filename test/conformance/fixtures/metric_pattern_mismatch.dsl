; The declared length disagrees with the contents: a SEMANTIC error (the
; parse itself is fine, which is why this fixture's golden is a successful
; parse - see test/dsl/analyze.test.js for the diagnostic).
tempo 120 beats 4
pattern lead 32 = "c4@1.25 d4 e4@2"
voice 0 $lead
