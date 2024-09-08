package rest

import (
	"fmt"
	"github.com/fmo/hexagonal-blog/internal/ports"
	"log"
	"net/http"
)

type Adapter struct {
	api  ports.APIPorts
	port int
}

func NewAdapter(api ports.APIPorts, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		a.api.SavePost()
	})

	http.HandleFunc("/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {

	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), nil)
	if err != nil {
		log.Panic(err)
	}
}
