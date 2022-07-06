package repo

import (
	pb "github.com/Muhammadjon226/first_service/genproto/first_service"
)

// FirstStorageI ...
type FirstStorageI interface {
	CreatePostsFromApi([]*pb.Post) error
	// Get(id string) (*pb.Post, error)
	// List(page, limit int64) ([]*pb.Post, int64, error)
	// Update(pb.Post) (pb.Post, error)
	// Delete(id string) error
}
