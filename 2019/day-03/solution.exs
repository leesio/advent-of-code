defmodule DayThree do

  def key({x, y}) do
    String.to_atom(to_string(x) <> ":" <> to_string(y))
  end

  def get_distance({x, y}) do
    abs(x) + abs(y)
  end

  def get_travelled_distance(a, b, ptr) do
    Map.get(a, key(ptr)) + Map.get(b, key(ptr))
  end

  def get_input do
    case File.read("input") do
      {:error, error} -> {:error, error}
      {:ok, body} ->
        {
          :ok,
          String.trim_trailing(body) |>
          String.split("\n") |>
          Enum.map(fn x -> String.split(x, ",") end)
        }
    end
  end

  def plot_sub_path(initial_count, next_ptr, distance, points_touched, path) do
    Enum.reduce(
      0..distance - 1,
      {points_touched, path},
      fn (n, {points_touched, path}) ->
        [last_ptr | _] = path
        next_ptr = next_ptr.(last_ptr)
        points_touched = case Map.get(points_touched, key(next_ptr))  do
          nil -> Map.put(points_touched, key(next_ptr), initial_count + n)
          _-> points_touched
        end
          {points_touched, [next_ptr | path]}
      end)
  end

  def get_next_ptr_fn(direction) do
    case direction do
      "R" -> fn {x_ptr, y_ptr} -> {x_ptr + 1, y_ptr} end
      "L" -> fn {x_ptr, y_ptr} -> {x_ptr - 1, y_ptr} end
      "U" -> fn {x_ptr, y_ptr} -> {x_ptr, y_ptr + 1} end
      "D" -> fn {x_ptr, y_ptr} -> {x_ptr, y_ptr - 1} end
      _ -> throw("unrecognised direction")
    end
  end

  def plot_path(instructions, points_touched, path) do
    [instruction | tail] = instructions
    {direction, distance} = String.split_at(instruction, 1)
    distance = String.to_integer(distance)
    next_ptr_fn = get_next_ptr_fn(direction)

    initial_count = length(path)
    {points_touched, path} = plot_sub_path(
      initial_count,
      next_ptr_fn,
      distance,
      points_touched,
      path
    )

    case length(tail) do
      0 -> {points_touched, path}
      _-> plot_path(tail, points_touched, path)
      end
    end

  def get_intersections(path, points) do
    Enum.filter(path, fn p -> Map.get(points, key(p)) !== nil && p != {0,0} end) |>
    Enum.sort() |>
    Enum.dedup()
  end

  def part_one do
    case get_input() do
      {:error, e} ->
        IO.puts "error fetching inout #{e}"
      {:ok, [wire_one, wire_two]} ->
        {_points_one, path_one} = plot_path(wire_one, %{}, [{0, 0}])
        {points_two, _path_two} = plot_path(wire_two, %{}, [{0, 0}])
        get_intersections(path_one, points_two) |>
        Enum.map(&get_distance/1) |>
        Enum.min()
    end
  end

  def part_two do
    case get_input() do
      {:error, e} ->
        IO.puts "error fetching inout #{e}"
      {:ok, [wire_one, wire_two]} ->
        {points_one, path_one} = plot_path(wire_one, %{}, [{0, 0}])
        {points_two, _path_two} = plot_path(wire_two, %{}, [{0, 0}])
        get_intersections(path_one, points_two) |>
        Enum.map(fn x ->
          get_travelled_distance(points_two, points_one, x)
        end) |>
        Enum.min()
    end
  end

end

DayThree.part_one |> IO.puts
DayThree.part_two |> IO.puts
