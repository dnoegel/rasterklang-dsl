; `each 1/D` (ADR 0003 D3): every token is one 1/D note, length derived.
tempo 120 meter 4/4
pattern hats each 1/16 = "c6 d6 e6 f6 g6 a6 b6 c7"
instrument lead
  wave pulse
  adsr 0 0 15 0
end
voice 0 lead $hats
