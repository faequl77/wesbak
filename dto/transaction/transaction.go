package transactiondto

type CheckoutRequest struct {
	Name          string `json:"name" form:"name"`
	Email         string `json:"email" form:"email"`
	RecepientName string `json:"recepient_name" form:"recepient_name"`
	Phone         string `json:"phone" form:"phone"`
	PosCode       string `json:"pos_code" form:"pos_code"`
	Address       string `json:"address" form:"address"`
	TotalPrice    int    `json:"total_price" form:"total_price"`
	Status        string `json:"status" form:"status"`
	CartID        []int  `json:"cart_id" form:"cart_id"`
}

type DeleteTransactionRequest struct {
	ID []int `json:"id" form:"id"`
}
