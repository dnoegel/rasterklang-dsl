speed 6 beats 4

instrument lead
  wave pulse
  adsr 0 9 6 4
  hardRestart 2 ad 0 0 sr 0 0 test
end

instrument bass
  wave saw
  hardRestart 3 sr 0 0
end

voice 0 lead "c4 e4 g4 c5"
voice 1 bass "c2 ~ g2 ~"
