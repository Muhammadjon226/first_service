package postgres

import (
	"github.com/jmoiron/sqlx"

	pb "github.com/Muhammadjon226/first_service/genproto/first_service"
)

type firstRepo struct {
	db *sqlx.DB
}

// NewFirstRepo
func NewFirstRepo(db *sqlx.DB) *firstRepo {
	return &firstRepo{db: db}
}

// CreatePostsFromApi ...
func (fr *firstRepo) CreatePostsFromApi(in []*pb.Post) error {

	return nil
}
