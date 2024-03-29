package data

import (
	"fmt"
	"review-service/internal/conf"
	"review-service/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewReviewRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	query  *query.Query
	logger *log.Helper
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	query.SetDefault(db)
	return &Data{query: query.Q, logger: log.NewHelper(logger)}, cleanup, nil
}
