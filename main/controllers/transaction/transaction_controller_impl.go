package transaction

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/transactions"
	transaction2 "github.com/ramdanariadi/grocery-be-golang/main/services/transaction"
	"io"
	"net/http"
)

type TransactionControllerImpl struct {
	Service transaction2.TransactionService
}

func NewTransactionController(db *sql.DB) *TransactionControllerImpl {
	return &TransactionControllerImpl{
		Service: transaction2.TransactinoServiceImpl{
			Repository: transactions.TransactionRepositoryImpl{
				DB: db,
			},
		},
	}
}

func (controller TransactionControllerImpl) FindByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responses := controller.Service.FindByUserId(id)
	customresponses.SendResponse(w, responses, http.StatusOK)
}

func (controller TransactionControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responses := controller.Service.FindByTransactionId(id)
	if responses.Id == "" {
		customresponses.SendResponse(w, nil, http.StatusNoContent)
		return
	}
	customresponses.SendResponse(w, responses, http.StatusOK)
}

func (controller TransactionControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	reader := r.Body
	bytes, err := io.ReadAll(reader)
	helpers.PanicIfError(err)
	transactionModel := models.TransactionModel{}
	err = json.Unmarshal(bytes, &transactionModel)
	helpers.PanicIfError(err)
	code := http.StatusInternalServerError

	if controller.Service.Save(transactionModel) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller TransactionControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	code := http.StatusNotModified
	if controller.Service.Delete(id) {
		code = http.StatusOK
	}
	customresponses.SendResponse(w, nil, code)
}
