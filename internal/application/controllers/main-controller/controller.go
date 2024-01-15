package main_controller

import (
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"net/http"
)

type controller struct {
	log logger.Logger
}

func NewController(log logger.Logger) *controller {
	return &controller{
		log: log,
	}
}

func (c *controller) MainHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Infof("handle default at: %s", r.URL.Path)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello cakes!"))
}
