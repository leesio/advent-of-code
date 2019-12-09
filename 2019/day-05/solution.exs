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
    IO.puts "registers"
    IO.inspect registers
    IO.puts "instruction"
    IO.inspect instruction
    registers = Enum.map(0..2, fn n ->
      case String.at(registers, n) do
        nil -> :positional
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
        IO.inspect sub_list
    # the parse argument should always return ["0", "0", "0"] even if there are
    # no registers (or :immediate, :position, or whatever)
    # Then we can zip the params with the register-type and get the values
    # (using something like `get_val` before we do any logic

    case instruction do
      "99" -> list
      _ ->
      {opcode, registers} = parse_instruction(instruction)
      chunk = get_chunk(opcode)

      {params_and_idx, _} = Enum.split(remainder, chunk)
      {params, idx} = Enum.split(params_and_idx, length(params_and_idx) - 1)
      idx = List.first(idx) |> String.to_integer()

      get_val = get_val_fn(list)
      vals = Enum.zip(params, registers)
             |> Enum.map(get_val)

      {list, ptr} = cond do
        opcode == "3" ->
          ptr = ptr + chunk + 1
          input = Task.async(fn -> IO.gets "Enter input\n" end)
          val = Task.await(input) |> String.trim_trailing
          list = List.replace_at(list, idx, to_string(val))
          {list, ptr}
        opcode == "4" ->
          ptr = ptr + chunk + 1
          val = case List.first(registers) do
            :positional -> Enum.at(list, idx)
            :immediate -> idx
          end
          IO.puts "**********"
          IO.inspect val
          IO.puts "**********"
          {list, ptr}
        opcode == "5" ->
          [val] = vals
          ptr = case val do
            "0" -> ptr + chunk + 1
            _ -> idx
          end
          {list, ptr}
        opcode == "6" ->
          [val] = vals
          ptr = case val do
            "0" -> idx
            _ -> ptr + chunk + 1
          end
          {list, ptr}
        opcode == "7" ->
          ptr = ptr + chunk + 1
          [a, b] = vals
          list = cond do
            a < b -> List.replace_at(list, idx, "1")
            true -> List.replace_at(list, idx, "0")
          end
          {list, ptr}
        opcode == "8" ->
          ptr = ptr + chunk + 1
          [a, b] = vals
          list = cond do
            a == b -> List.replace_at(list, idx, "1")
            true -> List.replace_at(list, idx, "0")
          end
          {list, ptr}
        (opcode == "1" || opcode == "2") ->
          ptr = ptr + chunk + 1
          combiner = case opcode do
            "1" -> fn {a, b} -> String.to_integer(a) + String.to_integer(b) end
            "2" -> fn {a, b} -> String.to_integer(a) * String.to_integer(b) end
          end
          val = List.to_tuple(vals)
                |> combiner.()

          list = List.replace_at(list, idx, to_string(val))
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
    example = ["3","9","7","9","10","9","4","9","99","-1","8"]
    example = ["3","3","1108","-1","8","3","4","3","99"]
    example = ["3","12","6","12","15","1","13","14","13","4","13","99","-1","0","1","9"]
    # example = ["3","3","1105","-1","9","1101","0","0","12","4","12","99","1"]
    traverse_list(example, 0)
  end
  def part_two do
  end
end


DayFive.part_one
DayFive.part_two

