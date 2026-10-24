; `switch bar` is 3.0.0's spelling of `switch cycle`.
tempo 120 beats 4
voice 0 "c4 e4"
filter
  cutoff switch bar
    < 4 : sine 300..1200 / 4
    else : saw 400..2000 / 8
  mode lowpass
  route 0
end
