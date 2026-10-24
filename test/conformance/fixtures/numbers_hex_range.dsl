voice 0 bass
  notes "c2 ~ c2 ~"
  reg control "9 8 9 8"
  reg 0x15 "3 3 3 3"
  vib 0.2 6
  slide 0.1
  gate 0.9
end

filter
  cutoff sine 300..2200 / 4
  resonance 10
  mode lowpass
  route 0 1
end
