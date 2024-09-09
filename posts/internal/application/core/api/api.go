package api

import (
	"context"
	"fmt"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
	"github.com/fmo/hexagonal-blog/internal/ports"
)

type Application struct {
	db    ports.DBPort
	image ports.ImagePort
}

func NewApplication(db ports.DBPort, image ports.ImagePort) *Application {
	return &Application{
		db:    db,
		image: image,
	}
}

func (a Application) SavePost(ctx context.Context, post domain.Post) (domain.Post, error) {
	err := a.db.Save(ctx, &post)
	if err != nil {
		return domain.Post{}, err
	}

	err = a.image.Upload(fmt.Sprintf("%d.png", post.ID))
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

func (a Application) GetPost(ctx context.Context, id int64) (domain.Post, error) {
	return a.db.Get(ctx, id)
}
