package ports

import (
	"context"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
)

type APIPorts interface {
	SavePost(ctx context.Context, post domain.Post) (domain.Post, error)
	GetPost(ctx context.Context, id int64) (domain.Post, error)
}
