package data

import (
	"io"
	"os"

	"github.com/sheypoor/sheypro-image/img"
	"mime/multipart"
)

func NewFileImageStorer() img.ImageStorer {
	return &fileImageStorer{}
}

type fileImageStorer struct {
}

func (fis *fileImageStorer) Store(srcFile multipart.File, dstPath string) error {
	f, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = io.Copy(f, srcFile)
	return err
}
