; Different declared lengths per voice is polymetre, which is legal (ADR
; 0001's D5). The arrangement's true period is reported, not warned about.
tempo 120 beats 4
pattern a 4 = "c2 ~ c2 ~"
pattern b 3 = "e3 f3 g3"
pattern c 8 = "c5@4 g5@4"
voice 0 $a
voice 1 $b
voice 2 $c
