package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/hyperweb-io/starship/registry/registry"
)

type AppServer struct {
	pb.UnimplementedRegistryServer

	config *Config
	logger *zap.Logger

	// clients for various chains
	chainClients ChainClients

	grpcServer *grpc.Server
	httpServer *http.Server
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

	chainClients, err := NewChainClients(
		log,
		config,
		os.Getenv("HOME"),
	)
	if err != nil {
		return nil, err
	}

	app := &AppServer{
		config:       config,
		logger:       log,
		chainClients: chainClients,
	}

	// Validate config
	err = app.ValidateConfig()
	if err != nil {
		return nil, err
	}

	// Create grpc server
	grpcServer := grpc.NewServer(app.grpcMiddleware()...)
	pb.RegisterRegistryServer(grpcServer, app)
	app.grpcServer = grpcServer

	// Create http server
	mux := runtime.NewServeMux()
	err = pb.RegisterRegistryHandlerFromEndpoint(
		context.Background(),
		mux,
		fmt.Sprintf("%s:%s", config.Host, config.GRPCPort),
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		return nil, err
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.HTTPPort),
		Handler: app.panicRecovery(app.corsMiddleware(app.loggingMiddleware(mux))),
	}
	app.httpServer = httpServer

	return app, err
}

func (a *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.httpServer.Handler.ServeHTTP(w, r)
}

func (a *AppServer) ValidateConfig() error {
	// Verify chain ids and chain registry information
	err := verifyChainIDs(a.config)
	if err != nil {
		return err
	}

	return nil
}

func (a *AppServer) grpcMiddleware() []grpc.ServerOption {
	opts := []grpcrecovery.Option{
		grpcrecovery.WithRecoveryHandler(
			func(p interface{}) error {
				err := status.Errorf(codes.Unknown, "panic triggered: %v", p)
				a.logger.Error("panic error", zap.Error(err))
				return err
			},
		),
	}

	serverOpts := []grpc.ServerOption{
		grpcmiddleware.WithUnaryServerChain(
			grpcctxtags.UnaryServerInterceptor(),
			grpczap.UnaryServerInterceptor(a.logger),
			grpcrecovery.UnaryServerInterceptor(opts...),
		),
		grpcmiddleware.WithStreamServerChain(
			grpcctxtags.StreamServerInterceptor(),
			grpczap.StreamServerInterceptor(a.logger),
			grpcrecovery.StreamServerInterceptor(opts...),
		),
	}

	return serverOpts
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
				a.logger.Error("panic error", zap.Error(err))

				http.Error(w, ErrInternalServer.Error(), 500)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (a *AppServer) corsMiddleware(next http.Handler) http.Handler {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust this to your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
	return cors.New(corsOptions).Handler(next)
}

func (a *AppServer) Run() error {
	a.logger.Info("App starting", zap.Any("Config", a.config))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.config.Host, a.config.GRPCPort))
	if err != nil {
		a.logger.Error("failed to listen", zap.Error(err))
	}

	// Start grpc server as long-running go routine
	go func() {
		if err := a.grpcServer.Serve(lis); err != nil {
			a.logger.Error("failed to start the App gRPC server", zap.Error(err))
		}
	}()

	// Start http server
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			a.logger.Error("failed to start the App http server", zap.Error(err))
		}
	}()

	return nil
}
