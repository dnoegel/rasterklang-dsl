; a declared length alongside `each` is legal and checked against the
; derived one (8 tokens * 1/16 = 2 beats).
tempo 120 meter 4/4
pattern hats 2 each 1/16 = "c6 d6 e6 f6 g6 a6 b6 c7"
instrument lead
  wave pulse
  adsr 0 0 15 0
end
voice 0 lead $hats
