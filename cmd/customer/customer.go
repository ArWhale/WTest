package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/SArtemJ/WTest/internal/adapters"
	"github.com/SArtemJ/WTest/internal/config"
	"github.com/SArtemJ/WTest/internal/consts"
	"github.com/SArtemJ/WTest/internal/customer"
	"github.com/SArtemJ/WTest/internal/server"
	"github.com/SArtemJ/WTest/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	viper, err := config.NewConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger(viper)
	RunServer(viper, logger)
}

func RunServer(viper *viper.Viper, logger logrus.FieldLogger) {

	//custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("genderCustom", customer.GenderValidation())
		v.RegisterValidation("birthdateCustom", customer.BirthDateValidation())
	}

	r, err := adapters.NewRepositories(viper.GetString(consts.ServiceDBUrlKey))
	if err != nil {
		logger.Panic(err)
	}
	defer r.Close()

	ch := server.NewCustomerHandlers(r.CustomerRepository, logger)

	router := gin.Default()
	router.LoadHTMLGlob("./web/templates/*")
	server.InitServiceRoutes(router, ch)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%v", viper.GetString(consts.ServiceHostKey), viper.GetInt64(consts.ServicePortKey)),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Infof("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server shut down")
}
