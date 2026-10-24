; --- tempo & meter -----------------------------------------------------
tempo 120 beats 4

; --- reusable named pattern ---------------------------------------------
pattern hook = "c4!3 e4!2 g4"

; --- instruments ----------------------------------------------------------
instrument bass
  wave pulse
  adsr 0 4 10 4
  pw   1536
  vib  0.15 5
  gate 0.85
end

instrument lead
  wave triangle pulse triangle noise
  arp  0 4 7 12
  pw   2048 2400 2800 3200 2800 2400
  filterTable 0 20 40 60 40 20
  adsr 0 9 0 0
  loop 2
  slide 0.08
end

instrument perc
  wave noise
  adsr 0 2 0 2
  test 1 0 0 0
end

; --- voices ----------------------------------------------------------------
voice 0 bass
  notes "c2!3 e2!2 g1 <c2 f2>"
  |> every 4 rev
end

voice 1 lead
  notes $hook
  |> every 4 rev
  |> degradeBy 0.1
  |> sometimesBy 0.3 (ply 2)
end

voice 2 perc
  notes "x(5,8,2)"
end

; --- shared filter -----------------------------------------------------------
filter
  cutoff switch cycle
    < 3  : saw 160..0 / 3
    else : saw 2047..1440 / 9 >> 3
  resonance 15
  mode      bandpass
  route     1 2
end
