
package cmd

import (
	"GOLAND/database"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fullname := cmd.Flag("fullname").Value.String()
		database.CreatePerson(fullname)
	},
}

func init() {
	createCmd.Flags().String("fullname","","create person")
	createCmd.MarkFlagRequired("fullname")
	rootCmd.AddCommand(createCmd)

}

