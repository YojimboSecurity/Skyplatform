/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"Skyplatform/log"

	"github.com/spf13/cobra"
)

// suspendCmd represents the suspend command
var suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "Suspend a virtual machine",
	Long: `Suspend a virtual machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmInit()
		previousState := vm.Runstate
		vm, err = vm.Suspend(*client)
		if err != nil {
			log.Error(err)
		}
		log.Infof("Virtual machine state changed from (%s) to (%s): ", previousState, vm.Runstate)
	},
}

func init() {
	vmCmd.AddCommand(suspendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suspendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// suspendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
