package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type BlurData struct {
	FilePath string
	Scale    float64
}

type BlurService struct{}

func NewBlurService() BlurService {
	return BlurService{}
}

func (rc BlurService) Blur(data BlurData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Blur(src, data.Scale)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
