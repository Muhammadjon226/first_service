package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbFirst "github.com/Muhammadjon226/first_service/genproto/first_service"
	l "github.com/Muhammadjon226/first_service/pkg/logger"
	"github.com/Muhammadjon226/first_service/pkg/utils"
	"github.com/Muhammadjon226/first_service/storage"
)

// FirstService ...
type FirstService struct {
	logger  l.Logger
	storage storage.IStorage
}

// NewFirstService ...
func NewFirstService(storage storage.IStorage, log l.Logger) *FirstService {

	return &FirstService{
		logger:  log,
		storage: storage,
	}
}

//GetPostsFromOpenAPI for get posts from open api
func (fs *FirstService) GetPostsFromOpenAPI(ctx context.Context, req *pbFirst.EmptyResp) (*pbFirst.EmptyResp, error) {

	posts, err := utils.GetPostsFromOpenAPIHandler()
	if err != nil {
		fs.logger.Error("failed to get posts from api", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get posts from api")
	}
	err = fs.storage.FirstService().CreatePostsFromAPI(posts)
	if err != nil {
		fs.logger.Error("failed to create Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create Post")
	}

	return &pbFirst.EmptyResp{}, nil
}

//CreatePost ...
func (fs *FirstService) CreatePost(ctx context.Context, req *pbFirst.Post) (*pbFirst.PostResponse, error) {

	post, err := fs.storage.FirstService().Create(req)
	if err != nil {
		fs.logger.Error("failed to create Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create Post")
	}

	return post, nil
}

//GetPostByID ...
func (fs *FirstService) GetPostByID(ctx context.Context, req *pbFirst.ByIdReq) (*pbFirst.PostResponse, error) {

	post, err := fs.storage.FirstService().Get(req)
	if err != nil {
		fs.logger.Error("failed to get Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get Post")
	}

	return post, nil
}

//ListPosts ...
func (fs *FirstService) ListPosts(ctx context.Context, req *pbFirst.ListReq) (*pbFirst.ListResp, error) {
	posts, err := fs.storage.FirstService().List(req)
	if err != nil {
		fs.logger.Error("failed to list Posts", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list Posts")
	}

	return &pbFirst.ListResp{
		Posts: posts.Posts,
		Count: posts.Count,
	}, nil
}

//UpdatePost ...
func (fs *FirstService) UpdatePost(ctx context.Context, req *pbFirst.Post) (*pbFirst.PostResponse, error) {
	post, err := fs.storage.FirstService().Update(req)
	if err != nil {
		fs.logger.Error("failed to update Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update Post")
	}

	return post, nil
}

//DeletePost ...
func (fs *FirstService) DeletePost(ctx context.Context, req *pbFirst.ByIdReq) (*pbFirst.EmptyResp, error) {
	err := fs.storage.FirstService().Delete(req)
	if err != nil {
		fs.logger.Error("failed to delete Post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete Post")
	}

	return &pbFirst.EmptyResp{}, nil
}
