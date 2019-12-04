defmodule DayFour do
  def get_input do
    {246515,739105}
  end

  def get_digits(n) do
    Integer.to_string(n) |>
    String.split("") |>
    Enum.filter(fn x -> x !== "" end)
  end

  def max_two_adjacent(digits) do
    digits = Enum.map(digits, &String.to_integer/1)
    acc = [[-1]]

    Enum.reduce(digits, acc, fn (n, acc) ->
      [head_list | tail_list] = acc
      [head_val | _] = head_list
      if head_val == n do
        [[n | head_list] | tail_list]
      else
        [[n] | acc]
      end
    end) |>
    Enum.filter(fn l -> length(l) === 2 end) |>
    length > 0
  end

  def two_adjacent(digits) do
    Enum.reduce(digits, fn (curr, prev) ->
      cond do
        prev === -1 -> -1
        curr === prev -> -1
        true -> curr
      end
    end) === -1
  end

  def all_increase(digits) do
    Enum.reduce(digits, fn (curr, prev) ->
      cond do
        prev === -1 -> -1
        curr < prev -> -1
        true -> curr
      end
    end) > -1
  end

  def is_valid(v) do
    all_increase(v) && two_adjacent(v)
  end

  def part_one do
    {a, b} = get_input()
    Enum.filter(a..b, fn n ->
      digits = get_digits(n)
      all_increase(digits) && two_adjacent(digits)
    end) |> Enum.count
  end

  def part_two do
    {a, b} = get_input()
    Enum.filter(a..b, fn n ->
      digits = get_digits(n)
      all_increase(digits) &&  max_two_adjacent(digits)
    end) |> Enum.count
  end

end

DayFour.part_one |> IO.puts
DayFour.part_two |> IO.puts

