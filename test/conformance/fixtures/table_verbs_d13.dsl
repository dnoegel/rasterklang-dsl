; ADR 0003 D13: table verbs, an alternative to a bare token row.
instrument bass
  wave pulse
  pw
    set 400
    add 20 repeat 8
    add -20 repeat 8
    loop
  end
end
tempo 120 beats 4
voice 0 bass "c4"
