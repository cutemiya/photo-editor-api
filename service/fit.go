package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type FitData struct {
	FilePath string
	Width    int
	Height   int
	Filter   imaging.ResampleFilter
}

type FitService struct{}

func NewFitService() FitService {
	return FitService{}
}

func (rc FitService) Fit(data FitData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Fit(src, data.Width, data.Height, data.Filter)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
