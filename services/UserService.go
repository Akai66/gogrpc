package services

import (
	"context"
	"time"
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

func (this *UserService) GetUserScoreByStream(userReq *UserRequest, stream UserService_GetUserScoreByStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for index,user := range userReq.Users {
		user.UserScore = score
		score++
		users = append(users,user)
		if (index + 1) % 2 == 0 {
			//每获取2个用户积分，就返回一次
			err := stream.Send(&UserResponse{Users: users})
			if err != nil {
				return err
			}
			//清空users
			users = users[0:0]
		}
		time.Sleep(1* time.Second)
	}
	//用户数量为奇数时，处理users中多出的部分
	if len(users) > 0 {
		err := stream.Send(&UserResponse{Users: users})
		if err != nil {
			return err
		}
	}
	return nil
}