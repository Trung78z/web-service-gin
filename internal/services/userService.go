package services

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/Trung78z/web-service-gin/internal/repositories/queries"
	"github.com/Trung78z/web-service-gin/pkg/utils"
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
	cacheKey := "user:" + strconv.Itoa(int(id))
	val, err := utils.GetCache(ctx, cacheKey)
	if err == nil {
		// If cache hit, unmarshal the value
		var user queries.User
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			return queries.User{}, errors.New("failed to unmarshal user data from cache")
		}

		return user, nil
	}

	// If cache miss, get user from database
	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		return queries.User{}, err
	}

	// Set the user data into cache for future requests
	userData, _ := json.Marshal(user)
	err = utils.SetCache(ctx, cacheKey, string(userData), 3600) // Cache for 1 hour
	if err != nil {
		log.Printf("Failed to set cache for user %d: %v", id, err)
	}

	return user, nil
}

// ListUsers lấy danh sách tất cả người dùng
func (s *UserService) ListUsers(ctx context.Context) ([]queries.User, error) {
	users, err := s.queries.ListUsers(ctx)
	return users, err
}
