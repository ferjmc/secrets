package cobra

import (
	"fmt"

	"github.com/ferjmc/secrets"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a secret",
	Run: func(cmd *cobra.Command, args []string) {
		v := secrets.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("no value set")
		}
		fmt.Printf("%s = %s", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
