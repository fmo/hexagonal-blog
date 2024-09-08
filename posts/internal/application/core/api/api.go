package api

import (
	"context"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
	"github.com/fmo/hexagonal-blog/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) SavePost(ctx context.Context, post domain.Post) (domain.Post, error) {
	err := a.db.Save(ctx, &post)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

func (a Application) GetPost(ctx context.Context, id int64) (domain.Post, error) {
	return a.db.Get(ctx, id)
}
