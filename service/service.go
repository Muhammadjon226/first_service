package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbFirst "github.com/Muhammadjon226/first_service/genproto/first_service"
	l "github.com/Muhammadjon226/first_service/pkg/logger"
	"github.com/Muhammadjon226/first_service/storage"
	"github.com/jmoiron/sqlx"
)

// FirstService ...
type FirstService struct {
	logger  l.Logger
	storage storage.IStorage
}

// NewFirstService ...
func NewFirstService(db *sqlx.DB, log l.Logger) *FirstService {

	return &FirstService{
		logger: log,
	}
}

func (fs *FirstService) GetPostsFromOpenApi(ctx context.Context, req *pbFirst.EmptyResp) (*pbFirst.EmptyResp, error) {

	var request []*pbFirst.Post
	err := fs.storage.FirstService().CreatePostsFromApi(request)
	if err != nil {
		fs.logger.Error("failed to create Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create Post")
	}

	return &pbFirst.EmptyResp{}, nil
}