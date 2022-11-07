package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type ResizeData struct {
	FilePath string
	Width    int
	Height   int
	Filter   imaging.ResampleFilter
}

type ResizeService struct{}

func NewResizeService() ResizeService {
	return ResizeService{}
}

func (rc ResizeService) Resize(data ResizeData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Resize(src, data.Width, data.Height, data.Filter)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
