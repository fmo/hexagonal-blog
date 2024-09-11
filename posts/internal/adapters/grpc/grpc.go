package grpc

import (
	"context"
	"github.com/fmo/hexagonal-blog/golang/post"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
	log "github.com/sirupsen/logrus"
)

func (a Adapter) Create(ctx context.Context, request *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	log.WithContext(ctx).Info("Creating post...")
	newPost := domain.NewPost(request.Title, request.Body)
	result, err := a.api.SavePost(ctx, newPost)
	if err != nil {
		return nil, err
	}

	return &post.CreatePostResponse{
		PostId: result.ID,
	}, nil
}

func (a Adapter) Get(ctx context.Context, request *post.GetPostRequest) (*post.GetPostResponse, error) {
	result, err := a.api.GetPost(ctx, request.PostId)
	if err != nil {
		return nil, err
	}

	return &post.GetPostResponse{Title: result.Title, Body: result.Body}, nil
}
