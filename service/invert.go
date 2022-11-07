package service

import (
	"fmt"
	"github.com/disintegration/imaging"
)

type InvertData struct {
	FilePath string
}

type InvertService struct{}

func NewInvertService() InvertService {
	return InvertService{}
}

func (rc InvertService) Invert(data InvertData) error {
	src, err := imaging.Open(data.FilePath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("fail to open: %w", err)
	}

	image := imaging.Invert(src)
	if err = imaging.Save(image, "res.jpg"); err != nil {
		return err
	}

	return nil
}
