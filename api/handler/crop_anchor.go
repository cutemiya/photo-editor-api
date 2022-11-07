package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/service"
)

type CropAnchorRequest struct {
	FilePath string `json:"src"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Anchor   string `json:"anchor"`
}

func CropAnchor(l *zap.SugaredLogger, cropAnchorService service.CropAnchorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CropAnchorRequest
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

		cropAnchorData := service.CropAnchorData{
			FilePath: req.FilePath,
			Width:    req.Width,
			Height:   req.Height,
			Anchor:   anchor,
		}

		if err := cropAnchorService.CropAnchor(cropAnchorData); err != nil {
			l.Error(err)
			w.WriteHeader(400)
			return
		}
	}
}
