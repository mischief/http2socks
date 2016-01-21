package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
	goproxy "gopkg.in/elazarl/goproxy.v1"
)

var (
	listenAddr = flag.String("listen", "127.0.0.1:8118", "http proxy listen address")
	socksAddr  = flag.String("socks", "127.0.0.1:9050", "socks5 server address")
	verbose    = flag.Bool("v", false, "verbose")
)

func main() {
	logger := log.New(os.Stderr, "http2socks: ", log.LstdFlags|log.Lshortfile)

	flag.Parse()

	prxy := goproxy.NewProxyHttpServer()
	prxy.Logger = logger
	prxy.Verbose = *verbose

	dialer, err := proxy.SOCKS5("tcp", *socksAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}

	prxy.Tr = &http.Transport{Dial: dialer.Dial}

	logger.Fatal(http.ListenAndServe(*listenAddr, prxy))
}
