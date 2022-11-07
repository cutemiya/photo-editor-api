package service

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image/color"
)

type RotateData struct {
	FilePath string
	Angle    float64
	Color    color.Color
}

type RotateService struct{}

func NewRotateService() RotateService {
	return RotateService{}
}

func (rc RotateService) Rotate(data RotateData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Rotate(src, data.Angle, data.Color)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
