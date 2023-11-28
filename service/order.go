package service

import "github.com/YogeLiu/medical/dao"

type OrderService struct {
	dao *dao.OrderDao
}

func NewOrderService(path string) *OrderService {
	return &OrderService{dao: dao.NewOrderDao(path)}
}

func (svc *OrderService) CreateOrder() {}
