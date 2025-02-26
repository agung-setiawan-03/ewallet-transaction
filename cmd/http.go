package cmd

import (
	"ewallet-transaction/external"
	"ewallet-transaction/helpers"
	"ewallet-transaction/internal/api"
	"ewallet-transaction/internal/interfaces"
	"ewallet-transaction/internal/repository"
	"ewallet-transaction/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	r.GET("/healthcheck", d.HealthcheckAPI.HelathcheckHandlerHTTP)

	transactionV1 := r.Group("/transaction/v1")
	transactionV1.POST("/create", d.MiddlewareValidateToken, d.TransactionAPI.CreateTransaction)
	transactionV1.PUT("/update-status/:reference", d.MiddlewareValidateToken, d.TransactionAPI.UpdateStatusTransaction)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckAPI
	External       interfaces.IExternal
	TransactionAPI interfaces.ITransactionAPI
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HelathcheckServices: healthcheckSvc,
	}

	external := &external.External{}

	trxRepo := &repository.TransactionRepo{
		DB: helpers.DB,
	}
	trxService := &services.TransactionService{
		TransactionRepo: trxRepo,
		External:        external,
	}
	trxAPI := &api.TransactionAPI{
		TransactionService: trxService,
	}
	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		External:       external,
		TransactionAPI: trxAPI,
	}

}
