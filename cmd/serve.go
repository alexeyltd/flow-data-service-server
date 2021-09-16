package cmd

import (
	"context"
	"flow-data-service-server/internal/controller"
	"flow-data-service-server/internal/infra"
	"flow-data-service-server/internal/repository"
	"flow-data-service-server/internal/service"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := infra.NewConfig()
		if err != nil {
			panic(err)
		}
		gorm, err := infra.NewGorm(config)
		if err != nil {
			panic(err)
		}
		l, err := infra.NewLogger(config)
		if err != nil {
			panic(err)
		}
		graphRepositoryImpl := repository.NewGraphRepositoryImpl(gorm, l.Sugar())
		graphService := service.NewGraphServiceImpl(graphRepositoryImpl)
		userController := controller.NewGraphController(graphService)

		engine, err := controller.NewGin(userController)
		if err != nil {
			panic(err)
		}

		srv := &http.Server{
			Addr:    config.HttpAddr,
			Handler: engine,
		}

		// gracefull
		go func() {
			quit := make(chan os.Signal)
			signal.Notify(quit, os.Interrupt)
			<-quit
			log.Println("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			log.Println("Server exiting")
		}()

		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
