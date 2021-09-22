package services

import (
	"context"
	"io"
	"time"
)

// UserService 实现UserServiceServer接口
type UserService struct {

}

// GetUserScore 普通模式
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

// GetUserScoreByServerStream 服务端流模式，服务端按批次返回数据给客户端
func (this *UserService) GetUserScoreByServerStream(userReq *UserRequest, stream UserService_GetUserScoreByServerStreamServer) error {
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

// GetUserScoreByClientStream 客户端流模式，客户端按批次将请求数据发送给服务端
func (this *UserService) GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for {
		//分批接收
		userReq,err := stream.Recv()
		if err == io.EOF {
			//客户端发送完毕，服务端一次性响应数据给客户端
			return stream.SendAndClose(&UserResponse{Users: users})
		}
		if err != nil {
			return err
		}
		//处理客户端分批发送过来的数据，类似业务处理逻辑
		for _,user := range userReq.Users {
			user.UserScore = score
			score++
			users = append(users,user)
		}
	}
}

