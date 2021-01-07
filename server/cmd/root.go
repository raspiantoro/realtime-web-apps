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
	"strconv"

	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer"
	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer/natsstreaming"
	"github.com/raspiantoro/realtime-web-apps/server/internal/app/appcontext"
	"github.com/raspiantoro/realtime-web-apps/server/internal/app/server"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "realtime-web-apps server",
	Long:  `server apps for realtime-web-apps`,
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

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".server" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run() {
	port := os.Getenv("GRPC_PORT")
	host := os.Getenv("GRPC_HOST")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	natsHost := os.Getenv("NATS_HOST")
	natsPort, err := strconv.ParseUint(os.Getenv("NATS_PORT"), 10, 64)
	if err != nil {
		return
	}

	opt := streamer.Option{
		Streamer: streamer.StanStreamerType,
		StanOption: natsstreaming.ClientOption{
			ClientID:  "streamappserver",
			ClusterID: "streamapps",
			Host:      natsHost,
			Port:      natsPort,
		},
	}

	stremerOption := appcontext.StreamerOption{
		Stan: opt,
	}

	appStreamer := appcontext.InitStreamer(stremerOption)

	go func() {
		select {
		case <-c:
			log.Println("shutting down application...")
			cancel()
			appStreamer.Stan.Close()
		}
	}()

	consumer := appcontext.InitConsumer(appStreamer.Stan)

	consumer.Stream.Start(ctx)

	svc := appcontext.InitService(ctx)
	srv := server.NewServer()

	err = srv.RunServer(ctx, svc, host, port)
	if err != nil {
		log.Fatalf("failed to start gRPC server: %s", err)
	}
}
