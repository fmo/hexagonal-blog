package ports

import (
	"context"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
)

type DBPort interface {
	Save(ctx context.Context, post *domain.Post) error
	Get(ctx context.Context, id int64) (domain.Post, error)
}
