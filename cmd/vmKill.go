/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"Skyplatform/log"

	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill a virtual machine",
	Long: `Kill a virtual machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmInit()
		previousState := vm.Runstate
		vm, err = vm.Kill(*client)
		if err != nil {
			log.Error(err)
		}
		log.Infof("Virtual machine state changed from (%s) to (%s): ", previousState, vm.Runstate)
	},
}

func init() {
	vmCmd.AddCommand(killCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// killCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// killCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
