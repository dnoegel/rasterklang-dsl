; DSL v2 2.15.0: the per-row effect lane (`fx`) and absolute-pitch table
; cells (`=MIDI`), the two GoatTracker constructs the language previously
; had no spelling for at all.
tempo 125 beats 4

table drumPitches = =36 =43!2 0

instrument kick
  wave noise pulse!2
  arp $drumPitches
  adsr 0 8 0 0
  loop arp 3
end

instrument lead
  wave pulse
  arp 0 4 7
end

voice 0 kick "c2 ~ c2 c2"
voice 1 lead "c4 e4 g4 c5" fx:"~ tp.24 vib.16.100 ad.0.10"

voice 2
  notes "c3 ~ g3 ~"
  fx "pup.128 ~ pdn.s3 sr.4.4"
end
