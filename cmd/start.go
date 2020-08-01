package cmd

import (
	"context"
	"ecommerce/logger"
	"ecommerce/routes"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start magellan service",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		e.Pre(middleware.RemoveTrailingSlash())
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: logger.MiddlewareLog,
		}))
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}))

		// index route
		e.GET("/", func(c echo.Context) error {
			message := `This service called Magellan`
			return c.String(http.StatusOK, message)
		})

		// Main Routes
		routes.Employee(e)
		routes.User(e)

		// Start server
		// e.Logger.Fatal(e.Start(":8000"))
		go func() {
			if err := e.Start(":8000"); err != nil {
				e.Logger.Info("Shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}
