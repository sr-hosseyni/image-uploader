package img

import (
	"mime/multipart"
)

type ImageStorer interface {
	Store(srcFile multipart.File, dstPath string) error
}
