package handler

import (
	"github.com/disintegration/imaging"
	"image/color"
)

var (
	ResampleFilterMap = map[string]imaging.ResampleFilter{
		"lanczos":           imaging.Lanczos,
		"catmullrom":        imaging.CatmullRom,
		"mitchellnetravali": imaging.MitchellNetravali,
		"linear":            imaging.Linear,
		"box":               imaging.Box,
		"nearestneighbor":   imaging.NearestNeighbor,
	}
)

var (
	AnchorMap = map[string]imaging.Anchor{
		"center":      imaging.Center,
		"top":         imaging.Top,
		"left":        imaging.Left,
		"right":       imaging.Right,
		"bottom":      imaging.Bottom,
		"topleft":     imaging.TopLeft,
		"topright":    imaging.TopRight,
		"bottomleft":  imaging.BottomLeft,
		"bottomright": imaging.BottomRight,
	}
)

var (
	ColorMap = map[string]color.Color{
		"red":   color.RGBA{R: 255, A: 255},
		"green": color.RGBA{G: 255, A: 255},
		"blue":  color.RGBA{B: 255, A: 255},
		"white": color.White,
		"black": color.Black,
	}
)
