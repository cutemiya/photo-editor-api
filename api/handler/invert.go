package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type InvertRequest struct {
	FilePath string `json:"src"`
}

func Invert(l *zap.SugaredLogger, invertService service.InvertService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req InvertRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		invertData := service.InvertData{
			FilePath: req.FilePath,
		}

		if err := invertService.Invert(invertData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
