package services

import (
	"context"
)

// UserService 实现UserServiceServer接口
type UserService struct {

}


func (this *UserService) GetUserScore(ctx context.Context, userReq *UserRequest) (*UserResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for _,user := range userReq.Users {
		user.UserScore = score
		score++
		users = append(users,user)
	}
	return &UserResponse{Users: users},nil
}
