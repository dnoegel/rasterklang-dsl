; `meter N/D` (ADR 0003 D2) replaces `beats N` as the time signature: a bar
; is N beats and a beat is a 1/D note.
tempo 120 meter 7/8
pattern riff 7 = "c4 e4 g4 c5 g4 e4 c4"
voice 0 $riff wave:pulse
