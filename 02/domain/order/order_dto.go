package order

type OrderReq struct {
	OrderId         int              `json:"orderId"`
	CustomerName    string           `json:"customerName"`
	OrderedAt       string           `json:"orderedAt"`
	OrderDetailList []OrderDetailReq `json:"items"`
}

func (c OrderReq) OrderReqIntoOrder() Order {
	return Order{
		OrderId:      c.OrderId,
		CustomerName: c.CustomerName,
		OrderedAt:    c.OrderedAt,
	}
}

func (c Order) OrderIntoOrderResp() OrderResp {
	return OrderResp{
		OrderId:      c.OrderId,
		CustomerName: c.CustomerName,
		OrderedAt:    c.OrderedAt,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

type OrderDetailReq struct {
	OrderDetailId int    `json:"orderDetailId"`
	ItemCode      string `json:"itemCode"`
	Description   string `json:"description"`
	Quantity      int    `json:"quantity"`
}

func (c OrderDetailReq) IntoOrderDetail() OrderDetail {
	return OrderDetail{
		OrderDetailId: c.OrderDetailId,
		ItemCode:      c.ItemCode,
		Description:   c.Description,
		Quantity:      c.Quantity,
	}
}

type OrderResp struct {
	OrderId             int               `json:"orderId"`
	CustomerName        string            `json:"customerName"`
	OrderedAt           string            `json:"orderedAt"`
	CreatedAt           string            `json:"createdAt"`
	UpdatedAt           string            `json:"updatedAt"`
	OrderDetailRespList []OrderDetailResp `json:"items"`
}

type OrderDetailResp struct {
	OrderDetailId int    `json:"orderDetailId"`
	ItemCode      string `json:"itemCode"`
	Description   string `json:"description"`
	Quantity      int    `json:"quantity"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

func (c OrderDetail) OrderDetailIntoOrderDetailResp() OrderDetailResp {
	return OrderDetailResp{
		OrderDetailId: c.OrderDetailId,
		ItemCode:      c.ItemCode,
		Description:   c.Description,
		Quantity:      c.Quantity,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}
