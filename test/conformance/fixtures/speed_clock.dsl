; `speed N` is N frames per beat - the exact, frame-native spelling of the
; clock, mutually exclusive with `tempo`.
speed 6 beats 8 offset 1
pattern riff 8 = "c4 e4 g4 c5 g4 e4 c4 ~"
voice 0 $riff wave:pulse
