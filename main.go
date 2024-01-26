package main

import (
	"github.com/nomionz/fs-enc/cmd/decrypt"
	"github.com/nomionz/fs-enc/cmd/encrypt"
	"github.com/nomionz/fs-enc/pkg/util"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "",
	}

	cmd.CompletionOptions.HiddenDefaultCmd = true

	cmd.AddCommand(encrypt.Cmd(), decrypt.Cmd())

	util.ExitOnError(cmd.Execute())
}
