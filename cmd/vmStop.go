/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"Skyplatform/log"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a virtual machine",
	Long: `Stop a virtual machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmInit()
		previousState := vm.Runstate
		vm, err = vm.Stop(*client)
		if err != nil {
			log.Error(err)
		}
		log.Infof("Virtual machine state changed from (%s) to (%s): ", previousState, vm.Runstate)
	},
}

func init() {
	vmCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
