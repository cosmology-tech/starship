package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AppServer struct {
	config *Config
	logger *zap.Logger
	server *grpc.Server
	router http.Handler
}

func NewAppServer(config *Config) (*AppServer, error) {
	log, err := NewLogger(config)
	if err != nil {
		return nil, err
	}
	log.Info(
		"Starting the service",
		zap.String("prog", Prog),
		zap.String("version", Version),
		zap.Any("config", config),
	)

	app := &AppServer{
		config: config,
		logger: log,
	}

	server := grpc.NewServer()
	reflection.Register(server)

	// Setup routes
	router, err := app.Router()
	if err != nil {
		log.Error("Error setting up routes", zap.Error(err))
		return nil, err
	}
	app.router = router

	return app, err
}

func (a *AppServer) Router() (*chi.Mux, error) {
	router := chi.NewRouter()
	router.MethodNotAllowed(MethodNotAllowed)
	router.NotFound(NotFound)

	// Set middleware
	router.Use(a.panicRecovery)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Setup routes
	router.Get("/chain_ids", a.GetChainIDs)
	router.Route("/chains", func(r chi.Router) {
		r.Get("/", a.GetChains)
		r.Get("/{chain}", a.GetChain)
		r.Get("/{chain}/assets", a.GetChainAssets)
	})
	router.Route("/ibc", func(r chi.Router) {
		r.Get("/", a.GetAllIBC)
		r.Get("/{chain-a}/{chain-b}", a.GetIBCChainsData)
		r.Post("/{chain-a}/{chain-b}", a.SetIBCChainsData)
		r.Get("/{chain-a}/{chain-b}/channels", a.GetIBCChainsChannels)
		r.Patch("/{chain-a}/{chain-b}/channels", a.AddIBCChainChannel)
	})

	return router, nil
}

func (a *AppServer) loggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		defer func() {
			a.logger.Info("client request",
				zap.Duration("latency", time.Since(start)),
				zap.Int("status", ww.Status()),
				zap.Int("bytes", ww.BytesWritten()),
				zap.String("client_ip", r.RemoteAddr),
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("request-id", middleware.GetReqID(r.Context())))
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}

func (a *AppServer) panicRecovery(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rc := recover(); rc != nil {
				err, ok := rc.(error)
				if !ok {
					err = fmt.Errorf("panic: %v", rc)
				}
				a.logger.Error("panic error",
					zap.String("request-id", middleware.GetReqID(r.Context())),
					zap.Error(err))

				render.Render(w, r, ErrInternalServer)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (a *AppServer) Run() error {
	a.logger.Info("App starting", zap.Any("Config", a.config))

	// Setup server
	server := &http.Server{
		Addr:    a.config.Addr,
		Handler: a.router,
	}
	a.server = server

	// Start http server as long-running go routine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			a.logger.Error("failed to start the App HTTP server", zap.Error(err))
		}
	}()

	return nil
}
