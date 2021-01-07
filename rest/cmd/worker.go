/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os"
	"os/signal"

	"github.com/raspiantoro/realtime-web-apps/rest/internal/app/appcontext"
	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer"
	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer/natsstreaming"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runWorker()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runWorker() {

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
