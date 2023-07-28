package order

import (
	"context"
	"log"
)

type repository interface {
	writeOrderRepo
}

type writeOrderRepo interface {
	createOrder(ctx context.Context, order Order, orderDetailList []OrderDetail) (err error, _ Order, _ []OrderDetail)
	getOrderList(ctx context.Context) (err error, orderList []Order, orderDetailList []OrderDetail)
	editOrder(ctx context.Context, order Order, orderDetailList []OrderDetail) (err error, _ Order, _ []OrderDetail)
	removeOrder(ctx context.Context, order Order) (err error)
}

type orderService struct {
	repo repository
	r    orderRepo
}

func newOrderService(repo repository, r orderRepo) orderService {
	return orderService{
		repo: repo,
		r:    r,
	}
}

func (u orderService) createOrder(ctx context.Context, req OrderReq) (err error, resp OrderResp) {
	var order Order
	var orderDetailList []OrderDetail

	for _, orderDetail := range req.OrderDetailList {
		orderDetailList = append(orderDetailList, orderDetail.IntoOrderDetail())
	}
	err, order, orderDetailList = u.repo.createOrder(ctx, req.OrderReqIntoOrder(), orderDetailList)
	if err != nil {
		log.Println(err)
		return
	}

	resp = order.OrderIntoOrderResp()
	for _, orderDetail := range orderDetailList {
		resp.OrderDetailRespList = append(resp.OrderDetailRespList, orderDetail.OrderDetailIntoOrderDetailResp())
	}

	return err, resp
}

func (u orderService) getOrderList(ctx context.Context) (err error, respList []OrderResp) {
	err, orderList, orderDetailList := u.repo.getOrderList(ctx)
	_ = orderList
	_ = orderDetailList

	for _, order := range orderList {
		respList = append(respList, order.OrderIntoOrderResp())
	}

	for _, orderDetail := range orderDetailList {
		for idx, resp := range respList {
			if resp.OrderId == orderDetail.OrderId {
				respList[idx].OrderDetailRespList = append(respList[idx].OrderDetailRespList, orderDetail.OrderDetailIntoOrderDetailResp())
			}
		}
	}

	return err, respList
}

func (u orderService) editOrder(ctx context.Context, req OrderReq) (err error, resp OrderResp) {
	var order Order
	var orderDetailList []OrderDetail

	for _, orderDetail := range req.OrderDetailList {
		orderDetailList = append(orderDetailList, orderDetail.IntoOrderDetail())
	}
	err, order, orderDetailList = u.repo.editOrder(ctx, req.OrderReqIntoOrder(), orderDetailList)
	if err != nil {
		log.Println(err)
		return
	}

	resp = order.OrderIntoOrderResp()
	for _, orderDetail := range orderDetailList {
		resp.OrderDetailRespList = append(resp.OrderDetailRespList, orderDetail.OrderDetailIntoOrderDetailResp())
	}

	return err, resp
}

func (u orderService) removeOrder(ctx context.Context, req OrderReq) (err error) {
	err = u.repo.removeOrder(ctx, req.OrderReqIntoOrder())
	return
}
