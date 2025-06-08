defmodule Workshop1.Application do
  use Application

  def start(_type, _args) do
    children = [
      {HelloWorldServer, []},
      {Plug.Cowboy, scheme: :http, plug: Workshop1.Router, options: [port: 4000]}
    ]

    opts = [strategy: :one_for_one, name: Workshop1.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
