package handlers

import (
	"net/http"

	"github.com/showntop/circle-core/logger"
)

type Home struct {
	Handler
}

func (h Home) Index(w http.ResponseWriter, r *http.Request) {
	logger.Info("strstrstrstrstrstrstrstrstrstrstrstrstrstrstrstrstrstr")
	return Handler{func(w, r) {

	}}
}
