// Package cmd /*
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/HeRoMo/toggl-cli/toggl"
	"github.com/spf13/cobra"
)

// reportsCmd represents the reports command
var reportsCmd = &cobra.Command{
	Use:   "reports",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		togglToken := os.Getenv("TOGGL_TOKEN")
		workspaceID, _ := cmd.Flags().GetInt("workspace")
		since, _ := cmd.Flags().GetString("since")
		until, _ := cmd.Flags().GetString("until")

		ctx, cancelFunc := context.WithCancel(context.Background())
		togglClient := toggl.Client(togglToken)
		report, err := togglClient.Reports(ctx, workspaceID, since, until)
		if err != nil {
			log.Fatal(err)
		}

		cancelFunc()

		fmt.Printf("%+v\n", report)
	},
}

func init() {
	rootCmd.AddCommand(reportsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//reportsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	var workspace int
	var since string
	var until string
	reportsCmd.Flags().IntVar(&workspace, "workspace", 0, "Workspace ID")
	reportsCmd.Flags().StringVar(&since, "since", "", "Since")
	reportsCmd.Flags().StringVar(&until, "until", "", "Until")
}
