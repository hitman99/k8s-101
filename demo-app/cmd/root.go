//go:generate goagen bootstrap -d github.com/hitman99/k8s-101/demo-app/design
package cmd

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hitman99/k8s-101/demo-app/app"
	"github.com/hitman99/k8s-101/demo-app/controllers"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "demo",
	Run: func(cmd *cobra.Command, args []string) {
		// Create service
		service := goa.New("demo")

		// Mount middleware
		service.Use(middleware.RequestID())
		service.Use(middleware.LogRequest(true))
		service.Use(middleware.ErrorHandler(service, true))
		service.Use(middleware.Recover())

		// Mount controllers
		c := controllers.NewHealthController(service)
		app.MountHealthController(service, c)
		f := controllers.NewPublicController(service)
		app.MountPublicController(service, f)

		// Start service
		if err := service.ListenAndServe(":80"); err != nil {
			service.LogError("startup", "err", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
