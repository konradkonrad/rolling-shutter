package main

import (
	"github.com/spf13/cobra"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/chain"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/gnoshkeyper"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/rootcmd"
)

func subcommands() []*cobra.Command {
	return []*cobra.Command{
		chain.Cmd(),
		gnoshkeyper.Cmd(),
	}
}

func cmd() *cobra.Command {
	cmd := rootcmd.Cmd()
	cmd.Use = "gnosh"
	cmd.Short = "gnosh - gnosis shutterized chain"
	cmd.AddCommand(subcommands()...)
	return cmd
}

func main() {
	rootcmd.Main(cmd())
}
