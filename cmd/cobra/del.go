package cobra

import (
	"fmt"

	"github.com/ferjmc/secrets"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a secret by his key",
	Run: func(cmd *cobra.Command, args []string) {
		v := secrets.File(encodingKey, secretsPath())
		key := args[0]
		err := v.Del(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s delete successfully", key)
	},
}

func init() {
	RootCmd.AddCommand(delCmd)
}
