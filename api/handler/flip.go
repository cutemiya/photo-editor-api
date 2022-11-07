package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type FlipRequest struct {
	FilePath string `json:"src"`
}

func Flip(l *zap.SugaredLogger, flipService service.FlipService, methodType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req FlipRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		flipData := service.FlipData{
			FilePath: req.FilePath,
		}

		if err := flipService.Flip(flipData, methodType); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
