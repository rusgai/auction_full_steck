package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	resve := http.Server{
		Addr:    ":8001",
		Handler: mux,
	}
	if err := resve.ListenAndServe(); err != nil {
		panic(err)
	}
}
