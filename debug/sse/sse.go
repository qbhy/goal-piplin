package main

import (
	"bufio"
	"bytes"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

func handler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == "OPTIONS" {
		ctx.WriteString("ok")
		return
	}
	if bytes.HasPrefix(ctx.Path(), []byte("/source")) {
		ctx.SetContentType("text/event-stream")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type,Accept")
		ctx.SetBodyStreamWriter(func(w *bufio.Writer) {
			for {
				w.Write([]byte(fmt.Sprintf("data: %s\n\n", time.Now())))

				if err := w.Flush(); err != nil {
					println("client disconnected")
					return
				}

				time.Sleep(time.Second)
			}
		})
	} else {
		ctx.SetContentType("text/html")
		ctx.SetBody([]byte(`
      <!doctype html>
      <body>
      <div id=d></div>
      <script>
        var source = new EventSource('http://localhost:8080/source');

        source.onmessage = function(e) {
          document.getElementById('d').innerHTML = e.data;
        };
      </script>
  `))
	}
}

func main() {

	server := &fasthttp.Server{
		Handler: handler,
	}

	if err := server.ListenAndServe(":8080"); err != nil {
		panic(err)
	}
}
