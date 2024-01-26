package decrypt

import (
	"github.com/nomionz/fs-enc/pkg/fs"
	"github.com/nomionz/fs-enc/pkg/util"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dec [file path] [passphrase]",
		Short: "Decrypt the file",
		Long:  `Decrypt the file at the given path using the provided passphrase.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			passphrase := args[1]

			util.ExitOnError(fs.Decrypt(path, []byte(passphrase)))
		},
	}

	return cmd
}
