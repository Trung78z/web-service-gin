package repositories

import (
	"github.com/Trung78z/web-service-gin/internal/repositories/queries"
	"github.com/Trung78z/web-service-gin/pkg/database"
)

var Queries *queries.Queries

func InitRepository() {
	Queries = queries.New(database.DB)
}
