package main

import (
	"fmt"
	"net"

	"github.com/valyala/fasthttp"
)

func main() {
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		default:
			ip := ctx.Request.Header.Peek("X-Forwarded-For")
			if len(ip) == 0 {
				switch addr := ctx.RemoteAddr().(type) {
				case *net.UDPAddr:
					ip = []byte(addr.IP.String())
				case *net.TCPAddr:
					ip = []byte(addr.IP.String())
				}
			}
			ctx.Success("text", []byte(ip))
		}
	}

	ln, err := net.Listen("tcp", ":http")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := fasthttp.Serve(ln.(*net.TCPListener), m); err != nil {
		fmt.Println(err)
	}
}
