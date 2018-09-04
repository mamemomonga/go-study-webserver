package webserver

import (
	"fmt"
	"log"
	"net/http"
)

type Webserver struct {
	Message string
	listen  string
}

func NewWebserver(listen string) (this *Webserver, err error) {
	err = nil
	this = new(Webserver)
	this.listen = listen
	return
}

func (this *Webserver) RunServer() (err error) {
	err = nil
	http.HandleFunc("/", this.putMessage)

	log.Printf("Start Listening at http://%s/", this.listen)
	err = http.ListenAndServe(this.listen, nil)
	return
}

func (this *Webserver) putMessage(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATH: %s", r.URL.Path)
	fmt.Fprintf(w, this.Message)
}
