package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type RotateRequest struct {
	FilePath string  `json:"src"`
	Angle    float64 `json:"angle"`
	Color    string  `json:"color"`
}

func Rotate(l *zap.SugaredLogger, rotateService service.RotateService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RotateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		color, ok := ColorMap[req.Color]
		if !ok {
			l.Errorf("incorrect filter: %s", color)
			w.WriteHeader(400)
			return
		}

		rotateData := service.RotateData{
			FilePath: req.FilePath,
			Angle:    req.Angle,
			Color:    color,
		}

		if err := rotateService.Rotate(rotateData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
