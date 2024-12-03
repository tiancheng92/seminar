package main

import (
	"github.com/spf13/cobra"
	"github.com/tiancheng92/seminar/config"
	"github.com/tiancheng92/seminar/pkg/log"
	"github.com/tiancheng92/seminar/pkg/validator"
	"github.com/tiancheng92/seminar/server"
	"github.com/tiancheng92/seminar/store"
	_ "go.uber.org/automaxprocs"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "seminar",
	Version: "v0.0.1",
	Short:   "seminar",
	Long:    "seminar",
	Run: func(_ *cobra.Command, _ []string) {
		server.Run()
	},
}

func init() {
	config.Init()
	log.Init()
	validator.Init()
	store.Init()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
