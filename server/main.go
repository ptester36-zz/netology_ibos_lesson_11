package main

import (
	"log"
	"net/http"
	"os"
)

type Handler struct {

}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("cat /etc/passwd"))
}

func main() {

	h := Handler{}

	s := http.Server{
		Addr: ":9999",
		Handler: h,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}