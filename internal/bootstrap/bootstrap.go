package bootstrap

import (
	"fmt"

	hcontract "github.com/Clausia-Ifest/clausia-server/internal/app/contract/handler"
	rcontract "github.com/Clausia-Ifest/clausia-server/internal/app/contract/repository"
	ucontract "github.com/Clausia-Ifest/clausia-server/internal/app/contract/usecase"
	hdocument "github.com/Clausia-Ifest/clausia-server/internal/app/documents/handler"
	rdocument "github.com/Clausia-Ifest/clausia-server/internal/app/documents/repository"
	udocument "github.com/Clausia-Ifest/clausia-server/internal/app/documents/usecase"
	huser "github.com/Clausia-Ifest/clausia-server/internal/app/user/handler"
	ruser "github.com/Clausia-Ifest/clausia-server/internal/app/user/repository"
	uuser "github.com/Clausia-Ifest/clausia-server/internal/app/user/usecase"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/blockchain"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/config"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/db"
	rpc "github.com/Clausia-Ifest/clausia-server/internal/infra/grpc"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/http"
	logger "github.com/Clausia-Ifest/clausia-server/internal/infra/log"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/storage"
	"github.com/Clausia-Ifest/clausia-server/internal/middleware"
	"github.com/Clausia-Ifest/clausia-server/pkg/hash"
	querybuilder "github.com/Clausia-Ifest/clausia-server/pkg/query_builder"
	jwt "github.com/Clausia-Ifest/clausia-server/pkg/token"
	"github.com/Clausia-Ifest/clausia-server/pkg/transactor"
	"github.com/Clausia-Ifest/clausia-server/pkg/validator"
	clausiapb "github.com/Clausia-Ifest/clausia-server/proto"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config *config.Env
	db     db.Postgres
	server *fiber.App
	s3     storage.IS3
	grpc   clausiapb.ClausIAClient
}

func Initialize() error {
	_config, err := config.Load()
	if err != nil {
		return err
	}

	_db, err := db.NewPostgres(_config.DB)
	if err != nil {
		return err
	}

	if err := logger.NewZeroLog(_config.AppEnv); err != nil {
		return err
	}

	rpc, err := rpc.New(_config.GRPCHost, _config.GRPCPort)
	if err != nil {
		return err
	}

	grpcConn := clausiapb.NewClausIAClient(rpc)

	app := &App{
		config: _config,
		db:     _db,
		server: http.NewFiber(_config.AppEnv),
		s3:     storage.New(_config.S3AccessKey, _config.S3SecretKey, _config.S3Region, _config.S3Endpoint, _config.S3BucketName),
		grpc:   grpcConn,
	}

	if err := app.handleFlags(); err != nil {
		return err
	}

	return app.start()
}

func (app *App) start() error {
	app.health()
	app.registerRoutes()

	return app.server.Listen(fmt.Sprintf(":%d", app.config.AppPort))
}

func (app *App) health() {
	app.server.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK 🧘",
		})
	})
}

func (app *App) registerRoutes() {
	transactor := transactor.New(app.db.GetConnection())
	queryBuilder := querybuilder.New()
	blockchain := blockchain.New(app.config.WEB3InfuraRPC, app.config.WEB3PrivateKey, app.config.WEB3ContractAddress)

	val10 := validator.New()
	hash := hash.New()
	token := jwt.New(app.config.JWTSecret, app.config.JWTDuration)
	middleware := middleware.New(token)

	repositoryUser := ruser.New(queryBuilder)
	usecaseUser := uuser.New(transactor, hash, token, repositoryUser)
	handlerUser := huser.New(val10, middleware, usecaseUser)

	repositoryDocument := rdocument.New(queryBuilder)
	usecaseDocument := udocument.New(transactor, app.s3, app.grpc, repositoryDocument)
	handlerDocument := hdocument.New(val10, middleware, usecaseDocument)

	repositoryContract := rcontract.New(queryBuilder)
	usecaseContract := ucontract.New(transactor, app.s3, app.grpc, blockchain, repositoryContract, repositoryDocument)
	handlerContract := hcontract.New(val10, middleware, usecaseContract)

	router := app.server.Group("")
	handlerUser.MountRoutes(router)
	handlerDocument.MountRoutes(router)
	handlerContract.MountRoutes(router)
}
