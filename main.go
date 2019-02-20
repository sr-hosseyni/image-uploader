package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sheypoor/sheypro-common/util"
	"github.com/sheypoor/sheypro-image/api/v1"

	"github.com/sheypoor/sheypro-image/data"
	"github.com/sheypoor/sheypro-image/img"

	"go.uber.org/dig"
)

func main() {
	var err error

	// Check configuration
	if key := util.GetEnv("ENC_KEY", ""); key == "" {
		panic("encryption key is not set")
	}

	// Initialize DI
	container := buildContainer()

	// Start server
	err = container.Invoke(func(r *httprouter.Router) error {
		addr := ":" + util.GetEnv("PORT", "8080")
		log.Print("Start listening on " + addr)
		return http.ListenAndServe(addr, r)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	// Storers
	_ = container.Provide(data.NewFileImageStorer)

	// Services
	container.Provide(img.NewImageService)

	// Routes
	container.Provide(initRouter(container))

	return container
}

func initRouter(c *dig.Container) func() (*httprouter.Router, error) {
	return func() (r *httprouter.Router, err error) {
		r = httprouter.New()
		if err = v1.Build(r, c); err != nil {
			return
		}
		return
	}
}
