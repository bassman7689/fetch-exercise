package requests

type ProcessReceipt struct {
	Retailer string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items []*ProcessReceiptItem `json:"items" validate:"required,dive,required"`
	Total string `json:"total" validate:"required"`
}

type ProcessReceiptItem struct {
	ShortDescription string  `json:"shortDescription" validate:"required"`
	Price string `json:"price" validate:"required"`
}
