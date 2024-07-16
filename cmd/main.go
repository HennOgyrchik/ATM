package main

import (
	"atm/repository"
	"atm/services"
	"flag"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "ip-address")
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	router := gin.Default()

	repos, err := repository.New()
	if err != nil {
		slog.Error("Creating a repository", "Error", err.Error())
		return
	}
	serv := services.New(repos)

	router.POST("/accounts", serv.CreateNewAccount)
	router.POST("/accounts/:id/deposit", serv.Deposit)
	router.POST("/accounts/:id/withdraw", serv.Withdraw)
	router.GET("/accounts/:id/balance", serv.GetBalance)

	err = router.Run(*addr)
	if err != nil {
		slog.Error("Start router", "Error", err.Error())
		return
	}
}
