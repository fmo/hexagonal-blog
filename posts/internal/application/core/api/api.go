package api

import (
	"context"
	"fmt"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
	"github.com/fmo/hexagonal-blog/internal/ports"
	log "github.com/sirupsen/logrus"
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

	imageName := fmt.Sprintf("%d.png", post.ID)

	if a.image.CheckImageAlreadyUploaded(imageName) {
		log.Infof("Image %s is already uploaded", imageName)
	} else {
		log.Infof("Image %s is not uploaded, so doing it.", imageName)
		err = a.image.Upload(
			imageName,
			"https://upload.wikimedia.org/wikipedia/commons/a/af/Tux.png",
		)
		if err != nil {
			return domain.Post{}, err
		}
	}

	return post, nil
}

func (a Application) GetPost(ctx context.Context, id int64) (domain.Post, error) {
	return a.db.Get(ctx, id)
}
