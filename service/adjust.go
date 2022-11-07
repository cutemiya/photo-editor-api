package service

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
)

type AdjustData struct {
	FilePath string
	Scale    float64
}

type AdjustService struct{}

func NewAdjustService() AdjustService {
	return AdjustService{}
}

func (rc AdjustService) Adjust(data AdjustData, methodType string) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	var image *image.NRGBA
	switch methodType {
	case "gamma":
		image = imaging.AdjustGamma(src, data.Scale)
	case "contrast":
		image = imaging.AdjustContrast(src, data.Scale)
	case "saturation":
		image = imaging.AdjustSaturation(src, data.Scale)
	case "brightness":
		image = imaging.AdjustBrightness(src, data.Scale)
	}

	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
