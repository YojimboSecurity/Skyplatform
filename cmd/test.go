/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"Skyplatform/pkg"
	"Skyplatform/log"
	"fmt"

	"github.com/skytap/skytap-sdk-go/api"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		client := pkg.NewClient()
		//resp := interface{}(nil)
		env ,err := api.GetEnvironment(*client, "122243460")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(env.Runstate)
		vm, err := api.GetVirtualMachine(*client, "90540245")
		if err != nil {
			log.Error(err)
		}
		fmt.Println(vm.Runstate)
		vm, err = vm.Stop(*client)
		if err != nil {
			log.Error(err)
		}
		fmt.Println(vm.Runstate)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
