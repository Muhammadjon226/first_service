package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/Muhammadjon226/first_service/storage/postgres"
	"github.com/Muhammadjon226/first_service/storage/repo"
)

// IStorage ...
type IStorage interface {
	FirstService() repo.FirstStorageI
}

//storagePg ...
type storagePg struct {
	db       *sqlx.DB
	firstRepo repo.FirstStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) IStorage {
	return &storagePg{
		db:       db,
		firstRepo: postgres.NewFirstRepo(db),
	}
}

func (s storagePg) FirstService() repo.FirstStorageI {
	return s.firstRepo
}
