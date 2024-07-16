package services

import (
	"atm/models"
	"atm/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

var errParseID = errors.New("error parse ID")

type Service struct {
	repository repository.Repository
}

func New(storage repository.Repository) *Service {
	return &Service{repository: storage}
}

func (s *Service) CreateNewAccount(c *gin.Context) {
	id, err := uuid.NewUUID()
	if err != nil {
		slog.Error("Method CreateNewAccount", "Error", err.Error())

		c.JSON(http.StatusInternalServerError, struct {
			Error string
		}{err.Error()})
		return
	}

	_, ok := s.repository.GetAccount(id)
	if ok {
		s.CreateNewAccount(c)
		return
	}

	err = s.repository.AddAccount(models.NewAccount(id))

	if err != nil {
		slog.Error("Method CreateNewAccount", "Error", err.Error())

		c.JSON(http.StatusInternalServerError, struct {
			Error string
		}{err.Error()})
		return
	}

	slog.Info("A new account has been created", "id", id)
	c.JSON(http.StatusCreated, struct {
		ID uuid.UUID
	}{id})
}

func (s *Service) Deposit(c *gin.Context) {
	acc, err := s.getAccount(c)
	if errors.Is(err, errParseID) {
		return
	}

	req := models.Amount{}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		slog.Error("Method Deposit", "Error", err.Error())
		c.JSON(http.StatusBadRequest, struct {
			Error string
		}{err.Error()})
		return
	}

	err = acc.Deposit(req.Amount)
	if err != nil {
		slog.Error("Method Deposit", "Error", err.Error())
		c.JSON(http.StatusInternalServerError, struct {
			Error string
		}{err.Error()})
		return
	}
	slog.Info("Deposit success", "id", acc.GetID())
}
func (s *Service) Withdraw(c *gin.Context) {
	acc, err := s.getAccount(c)
	if errors.Is(err, errParseID) {
		return
	}

	req := models.Amount{}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		slog.Error("Method Withdraw", "Error", err.Error())
		c.JSON(http.StatusBadRequest, struct {
			Error string
		}{err.Error()})
		return
	}

	err = acc.Withdraw(req.Amount)
	if err != nil {
		slog.Error("Method Withdraw", "Error", err.Error())
		c.JSON(http.StatusBadRequest, struct {
			Error string
		}{err.Error()})
		return
	}

	slog.Info("Withdraw success", "id", acc.GetID())

}
func (s *Service) GetBalance(c *gin.Context) {
	acc, err := s.getAccount(c)
	if errors.Is(err, errParseID) {
		return
	}

	balance := acc.GetBalance()

	c.JSON(http.StatusOK, struct {
		Balance float64
	}{balance})
	slog.Info("GetBalance success", "id", acc.GetID())
}

func (s *Service) getAccount(c *gin.Context) (models.BankAccount, error) {
	idStr, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, struct {
			Error string
		}{"Bad request"})
		return nil, errParseID
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		slog.Error("Method GetBalance", "Error", err.Error())
		c.JSON(http.StatusBadRequest, struct {
			Error string
		}{err.Error()})
		return nil, errParseID
	}

	acc, ok := s.repository.GetAccount(id)
	if !ok {
		c.JSON(http.StatusNotFound, struct {
			Error string
		}{"Account not found"})
		return nil, errParseID
	}
	return acc, nil
}
