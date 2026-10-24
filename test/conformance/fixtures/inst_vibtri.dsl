; GoatTracker's own instrument vibrato (2.18.0): both SPEED spellings, and
; the note-relative `sN` one next to a plain register-unit number.
instrument lead
  wave pulse
  vibTri 5 s3 8
end

instrument bass
  wave triangle
  vibTri 16 100 0
end

voice 0 lead "c4 e4"
voice 1 "c2" vibTri:16,100,0
