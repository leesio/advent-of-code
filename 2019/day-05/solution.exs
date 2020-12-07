defmodule DayFive do
  def get_input do
    case File.read("input") do
      {:error, error} -> {:error, error}
      {:ok, body} ->
        {
          :ok,
          String.trim_trailing(body)
          |> String.split(",")
        }
    end
  end

  def parse_instruction (instruction) do
    {opcode, registers} = String.reverse(instruction) |> String.split_at(2)
    opcode = String.reverse(opcode) |> String.trim_leading("0")

    registers =
      String.codepoints(registers)
      |> Enum.map(fn x ->
        case x do
          "0" -> :positional
          "1" -> :immediate
        end
      end)

    {opcode, registers}
  end

  def get_val_fn(list) do
    fn {param, register} ->
      case register do
        :positional -> Enum.at(list, String.to_integer(param))
        :immediate -> param
      end
    end
  end

  def get_chunk(opcode) do
    cond do
      opcode == "1" || opcode == "2" -> 3
      opcode == "3" || opcode == "4" -> 1
      opcode == "5" || opcode == "6" -> 2
      opcode == "7" || opcode == "8" -> 3
      true -> 0
    end
  end

  def traverse_list(list, ptr) do
    {_, sub_list} = Enum.split(list, ptr)
    [instruction | remainder] = sub_list

    case instruction do
      "99" -> list
      _ ->
      {opcode, registers} = parse_instruction(instruction)
      chunk = get_chunk(opcode)

      {params, _} = Enum.split(remainder, chunk)

      get_val = get_val_fn(list)
      {list, ptr} = case {opcode, params} do
        {"3", [idx]} ->
          idx = String.to_integer(idx)
          ptr = ptr + chunk + 1
          input = Task.async(fn -> IO.gets "Enter input\n" end)
          val = Task.await(input) |> String.trim_trailing
          list = List.replace_at(list, idx, to_string(val))
          {list, ptr}
        {"4", [idx]} ->
          ptr = ptr + chunk + 1
          idx = String.to_integer(idx)
          IO.puts "**********"
          IO.inspect idx
          IO.puts "**********"
          {list, ptr}
        {"5", [a, b]} ->
          ptr = case a do
            "0" -> ptr + chunk + 1
            _ -> b
          end
          {list, ptr}
        {"6", [a, b]} ->
          ptr = case a do
            "0" -> b
            _ -> ptr + chunk + 1
          end
          {list, ptr}
        {"7", [a, b, idx]} ->
          ptr = ptr + chunk + 1
          list = cond do
            a < b -> List.replace_at(list, String.to_integer(idx), "1")
            true -> List.replace_at(list, String.to_integer(idx), "0")
          end
          {list, ptr}
        {"8", [a, b, idx]} ->
          ptr = ptr + chunk + 1
          list = cond do
            a == b -> List.replace_at(list, String.to_integer(idx), "1")
            true -> List.replace_at(list, String.to_integer(idx), "0")
          end
          {list, ptr}
        {_, [a, b, idx]} ->
          ptr = ptr + chunk + 1
          combiner = case opcode do
            "1" -> fn {a, b} -> String.to_integer(a) + String.to_integer(b) end
            "2" -> fn {a, b} -> String.to_integer(a) * String.to_integer(b) end
          end
          val = combiner.({a, b})
          list = List.replace_at(list, String.to_integer(idx), to_string(val))
          {list, ptr}
        true ->
          IO.puts "unknown op code #{opcode}"
          IO.inspect opcode
      end
      traverse_list(list, ptr)
    end
  end


  def part_one do
    {:ok, input} = get_input()
    # example = ["3","9","7","9","10","9","4","9","99","-1","8"]
    # example = ["3","3","1108","-1","8","3","4","3","99"]
    example = ["3","12","6","12","15","1","13","14","13","4","13","99","-1","0","1","9"]
    # example = ["3","3","1105","-1","9","1101","0","0","12","4","12","99","1"]
    traverse_list(example, 0)
  end
  def part_two do
  end
end


DayFive.part_one
DayFive.part_two

