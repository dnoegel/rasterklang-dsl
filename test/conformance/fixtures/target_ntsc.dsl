; ADR 0002's `target` block: the song's DRIVE rate (how many CPU cycles
; pass between two play calls), not its musical tempo. `video NTSC` is
; named sugar for the 1022727 Hz machine clock's 17045-cycle video frame.
target { video NTSC }
speed 6 beats 8
pattern riff 8 = "c4 e4 g4 c5 g4 e4 c4 ~"
voice 0 $riff wave:pulse
