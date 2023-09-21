package models

type Receipt struct {
	ID string
	Retailer string
	PurchaseDate string
	PurchaseTime string
	Items []*ReceiptItem
	Total string
}

type ReceiptItem struct {
	ShortDescription string
	Price string
}
