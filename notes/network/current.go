package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://yandex.ru")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range resp.Header {
		fmt.Println(k, "\t\t", v)
	}

	fmt.Println(resp.Proto)
	fmt.Println(resp.TLS.Version, resp.TLS.NegotiatedProtocol)

	// b, _ := httputil.DumpResponse(resp, false)
	// fmt.Println(string(b))
}
