package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type ResizeRequest struct {
	FilePath string `json:"src"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filter   string `json:"filter"`
}

func Resizing(l *zap.SugaredLogger, resizeService service.ResizeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ResizeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		filter, ok := ResampleFilterMap[req.Filter]
		if !ok {
			l.Errorf("incorrect filter: %s", filter)
			w.WriteHeader(400)
			return
		}

		resizeData := service.ResizeData{
			FilePath: req.FilePath,
			Width:    req.Width,
			Height:   req.Height,
			Filter:   filter,
		}

		if err := resizeService.Resize(resizeData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
