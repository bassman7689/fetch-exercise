package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/bassman7689/fetch-exercise/pkg/requests"
	"github.com/bassman7689/fetch-exercise/pkg/responses"
	"github.com/bassman7689/fetch-exercise/pkg/store"
)

func TestProcessReceipt(t *testing.T) {
	tt := []struct{
		name string
		request *requests.ProcessReceipt
		expectedStatusCode int
	}{
		{
			name: "valid request",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "10.00",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Empty Retailer",
			request: &requests.ProcessReceipt{
				Retailer: "",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "10.00",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Empty PurchaseDate",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "",
				PurchaseTime: "13:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "10.00",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Empty PurchaseTime",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "10.00",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Empty Items",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "",
				Items: []*requests.ProcessReceiptItem{
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Empty Total",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "10.00",
					},
				},
				Total: "",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid PurchaseDate",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-13-01",
				PurchaseTime: "13:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "10.00",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid PurchaseTime",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "25:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "10.00",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid Item ShortDescription",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "",
						Price: "10.00",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid Item Price",
			request: &requests.ProcessReceipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*requests.ProcessReceiptItem{
					&requests.ProcessReceiptItem{
						ShortDescription: "A Product",
						Price: "",
					},
				},
				Total: "10.00",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}


	r := mux.NewRouter()
	Register(r, store.NewMemoryStore())
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			reqBody, err := json.Marshal(test.request)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewReader(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			bodyBytes := rr.Body.Bytes()
			if rr.Code != test.expectedStatusCode {
				t.Errorf("handler returned incorrect status code: got %v want %v", rr.Code, test.expectedStatusCode)
				t.Errorf("response body: %v", string(bodyBytes))
			}

			if test.expectedStatusCode != http.StatusOK {
				return
			}

			res := &responses.ProcessReceipt{}
			if err := json.Unmarshal(bodyBytes, &res); err != nil {
				t.Errorf("handler returned invalid json: got %v", rr.Body.String())
			}

			if err := validate.Var(res.ID, "required,uuid"); err != nil {
				t.Errorf("invalid id %q: %v", res.ID, err)
			}
		})
	}
}
