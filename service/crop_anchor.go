package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type CropAnchorData struct {
	FilePath string
	Width    int
	Height   int
	Anchor   imaging.Anchor
}

type CropAnchorService struct{}

func NewCropAnchorService() CropAnchorService {
	return CropAnchorService{}
}

func (rc CropAnchorService) CropAnchor(data CropAnchorData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.CropAnchor(src, data.Width, data.Height, data.Anchor)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
