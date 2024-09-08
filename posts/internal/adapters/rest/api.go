package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
	"github.com/fmo/hexagonal-blog/internal/ports"
	"log"
	"net/http"
	"strconv"
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
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var post domain.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		savedPost, err := a.api.SavePost(ctx, post)
		if err != nil {
			log.Fatalf("cant save")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Message: savedPost,
			Status:  200,
		})
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		postIdStr := query.Get("postId")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		post, _ := a.api.GetPost(ctx, int64(postId))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Message: post,
			Status:  200,
		})
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), nil)
	if err != nil {
		log.Panic(err)
	}
}
