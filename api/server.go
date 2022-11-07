package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"photo-editor/api/handler"
	"photo-editor/config"
	"photo-editor/lib/pctx"
	"photo-editor/service"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	resizeService service.ResizeService,
	cropAnchorService service.CropAnchorService,
	fillService service.FillService,
	fitService service.FitService,
	blurService service.BlurService,
	sharpeningService service.SharpeningService,
	adjustService service.AdjustService,
	grayScaleService service.GrayScaleService,
	invertService service.InvertService,
	flipService service.FlipService,
	rotateService service.RotateService,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/resizing", handler.Resizing(logger, resizeService)).Methods(http.MethodPost)
	router.HandleFunc("/cropanchor", handler.CropAnchor(logger, cropAnchorService)).Methods(http.MethodPost)
	router.HandleFunc("/fill", handler.Fill(logger, fillService)).Methods(http.MethodPost)
	router.HandleFunc("/fit", handler.Fit(logger, fitService)).Methods(http.MethodPost)
	router.HandleFunc("/blur", handler.Blur(logger, blurService)).Methods(http.MethodPost)
	router.HandleFunc("/sharpening", handler.Sharpening(logger, sharpeningService)).Methods(http.MethodPost)
	router.HandleFunc("/adjust/gamma", handler.Adjust(logger, adjustService, "gamma")).Methods(http.MethodPost)
	router.HandleFunc("/adjust/contrast", handler.Adjust(logger, adjustService, "contrast")).Methods(http.MethodPost)
	router.HandleFunc("/adjust/brightness", handler.Adjust(logger, adjustService, "brightness")).Methods(http.MethodPost)
	router.HandleFunc("/adjust/saturation", handler.Adjust(logger, adjustService, "saturation")).Methods(http.MethodPost)
	router.HandleFunc("/grayscale", handler.GrayScale(logger, grayScaleService)).Methods(http.MethodPost)
	router.HandleFunc("/invert", handler.Invert(logger, invertService)).Methods(http.MethodPost)
	router.HandleFunc("/flip/vertical", handler.Flip(logger, flipService, "vertical")).Methods(http.MethodPost)
	router.HandleFunc("/flip/horizontal", handler.Flip(logger, flipService, "horizontal")).Methods(http.MethodPost)
	router.HandleFunc("/rotate", handler.Rotate(logger, rotateService)).Methods(http.MethodPost)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
