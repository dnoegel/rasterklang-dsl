pattern hook = "c4!3 e4!2 g4"

voice 1 lead
  notes $hook
  |> off 0.125 (add 12)
  |> degradeBy 0.1
  |> sometimesBy 0.3 (ply 2)
end

voice 0 bass $bassLine
