package services

import (
	"context"
	"fmt"
)

// OrderService 实现OrderServiceServer接口
type OrderService struct {

}


func (this *OrderService) NewOrder(ctx context.Context,order *OrderMain) (*OrderResponse, error) {
	fmt.Println(order)
	return &OrderResponse{
		Status: "ok",
		Msg: "success",
	},nil
}
