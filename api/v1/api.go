package v1

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/dig"

	"github.com/sheypoor/sheypro-common/api"
)

func Build(r *httprouter.Router, c *dig.Container) error {
	builder := &apiBuilder{}
	if err := builder.init(c); err != nil {
		return err
	}

	// Serve image
	r.GET("/v1/image", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// TODO implement
	})

	// Handle image upload
	r.POST("/v1/image", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		r.ParseMultipartForm(32 << 20)
		//ParseMultipartForm parses a request body as multipart/form-data

		file, _, err := r.FormFile("image")
		if err != nil {
			fmt.Println(err)
			api.Failure(w, err)
			return
		}
		defer file.Close()

		// if err != nil {
		// 	fmt.Println(err)
		// 	api.Failure(w, err)
		// 	return
		// }

		path, err := builder.is.Upload(file)
		if err != nil {
			fmt.Println(err)
			api.Failure(w, err)
			return
		}

		api.Success(w, path)
	})

	return nil
}
