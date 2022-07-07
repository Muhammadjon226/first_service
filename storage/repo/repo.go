package repo

import (
	"github.com/Muhammadjon226/first_service/models"
	pbFirst "github.com/Muhammadjon226/first_service/genproto/first_service"
)

// FirstStorageI ...
type FirstStorageI interface {
	CreatePostsFromAPI([]models.Data) error
	Create(*pbFirst.Post) (*pbFirst.PostResponse, error)
	Get(*pbFirst.ByIdReq) (*pbFirst.PostResponse, error)
	List(*pbFirst.ListReq) (*pbFirst.ListResp, error)
	Update(*pbFirst.Post) (*pbFirst.PostResponse, error)
	Delete(*pbFirst.ByIdReq) error
}
