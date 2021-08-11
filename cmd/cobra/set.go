package cobra

import (
	"fmt"

	"github.com/ferjmc/secrets"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a secret",
	Run: func(cmd *cobra.Command, args []string) {
		v := secrets.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s=%s sets successfully", key, value)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
