package main

import (
	"context"
	"github.com/fmo/hexagonal-blog/config"
	"github.com/fmo/hexagonal-blog/internal/adapters/db/mysql"
	"github.com/fmo/hexagonal-blog/internal/adapters/grpc"
	"github.com/fmo/hexagonal-blog/internal/adapters/image/s3"
	"github.com/fmo/hexagonal-blog/internal/application/core/api"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	ctx := context.Background()

	dbAdapter, err := mysql.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	imageAdapter, err := s3.NewAdapter(config.GetAwsRegion(), config.GetS3Bucket())
	if err != nil {
		log.Fatalf("Failed to connect to aws. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, imageAdapter)
	//restAdapter := rest.NewAdapter(application, config.GetApplicationPort())
	//restAdapter.Run(ctx)
	grpcAdapter := grpc.NewAdapter(application, config.GetGrpcPort())
	grpcAdapter.Run(ctx)
}
