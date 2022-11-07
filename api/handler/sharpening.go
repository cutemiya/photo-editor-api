package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type SharpeningRequest struct {
	FilePath string  `json:"src"`
	Scale    float64 `json:"scale"`
}

func Sharpening(l *zap.SugaredLogger, sharpeningService service.SharpeningService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SharpeningRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		sharpeningData := service.SharpeningData{
			FilePath: req.FilePath,
			Scale:    req.Scale,
		}

		if err := sharpeningService.Sharpening(sharpeningData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
