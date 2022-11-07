package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type SharpeningData struct {
	FilePath string
	Scale    float64
}

type SharpeningService struct{}

func NewSharpeningService() SharpeningService {
	return SharpeningService{}
}

func (rc SharpeningService) Sharpening(data SharpeningData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Sharpen(src, data.Scale)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
