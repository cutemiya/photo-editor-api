package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type FillRequest struct {
	FilePath string `json:"src"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Anchor   string `json:"anchor"`
	Filter   string `json:"filter"`
}

func Fill(l *zap.SugaredLogger, fillService service.FillService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req FillRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}

		anchor, ok := AnchorMap[req.Anchor]
		if !ok {
			l.Errorf("incorrect anchor: %s", anchor)
			w.WriteHeader(400)
			return
		}

		filter, ok := ResampleFilterMap[req.Filter]
		if !ok {
			l.Errorf("incorrect anchor: %s", filter)
			w.WriteHeader(400)
			return
		}

		fillData := service.FillData{
			FilePath: req.FilePath,
			Width:    req.Width,
			Height:   req.Height,
			Anchor:   anchor,
			Filter:   filter,
		}

		if err := fillService.Fill(fillData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
