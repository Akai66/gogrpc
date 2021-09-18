package services

import (
	"context"
	"fmt"
)

// OrderService 实现OrderServiceServer接口
type OrderService struct {

}


func (this *OrderService) NewOrder(ctx context.Context,orderReq *OrderRequest) (*OrderResponse, error) {
	fmt.Println(orderReq.OrderMain)
	return &OrderResponse{
		Status: "ok",
		Msg: "success",
	},nil
}
