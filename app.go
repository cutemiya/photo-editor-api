package main

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"photo-editor/api"
	"photo-editor/config"
	"photo-editor/lib/pctx"
	"photo-editor/service"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {

	var (
		resizeService     = service.NewResizeService()
		cropAnchorService = service.NewCropAnchorService()
		fillService       = service.NewFillService()
		fitService        = service.NewFitService()
		blurService       = service.NewBlurService()
		sharpeningService = service.NewSharpeningService()
		adjustService     = service.NewAdjustService()
		grayScaleService  = service.NewGrayScaleService()
		invertService     = service.NewInvertService()
		flipService       = service.NewFlipService()
		rotateService     = service.NewRotateService()

		server = api.NewServer(
			ctxProvider,
			logger,
			settings,
			resizeService,
			cropAnchorService,
			fillService,
			fitService,
			blurService,
			sharpeningService,
			adjustService,
			grayScaleService,
			invertService,
			flipService,
			rotateService)
	)

	return App{
		logger:   logger,
		settings: settings,
		server:   server,
	}
}

func (a App) Run() {
	go func() {
		_ = a.server.ListenAndServe()
	}()
	a.logger.Debugf("HTTP server started on %d", a.settings.Port)
}

func (a App) Stop(ctx context.Context) {
	_ = a.server.Shutdown(ctx)
	a.logger.Debugf("HTTP server stopped")
}
