package handlers

import (
	"net/http"

	"github.com/showntop/circle-core/logger"
)

type Home struct {
}

func (h Home) Index(w http.ResponseWriter, r *http.Request) {
	logger.Info("this is the home page")
	w.Write([]byte("this is the home page"))
}
