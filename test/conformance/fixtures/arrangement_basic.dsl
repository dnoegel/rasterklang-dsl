; ADR 0003 D4: an arrangement sequences patterns of different declared
; lengths in beat time - fixes F3 (a hard error in 3.0.0).
tempo 120 meter 4/4
pattern a 4 = "c4 d4 e4 f4"
pattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"
voice 0 $a $b
