defmodule HelloWorldServer do
  use GenServer

  ## Client API

  def start_link(_args \\ []) do
    GenServer.start_link(__MODULE__, :ok, name: __MODULE__)
  end

  def say_hello do
    GenServer.call(__MODULE__, :say_hello)
  end

  def async_say_hello do
    GenServer.cast(__MODULE__, :say_hello)
  end

  @impl true
  def init(:ok) do
    {:ok, %{}}
  end

  @impl true
  def handle_call(:say_hello, _from, state) do
    {:reply, "Hello, world from HelloWorldServer!", state}
  end

  @impl true
  def handle_cast(:say_hello, state) do
    IO.puts("Hello, world from HelloWorldServer!")
    {:noreply, state}
  end
end
