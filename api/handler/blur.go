package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type BlurRequest struct {
	FilePath string  `json:"src"`
	Scale    float64 `json:"scale"`
}

func Blur(l *zap.SugaredLogger, blurService service.BlurService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req BlurRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		blurData := service.BlurData{
			FilePath: req.FilePath,
			Scale:    req.Scale,
		}

		if err := blurService.Blur(blurData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
