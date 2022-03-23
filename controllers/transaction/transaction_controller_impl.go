package transaction

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-tunas/customresponses"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/repositories/transactions"
	"go-tunas/services/transaction"
	"io"
	"net/http"
)

type TransactionControllerImpl struct {
	Service transaction.TransactionService
}

func NewTransactinController(db *sql.DB) *TransactionControllerImpl {
	return &TransactionControllerImpl{
		Service: transaction.TransactinoServiceImpl{
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
