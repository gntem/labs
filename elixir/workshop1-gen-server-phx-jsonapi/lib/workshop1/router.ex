defmodule Workshop1.Router do
  use Plug.Router

  plug :match
  plug :dispatch

  get "/api/hello" do
    HelloWorldServer.async_say_hello()
    json = Jason.encode!(%{message: "Async hello request received"})
    conn
    |> put_resp_content_type("application/json")
    |> send_resp(200, json)
  end

  match _ do
    send_resp(conn, 404, "Not Found")
  end
end
