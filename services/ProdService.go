package services

import (
	"context"
)

// ProdService 定义ProdService用于实现ProdServiceServer接口
type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error)  {
	var stock int32 = 50
	if request.ProdArea == ProdAreas_A {
		stock = 41
	}else if request.ProdArea == ProdAreas_B {
		stock = 30
	}else if request.ProdArea == ProdAreas_C {
		stock = 20
	}
	return &ProdResponse{ProdStock: stock},nil
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