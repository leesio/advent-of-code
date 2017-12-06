
def balance_sequence(seq)
  current_iteration = nil
  last_iteration = nil
  count = 0

  map = {}
  while true do
    if !current_iteration then
      current_iteration = seq.clone
    end

    last_iteration = current_iteration
    current_iteration = run_balance_loop(last_iteration)
    count = count + 1
    if map[current_iteration.to_s] then
      puts "part 1: #{count}"
      puts "part 2: #{count - map[current_iteration.to_s]}"
      return
    end

    map[current_iteration.to_s] = count
  end
end

def run_balance_loop(master)
  seq = master.clone
  offset = 1
  max_index = get_max_index(seq)

  counter = max_index + offset
  distributable = seq[max_index]
  seq[max_index] = 0
  while distributable > 0 do
    ptr = counter % seq.length
    seq[ptr] = seq[ptr] + 1
    distributable = distributable - 1
    counter = counter + 1
  end
  seq
end

def get_max_index(seq)
  max_index = 0
  max = 0

  seq.each.with_index do |val, v|
    if val > max then
      max = val
      max_index = v
    end
  end

  max_index
end


File.open('input').each do | line |
  seq = line.split("\t").map { |s| s.to_i }
  balance_sequence(seq)
end
