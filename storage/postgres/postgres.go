package postgres

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	pbFirst "github.com/Muhammadjon226/first_service/genproto/first_service"
	"github.com/Muhammadjon226/first_service/models"
	"github.com/Muhammadjon226/first_service/pkg/helper"
	"github.com/Muhammadjon226/first_service/pkg/newerror"
	"github.com/Muhammadjon226/first_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type firstRepo struct {
	db *sqlx.DB
}

//NewFirstRepo ...
func NewFirstRepo(db *sqlx.DB) repo.FirstStorageI {
	return &firstRepo{db: db}
}

//CreatePostsFromApi ...
func (fr *firstRepo) CreatePostsFromAPI(in []models.Data) error {
	tx, err := fr.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			// rollback transaction before return in case of error
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	createPostsStr := `
		INSERT INTO posts (
			id,
			user_id,
			title,
			body)
		VALUES `
	vals := []interface{}{}

	for _, post := range in {
		createPostsStr += `(?, ?, ?, ?),`
		vals = append(vals,
			post.ID,
			post.UserID,
			post.Title,
			post.Body,
		)
	}
	// trim the last ,
	createPostsStr = strings.TrimSuffix(createPostsStr, ",")

	// Replacing ? with $n for postgres
	createPostsStr = helper.ReplaceSQL(createPostsStr, "?")

	stmt, err := tx.Prepare(createPostsStr)
	if err != nil {
		log.Println(err)
		return err
	}

	// format all vals at once
	_, err = stmt.Exec(vals...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (fr *firstRepo) Create(input *pbFirst.Post) (*pbFirst.PostResponse, error) {

	response := pbFirst.PostResponse{}
	var updatedAt sql.NullString
	_, err := fr.db.Exec(`
			INSERT INTO posts (id, user_id, title, body) VALUES ($1, $2, $3, $4)
		`,
		input.Id,
		input.UserId,
		input.Title,
		input.Body,
	)
	err = fr.db.QueryRow(`SELECT created_at, updated_at from posts WHERE id = $1`, input.Id).Scan(
		&response.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	response.Id = input.Id
	response.Title = input.Title
	response.Body = input.Body
	response.UserId = input.UserId
	response.UpdatedAt = updatedAt.String

	return &response, nil
}

func (fr *firstRepo) Get(input *pbFirst.ByIdReq) (*pbFirst.PostResponse, error) {

	response := pbFirst.PostResponse{}
	var updatedAt sql.NullString
	err := fr.db.QueryRow(`
			SELECT * FROM posts WHERE id = $1`, input.Id,
	).Scan(
		&response.Id,
		&response.UserId,
		&response.Title,
		&response.Body,
		&response.CreatedAt,
		&updatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, newerror.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	response.UpdatedAt = updatedAt.String

	return &response, nil
}

func (fr *firstRepo) Delete(input *pbFirst.ByIdReq) error {

	result, err := fr.db.Exec(`DELETE FROM posts WHERE id = $1`, input.Id)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return newerror.ErrNotFound
	}

	return nil
}

func (fr *firstRepo) Update(input *pbFirst.Post) (*pbFirst.EmptyResp, error) {

	result, err := fr.db.Exec(`
			UPDATE posts SET
				updated_at = CURRENT_TIMESTAMP,
				user_id = CASE WHEN $1 = 0 THEN user_id ELSE $1 END,
				title = CASE WHEN $2 = '' THEN title ELSE $2 END,
				body = CASE WHEN $3 = '' THEN body ELSE $3 END
			WHERE id = $4
		`, input.UserId, input.Title, input.Body, input.Id,
	)
	if err != nil {
		return nil, err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return nil, newerror.ErrNotFound
	}
	return &pbFirst.EmptyResp{}, nil
}
func (fr *firstRepo) List(input *pbFirst.ListReq) (*pbFirst.ListResp, error) {
	posts := make([]*pbFirst.PostResponse, 0, input.Limit)
	offset := (input.Page - 1) * input.Limit
	count := 0
	err := fr.db.QueryRow(`SELECT COUNT(*) FROM posts`).Scan(&count)
	if err != nil {
		return nil, err
	}
	rows, err := fr.db.Query(`
		SELECT * FROM posts LIMIT $1 OFFSET $2
		`, input.Limit, offset,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, newerror.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var updatedAt sql.NullString
		post := pbFirst.PostResponse{}
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Title,
			&post.Body,
			&post.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		post.UpdatedAt = updatedAt.String
		posts = append(posts, &post)
	}

	return &pbFirst.ListResp{
		Posts: posts,
		Count: int64(count),
	}, nil
}
