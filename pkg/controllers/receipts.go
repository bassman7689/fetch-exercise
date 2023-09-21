package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/bassman7689/fetch-exercise/pkg/requests"
	"github.com/bassman7689/fetch-exercise/pkg/responses"
	"github.com/bassman7689/fetch-exercise/pkg/store"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type ReceiptsController struct {
	store store.Store
}

func handleResponse(w http.ResponseWriter, statusCode int, res any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}


func handleError(w http.ResponseWriter, statusCode int, message string) {
	handleResponse(w, statusCode, &responses.APIError{Error: message})
}

func (rc *ReceiptsController) Process(w http.ResponseWriter, r *http.Request) {
	prr := &requests.ProcessReceipt{}
	if err := json.NewDecoder(r.Body).Decode(prr); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid json in request body")
		return
	}

	if err := validate.Struct(prr); err != nil {
		// TODO: better error response
		handleError(w, http.StatusBadRequest, "Invalid fields in request json")
		return
	}

	id, err := rc.store.ProcessReceipt(prr)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error Processing Receipt")
		return
	}

	handleResponse(w, http.StatusOK, &responses.ProcessReceipt{ID: id})
}

func (rc *ReceiptsController) Points(w http.ResponseWriter, r *http.Request) {
}
