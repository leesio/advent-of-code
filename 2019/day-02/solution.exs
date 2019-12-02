defmodule DayTwo do
  def get_input do
    case File.read("input") do
      {:error, error} -> {:error, error}
      {:ok, body} ->
        {
          :ok,
          String.trim_trailing(body) |>
          String.split(",") |>
          Enum.map(fn (x) -> String.to_integer(x) end)
        }
    end
  end

  def traverse_list(list, head) do
    chunk_size = 4
    case Enum.slice(list, head, chunk_size) do
      [99 | _] ->
        list
      [opcode, a, b, index] ->
        val = case opcode do
          1 -> Enum.at(list, a) + Enum.at(list, b)
          2 -> Enum.at(list, a) * Enum.at(list, b)
          _ -> throw "unrecognised opcode #{opcode}"
        end
        list = List.replace_at(list, index, val)
        head = head + chunk_size
        traverse_list(list, head)
    end
  end


  def part_one do
    case get_input() do
      {:error, error} ->
        IO.puts("error retrieving input: #{error}")
      {:ok, input} ->
        input = List.replace_at(input, 1, 12)
        input = List.replace_at(input, 2, 2)
        traverse_list(input, 0)
    end
  end
  def part_two do
    case get_input() do
      {:error, error} ->
        IO.puts("error retrieving input: #{error}")
      {:ok, clean_input} ->
        Enum.reduce_while(1..100, 0, fn noun, _acc ->
          res = Enum.reduce_while(1..100, 0, fn verb, _acc ->
            input = List.replace_at(clean_input, 1, noun)
            input = List.replace_at(input, 2, verb)
            case traverse_list(input, 0) do
              [19690720 | _] -> {:halt, (100 * noun + verb)}
              _ -> {:cont, 0}
            end
          end)
          case res do
            0 -> {:cont, 0}
            _ -> {:halt, res}
          end
        end)
    end
  end
end

DayTwo.part_one |> Enum.at(0) |> IO.puts
DayTwo.part_two |> IO.puts

