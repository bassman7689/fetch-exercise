package store

import (
	"sync"

	"github.com/google/uuid"

	"github.com/bassman7689/fetch-exercise/pkg/models"
	"github.com/bassman7689/fetch-exercise/pkg/requests"
)

type Store interface {
	ProcessReceipt(prr *requests.ProcessReceipt) (string, error)
	GetReceiptById(id string) (*models.Receipt, error)
}

type MemoryStore struct {
	receipts map[string]*models.Receipt
	sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	receipts := make(map[string]*models.Receipt)
	return &MemoryStore{receipts: receipts}
}

func (ms *MemoryStore) ProcessReceipt(prr *requests.ProcessReceipt) (string, error) {
	ms.Lock()
	defer ms.Unlock()

	items := make([]*models.ReceiptItem, 0, len(prr.Items))
	for _, item := range prr.Items {
		items = append(items, &models.ReceiptItem {
			ShortDescription: item.ShortDescription,
			Price: item.Price,
		})
	}

	id := uuid.NewString()
	_, found := ms.receipts[id];
	for  found {
		id := uuid.NewString()
		_, found = ms.receipts[id];
	}

	receipt := &models.Receipt{
		ID: id,
		Retailer: prr.Retailer,
		PurchaseDate: prr.PurchaseDate,
		PurchaseTime: prr.PurchaseTime,
		Items: items,
		Total: prr.Total,
	}

	receipt.CalculatePoints()
	ms.receipts[id] = receipt

	return id, nil
}

func (ms *MemoryStore) GetReceiptById(id string) (*models.Receipt, error) {
	ms.RLock()
	defer ms.RUnlock()

	receipt, ok := ms.receipts[id]
	if !ok {
		return nil, nil
	}

	return receipt, nil
}
