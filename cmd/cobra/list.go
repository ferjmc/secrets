package cobra

import (
	"fmt"

	"github.com/ferjmc/secrets"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the secrets keys",
	Run: func(cmd *cobra.Command, args []string) {
		v := secrets.File(encodingKey, secretsPath())
		value, err := v.List()
		if err != nil {
			fmt.Println("no secrets availables")
		}
		for _, v := range value {
			fmt.Println(v)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
