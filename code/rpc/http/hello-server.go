package main

import (
	"net/http"
	"net/rpc"
	"io"
	"net/rpc/jsonrpc"
)

// HelloService ...
type HelloService struct {}

// Hello ...
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(res http.ResponseWriter, req *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer: res,
			ReadCloser: req.Body,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
