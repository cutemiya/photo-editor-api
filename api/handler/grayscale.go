package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type GrayScaleRequest struct {
	FilePath string `json:"src"`
}

func GrayScale(l *zap.SugaredLogger, grayScaleService service.GrayScaleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GrayScaleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		grayScaleData := service.GrayScaleData{
			FilePath: req.FilePath,
		}

		if err := grayScaleService.GrayScale(grayScaleData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
