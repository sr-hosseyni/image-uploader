package img

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

// PUBLIC ------------------------------------------------------------

type ImageService interface {
	Upload(file multipart.File) (string, error)
}

func NewImageService(storer ImageStorer) ImageService {
	return &imageServiceImpl{storer}
}

// PRIVATE -----------------------------------------------------------

type imageServiceImpl struct {
	storer ImageStorer
}

func (imgs *imageServiceImpl) Upload(file multipart.File) (string, error) {
	os.Mkdir("public/storage", os.ModePerm)
	uuid, _ := imgs.generateName()
	dstPath := "public/storage/" + uuid
	err := imgs.storer.Store(file, dstPath)
	return dstPath, err
}

func (imgs *imageServiceImpl) generateName() (string, error) {
	uuid := make([]byte, 32)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", uuid), nil
}
