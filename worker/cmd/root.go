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
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/raspiantoro/realtime-web-apps/worker/internal/app/appcontext"
	"github.com/raspiantoro/realtime-web-apps/worker/pkg/streamer"
	"github.com/raspiantoro/realtime-web-apps/worker/pkg/streamer/natsstreaming"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "worker",
	Short: "realtime-web-apps worker",
	Long:  `worker apps for realtime-web-apps`,
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	opt := streamer.Option{
		Streamer: streamer.StanStreamerType,
		StanOption: natsstreaming.ClientOption{
			ClientID:  "streamappworker",
			ClusterID: "streamapps",
			Host:      "localhost",
			Port:      4222,
		},
	}

	appStreamer := appcontext.Streamer{
		Stan: appcontext.InitStanStreamer(opt),
	}

	go func() {
		select {
		case <-c:
			log.Println("shutting down application...")
			cancel()
			appStreamer.Stan.Close()
		}
	}()

	pubs := appcontext.InitPublisher(appStreamer.Stan)

	task := appcontext.InitTask(pubs)

	streamTask := task.StreamTask

	log.Println("tasks is running")
	streamTask.Run(ctx)

}
