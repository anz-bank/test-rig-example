package dbfront_impl

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/validator"
	"github.com/go-chi/chi"
)

type Callback struct {
	UpstreamTimeout   time.Duration
	DownstreamTimeout time.Duration
	RouterBasePath    string
	UpstreamConfig    validator.Validator
	db                *sql.DB
}

type Config struct{}

func (c Config) Validate() error {
	return nil
}

func (g Callback) AddMiddleware(ctx context.Context, r chi.Router) {
}

func (g Callback) BasePath() string {
	return "/"
}

func (g Callback) GetDBHandle() (*sql.DB, error) {
	return g.db, nil
}

func (g Callback) Config() validator.Validator {
	return Config{}
}

func (g Callback) HandleError(ctx context.Context, w http.ResponseWriter, kind common.Kind, message string, cause error) {
	se := common.CreateError(ctx, kind, message, cause)
	httpError := common.HandleError(ctx, se)
	httpError.WriteError(ctx, w)
}

func (g Callback) DownstreamTimeoutContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 120*time.Second)
}

func GetCallback(db *sql.DB) Callback {
	return Callback{
		UpstreamTimeout:   120 * time.Second,
		DownstreamTimeout: 120 * time.Second,
		RouterBasePath:    "/",
		UpstreamConfig:    nil,
		db:                db,
	}
}
