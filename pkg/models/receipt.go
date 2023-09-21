package models

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Receipt struct {
	ID string
	Retailer string
	PurchaseDate string
	PurchaseTime string
	Items []*ReceiptItem
	Total string
	Points int64
}

type ReceiptItem struct {
	ShortDescription string
	Price string
}

var QuarterMultiples = []string{".25", ".50", ".75", ".00"}

func (r *Receipt) CalculatePoints() error {
	points := int64(0)
	for _, c := range r.Retailer {
		if !(unicode.IsLetter(c) || unicode.IsNumber(c)) {
			continue
		}

		points += 1
	}

	if strings.HasSuffix(r.Total, ".00") {
		points += 50
	}

	for _, multiple := range QuarterMultiples {
		if strings.HasSuffix(r.Total, multiple) {
			points += 25
			break
		}
	}

	if len(r.Items) > 0 {
		points += int64(len(r.Items) / 2 * 5)
	}

	for _, item := range r.Items {
		itemPoints, err := item.CalculatePoints()
		if err != nil {
			return err
		}

		points += itemPoints
	}

	purchaseDate, err := time.Parse("2006-01-02", r.PurchaseDate)
	if err != nil {
		return err
	}

	if purchaseDate.Day() % 2 == 1 {
		points += 6
	}

	purchaseTime, err := time.Parse("15:04", r.PurchaseTime)
	if err != nil {
		return err
	}

	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10;
	}

	r.Points = points
	return nil
}

func (ri *ReceiptItem) CalculatePoints() (int64, error) {
	trimmed := strings.TrimSpace(ri.ShortDescription)
	if len(trimmed) % 3 != 0 {
		return 0, nil
	}

	price, err := strconv.ParseFloat(ri.Price, 64)
	if err != nil {
		return 0, err
	}

	return int64(math.Ceil(price * 0.2)), nil
}
