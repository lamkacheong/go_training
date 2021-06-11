package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"week3/brother_server"
)

func main() {

	s1 := http.NewServeMux()
	s1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "helloworld")
	})

	s := http.Server{Addr: ":8080", Handler: s1}
	s2 := http.Server{Addr: ":8081", Handler: http.DefaultServeMux}

	serversMap := map[string]http.Server{
		"http server":  s,
		"debug server": s2,
	}

	b := brother_server.Build(&serversMap)

	err := b.Run()
	if err != nil {
		fmt.Println(err)
	}
}
