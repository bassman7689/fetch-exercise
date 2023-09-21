package models

import (
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tt := []struct{
		receipt *Receipt
		expectedPoints int64
	}{
		{
			receipt: &Receipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*ReceiptItem {
					&ReceiptItem {
						ShortDescription: "Mountain Dew 12PK",
						Price: "6.49",
					},
					&ReceiptItem{
						ShortDescription: "Emils Cheese Pizza",
						Price: "12.25",
					},
					&ReceiptItem{
						ShortDescription: "Knorr Creamy Chicken",
						Price: "1.26",
					},
					&ReceiptItem{
						ShortDescription: "Doritos Nacho Cheese",
						Price: "3.35",
					},
					&ReceiptItem{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price: "12.00",
					},
				},
				Total: "35.35",
			},
			expectedPoints: 28,
		},
		{
			receipt: &Receipt{
				Retailer: "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []*ReceiptItem{
					&ReceiptItem{
						ShortDescription: "Gatorade",
						Price: "2.25",
					},
					&ReceiptItem{
						ShortDescription: "Gatorade",
						Price: "2.25",
					},
					&ReceiptItem{
						ShortDescription: "Gatorade",
						Price: "2.25",
					},
					&ReceiptItem{
						ShortDescription: "Gatorade",
						Price: "2.25",
					},
				},
				Total: "9.00",
			},
			expectedPoints: 109,
		},
	}

	for _, test := range tt {
		if err := test.receipt.CalculatePoints(); err != nil {
			t.Errorf("error while calculating points %v", err)
		}

		if test.receipt.Points != test.expectedPoints {
			t.Errorf("wrong number of points: got %v want %v", test.receipt.Points, test.expectedPoints)
		}
	}

}
