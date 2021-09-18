package services

import (
	"context"
)

// ProdService 定义ProdService用于实现ProdServiceServer接口
type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error)  {
	return &ProdResponse{ProdStock: 25},nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, in *QuerySize) (*ProdResponseList, error)  {
	prodres := []*ProdResponse{
		{ProdStock: 25}, //此处编译器会自动转换为指针类型
		{ProdStock: 26},
		{ProdStock: 27},
		{ProdStock: 28},
	}
	return &ProdResponseList{Prods:prodres},nil
}