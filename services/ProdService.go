package services

import "context"

// ProdService 定义ProdService用于实现ProdServiceServer接口
type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error)  {
	return &ProdResponse{ProdStock: 20},nil
}
