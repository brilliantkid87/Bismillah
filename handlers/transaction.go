package handlers

// import (
// 	"net/http"
// 	dto "waysbean/dto/result"
// 	transactiondto "waysbean/dto/transaction"
// 	"waysbean/models"
// 	"waysbean/repositories"

// 	"github.com/labstack/echo/v4"
// )

// type handlerTransaction struct {
// 	TransactionRepository repositories.TransactionRepository
// }

// func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
// 	return &handlerTransaction{TransactionRepository}
// }

// func (h *handlerTransaction) FindTransaction(c echo.Context) error {
// 	transactions, err := h.TransactionRepository.FindTransaction()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
// }

// func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
// 	request := new(transactiondto.CreateTransactionRequest)
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	transaction := models.Transaction{
// 		Name:       request.Name,
// 		Email:      request.Email,
// 		Address:    request.Address,
// 		Phone:      request.Phone,
// 		Status:     "waiting",
// 	}

// 	data, err := h.TransactionRepository.CreateTransaction(transaction)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})

// }

// func convertResponseTransaction(u models.Transaction) transactiondto.TransactionResponse {
// 	return transactiondto.TransactionResponse{
// 		UserID:     u.UserID,
// 		Name:       u.Name,
// 		Phone:      u.Phone,
// 		Address:    u.Address,
// 		Status:     u.Status,
// 		TotalPrice: u.TotalPrice,
// 		ProductID:  u.ProductID,
// 	}
// }
