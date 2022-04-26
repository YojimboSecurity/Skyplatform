/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"Skyplatform/pkg"

	bubbleSpin "github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/skytap/skytap-sdk-go/api"
	"github.com/spf13/cobra"
)


// environmentsCmd represents the environments command
var environmentsCmd = &cobra.Command{
	Use:   "environments",
	Short: "List all virtual environments",
	Long: `List all available virtual environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(pkg.SpinnerModel(bubbleSpin.Pulse, "205"))
		client := pkg.NewClient()
		resp := interface{}(nil)
		go func() {
			api.GetSkytapResource(*client, "https://cloud.skytap.com/v2/configurations?scope=company&count=40", &resp)
			p.Quit()

		}()
		if err := p.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, v := range resp.([]interface{}) {
			name := v.(map[string]interface{})["name"]
			id := v.(map[string]interface{})["id"]
			fmt.Println(id, name)
		}
	},
}

func init() {
	rootCmd.AddCommand(environmentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// environmentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// environmentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
