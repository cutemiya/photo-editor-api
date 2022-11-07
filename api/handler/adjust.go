package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type AdjustRequest struct {
	FilePath string  `json:"src"`
	Scale    float64 `json:"scale"`
}

func Adjust(l *zap.SugaredLogger, adjustService service.AdjustService, methodType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AdjustRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		adjustData := service.AdjustData{
			FilePath: req.FilePath,
			Scale:    req.Scale,
		}

		if err := adjustService.Adjust(adjustData, methodType); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
