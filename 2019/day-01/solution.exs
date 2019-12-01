defmodule DayOne do
  def calculate_weight(i) do
      div(i, 3) - 2
  end

  def calculate_total_weight(x) do
    w = calculate_weight(x)
    if w > 0 do
      w + calculate_total_weight(w)
    else
      0
    end
  end


  def get_input do
    case File.read("input") do
      {:ok, body} ->
        {
          :ok,
          Enum.map(
            Enum.filter(String.split(body, "\n"), fn x -> x != "" end),
            fn (x) ->
              {i, _d} = Integer.parse(x)
              i
            end
          )
        }
    end
  end

  def part_one do
    case get_input() do
      {:ok, input} ->
        Enum.reduce(Enum.map(input, &calculate_weight/1), fn (x, acc) -> x + acc end)
      {:error, error} ->
        IO.puts("error retrieving input: #{error}")
    end
  end

  def part_two do
    case get_input() do
      {:ok, input} ->
        Enum.reduce(Enum.map(input, &calculate_total_weight/1), fn (x, acc) -> x + acc end)
      {:error, error} ->
        IO.puts("error retrieving input: #{error}")
    end
  end
end

IO.puts(DayOne.part_one)
IO.puts(DayOne.part_two)
