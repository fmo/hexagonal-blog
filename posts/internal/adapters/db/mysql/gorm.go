package mysql

import (
	"context"
	"fmt"
	"github.com/fmo/hexagonal-blog/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Post{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Save(ctx context.Context, post *domain.Post) error {
	postModel := Post{
		Title: post.Title,
		Body:  post.Body,
	}
	res := a.db.WithContext(ctx).Create(&postModel)
	if res.Error == nil {
		post.ID = int64(postModel.ID)
	}
	return res.Error
}

func (a Adapter) Get(ctx context.Context, id int64) (domain.Post, error) {
	var postEntity Post
	res := a.db.WithContext(ctx).First(&postEntity, id)
	post := domain.Post{
		ID:        int64(postEntity.ID),
		Title:     postEntity.Title,
		Body:      postEntity.Body,
		CreatedAt: postEntity.CreatedAt.UnixNano(),
	}
	return post, res.Error
}
