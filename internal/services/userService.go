package services

import (
	"context"
	"github/Trung78z/web-service-gin/internal/repositories/queries"
	"log"
)

type UserService struct {
	queries *queries.Queries
}

// NewUserService tạo instance của UserService và inject đối tượng queries
func NewUserService(q *queries.Queries) *UserService {
	return &UserService{
		queries: q,
	}
}

// CreateUser tạo người dùng mới
func (s *UserService) CreateUser(ctx context.Context, name string, email string) (queries.User, error) {
	user, err := s.queries.CreateUser(ctx, queries.CreateUserParams{
		Name:  name,
		Email: email,
	})
	return user, err
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (queries.User, error) {
	log.Println(id)
	user, err := s.queries.GetUserByID(ctx, id)
	return user, err
}

// ListUsers lấy danh sách tất cả người dùng
func (s *UserService) ListUsers(ctx context.Context) ([]queries.User, error) {
	users, err := s.queries.ListUsers(ctx)
	return users, err
}
