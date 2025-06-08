defmodule TpointsTest do
  use ExUnit.Case
  doctest workshop1

  test "greets the world" do
    assert workshop1.hello() == :world
  end
end
