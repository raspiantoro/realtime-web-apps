/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raspiantoro/realtime-web-apps/rest/internal/app/appcontext"
	"github.com/raspiantoro/realtime-web-apps/rest/internal/app/message"
	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer"
	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer/natsstreaming"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest",
	Short: "realtime-web-apps rest",
	Long:  `rest api for realtime-web-apps`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() {
	natsHost := os.Getenv("NATS_HOST")
	natsPort, err := strconv.ParseUint(os.Getenv("NATS_PORT"), 10, 64)
	if err != nil {
		return
	}

	natsClientID := os.Getenv("NATS_CLIENT_ID")
	natsClusterID := os.Getenv("NATS_CLUSTER_ID")

	opt := streamer.Option{
		Streamer: streamer.StanStreamerType,
		StanOption: natsstreaming.ClientOption{
			ClientID:  natsClientID,
			ClusterID: natsClusterID,
			Host:      natsHost,
			Port:      natsPort,
		},
	}

	appStreamer := appcontext.Streamer{
		Stan: appcontext.InitStanStreamer(opt),
	}

	pubs := appcontext.InitPublisher(appStreamer.Stan)

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/send", func(c echo.Context) (err error) {
		min := 300
		max := 1000

		msg := message.StreamData{
			Count: rand.Intn(max-min) + min,
		}

		pubs.StreamPublisher.Publish(msg)
		log.Print("new message has been published")

		return c.String(http.StatusOK, "Data sent from rest v0.2.0")
	})

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))
}
