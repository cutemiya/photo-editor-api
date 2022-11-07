package service

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
)

type FlipData struct {
	FilePath string
}

type FlipService struct{}

func NewFlipService() FlipService {
	return FlipService{}
}

func (rc FlipService) Flip(data FlipData, methodType string) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	var image *image.NRGBA
	switch methodType {
	case "vertical":
		image = imaging.FlipV(src)
	case "horizontal":
		image = imaging.FlipH(src)
	}

	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
