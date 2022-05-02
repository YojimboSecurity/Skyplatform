/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"Skyplatform/pkg"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/YojimboSecurity/skytap-sdk-go/api"
	bubbleSpin "github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var configurationId string

// environmentCmd represents the environment command
var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Display environment information",
	Long:  `Display environment information.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configurationId == "" {
			fmt.Println("Please provide a configuration id")
			return
		}
		resp := getEnvironmentInfo(configurationId)
		displayEnvironmentInfo(resp)
	},
}

func init() {
	rootCmd.AddCommand(environmentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// environmentCmd.PersistentFlags().String("foo", "", "A help for foo")
	environmentCmd.Flags().StringVarP(&configurationId, "configuration-id", "c", "", "configuration id")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// environmentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// getEnvironmentInfo gets the environment information
func getEnvironmentInfo(configurationId string) interface{} {
	URL := fmt.Sprintf("https://cloud.skytap.com/v2/configurations/%s.json", configurationId)
	p := tea.NewProgram(pkg.SpinnerModel(bubbleSpin.Pulse, "205"))
	client := pkg.NewClient()
	resp := interface{}(nil)
	go func() {
		api.GetSkytapResource(*client, URL, &resp)
		p.Quit()

	}()
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

// displayEnvironmentInfo displays the environment information
func displayEnvironmentInfo(resp interface{}) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Key\tValue\n")
	fmt.Fprintf(tw, "----\t------\n")
	for k, v := range resp.(map[string]interface{}) {
		if k == "vms" {
			fmt.Fprintln(tw, "vms\t", len(v.([]interface{})))
			fmt.Fprintf(tw, "\tKey\tValue\n")
			fmt.Fprintf(tw, "\t---\t-----\n")
			for _, v := range v.([]interface{}) {
				for kk, vv := range v.(map[string]interface{}) {
					fmt.Fprintf(tw, "\t%s\t%v\n", kk, vv)
				}
				fmt.Fprintln(tw, "\t")
			}
			continue
		}
		if k == "networks" {
			fmt.Fprintln(tw, "Networks\t", len(v.([]interface{})))
			fmt.Fprintf(tw, "\tKey\tValue\n")
			fmt.Fprintf(tw, "\t---\t-----\n")
			for _, v := range v.([]interface{}) {
				for kk, vv := range v.(map[string]interface{}) {
					fmt.Fprintf(tw, "\t%s\t%v\n", kk, vv)
				}
				fmt.Fprintln(tw, "\t")
			}
			continue
		}
		fmt.Fprintf(tw, "%s\t%v\n", k, v)
	}
	tw.Flush()
}
