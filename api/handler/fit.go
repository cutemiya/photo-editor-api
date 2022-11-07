package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type FitRequest struct {
	FilePath string `json:"src"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filter   string `json:"filter"`
}

func Fit(l *zap.SugaredLogger, fitService service.FitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req FitRequest
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

		fitData := service.FitData{
			FilePath: req.FilePath,
			Width:    req.Width,
			Height:   req.Height,
			Filter:   filter,
		}

		if err := fitService.Fit(fitData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
