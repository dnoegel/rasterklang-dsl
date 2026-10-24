filter
  cutoff    sine 300..2200 / 4
  resonance switch cycle
    < 4   : saw 4..15 / 2
    else  : 15
  mode switch cycle
    < 8   : lowpass
    else  : bandpass
  route switch cycle
    < 8   : 0 1
    else  : 1 2
end
