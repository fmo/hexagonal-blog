package ports

type ImagePort interface {
	Upload(imageName, imageUrl string) error
}
