package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
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

type Response struct {
	Message domain.Post `json:"message"`
	Status  int         `json:"status"`
}

func (a Adapter) Run(ctx context.Context) {

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		_, err := a.api.SavePost(ctx, domain.NewPost("title", "body"))
		if err != nil {
			log.Fatalf("cant save")
		}
	})

	http.HandleFunc("/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		post, _ := a.api.GetPost(ctx, 1)

		response := Response{
			Message: post,
			Status:  200,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), nil)
	if err != nil {
		log.Panic(err)
	}
}
