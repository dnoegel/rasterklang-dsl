table sweep = ramp 1088..0 step -48 over 128

instrument lead
  pw $sweep 2048!104
  filterTable ramp 0..40 over 8
end
