package v1

import (
	"net/http"
	"strings"

	"go.uber.org/dig"
	// 	"golang.org/x/text/language"

	"github.com/sheypoor/sheypro-common/jwt"
	"github.com/sheypoor/sheypro-common/util"
	// 	"github.com/sheypoor/sheypro-image/data"
	"github.com/sheypoor/sheypro-image/img"
)

type context struct {
	is img.ImageService
}

type apiBuilder struct {
	context
}

func (api *apiBuilder) init(c *dig.Container) error {
	return c.Invoke(func(is img.ImageService) {
		api.is = is
	})
}

func (api *apiBuilder) parseSubjectFromRequest(r *http.Request) (*jwt.AccessTokenSubject, error) {
	header := r.Header.Get("Authorization")
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return nil, ErrUnauthorized
	}

	if strings.ToLower(parts[0]) != "bearer" {
		return nil, ErrUnauthorized
	}

	key := util.GetEnv("ENC_KEY", "")

	subject, err := jwt.ParseAccessToken(parts[1], key)
	if err != nil {
		switch err {
		case jwt.ErrInvalidAccessToken:
			return nil, ErrInvalidAccessToken
		case jwt.ErrExpiredAccessToken:
			return nil, ErrExpiredAccessToken
		default:
			return nil, err
		}
	}

	return subject, nil
}
