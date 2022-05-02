/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"Skyplatform/log"
	"Skyplatform/pkg"
	"fmt"

	"github.com/YojimboSecurity/skytap-sdk-go/api"
	bubbleSpin "github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var vmId string
var vm *api.VirtualMachine
var err error
var client *api.SkytapClient

// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Get virtual machine information",
	Long: `Get virtual machine information`,
	Run: func(cmd *cobra.Command, args []string) {
		vmInit()
		fmt.Printf("%s is %s\n\n", vm.Name, vm.Runstate)
		for i := range vm.Interfaces {
			for ps := range vm.Interfaces[i].PublishedServices {
				fmt.Println("Published Service:")
				fmt.Printf("\tId:%s\n", vm.Interfaces[i].PublishedServices[ps].Id)
				fmt.Printf("\tInternal Port:%d\n", vm.Interfaces[i].PublishedServices[ps].InternalPort)
				fmt.Printf("\tExternal IP:%s\n", vm.Interfaces[i].PublishedServices[ps].ExternalIp)
				fmt.Printf("\tExternal Port:%d\n", vm.Interfaces[i].PublishedServices[ps].ExternalPort)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(vmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmCmd.PersistentFlags().String("foo", "", "A help for foo")
	vmCmd.PersistentFlags().StringVarP(&vmId, "vm-id", "v", "", "Virtual machine Id")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func vmInit() {
	if vmId == "" {
		log.Error("Virtual machine Id is required")
		return
	}
	p := tea.NewProgram(pkg.SpinnerModel(bubbleSpin.Pulse, "205"))
	client = pkg.NewClient()
	go func() {
		vm, err = api.GetVirtualMachine(*client, vmId)
		if err != nil {
			log.Error(err)
		}
		p.Quit()
	}()
	if err := p.Start(); err != nil {
		log.Error(err)
		return
	}
}
