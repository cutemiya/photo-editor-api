package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type GrayScaleData struct {
	FilePath string
}

type GrayScaleService struct{}

func NewGrayScaleService() GrayScaleService {
	return GrayScaleService{}
}

func (rc GrayScaleService) GrayScale(data GrayScaleData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Grayscale(src)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
