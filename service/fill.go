package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type FillData struct {
	FilePath string
	Width    int
	Height   int
	Anchor   imaging.Anchor
	Filter   imaging.ResampleFilter
}

type FillService struct{}

func NewFillService() FillService {
	return FillService{}
}

func (rc FillService) Fill(data FillData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Fill(src, data.Width, data.Height, data.Anchor, data.Filter)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
