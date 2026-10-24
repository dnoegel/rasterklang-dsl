; Metric patterns: a declared length in beats (ADR 0001's D1) makes an
; unweighted token exactly one beat, and `@n` n beats.
tempo 120 beats 8
pattern bass 4 = "c2 ~ c2 ~"
pattern lead 8 = "c5@2 e5 g5 c6@4"
pattern hats = "bd*4"
voice 0 $bass wave:pulse
voice 1 $lead wave:triangle
voice 2 $hats wave:noise
