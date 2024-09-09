package ports

type ImagePort interface {
	Upload(imageName string) error
}
