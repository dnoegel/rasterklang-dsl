; `*n` repeats an element, `[...]` groups and repeats a whole run - both
; flattened into the timeline's entry list at interpret time (ADR 0003 D4).
tempo 120 meter 4/4
pattern a 4 = "c4 d4 e4 f4"
pattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"
pattern verse = $a*2 $b
pattern chorus = [$a $b]*2
voice 0 lead $verse $chorus
