package server

import (
	"github.com/SyntSugar/ss-infra-go/api/server"
	"github.com/SyntSugar/ss-infra-go/log"
	"github.com/ccbond/cetus-ai/internal/config"
	"github.com/ccbond/cetus-ai/internal/logger"
	"go.uber.org/zap"

	openai "github.com/sashabaranov/go-openai"
)

type Services struct {
	OpenaiClient *openai.Client
}

type Server struct {
	apiServer *server.Server
	logger    *log.Logger
	config    *config.Config
	svcs      *Services
}

func NewServer(cfg *config.Config, services *Services) (*Server, error) {
	if err := logger.Init(cfg.LogConfig.LogLevel); err != nil {
		return nil, err
	}
	apiServer, err := server.New(&server.Config{
		API:   cfg.API,
		Admin: cfg.Admin,
	}, logger.Get())
	if err != nil {
		return nil, err
	}
	return &Server{
		apiServer: apiServer,
		config:    cfg,
		logger:    logger.Get(),
		svcs:      services,
	}, nil
}

func (srv *Server) Run() error {
	setupAPIRouters(srv)
	if err := srv.apiServer.Run(); err != nil {
		return err
	}
	logger.Get().With(
		zap.Any("api", srv.config.API),
		zap.Any("admin", srv.config.Admin),
	).Info("The server was listening")
	return nil
}

func (srv *Server) Shutdown() {
	if err := srv.apiServer.Shutdown(); err != nil {
		srv.logger.With(zap.String("err", err.Error())).Error("Shutdown error")
	}
	srv.logger.Info("The server was shutdown normally, see you lala.")
}
