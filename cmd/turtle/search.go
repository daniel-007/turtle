package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdSearch = &cobra.Command{
		Use:   "search",
		Short: "Print emojis with a name that contains the search string",
		Long:  "Print emojis with a name that contains the search string",
		RunE:  runSearch,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("require search string")
			}
			return nil
		},
	}
)

func runSearch(cmd *cobra.Command, args []string) error {
	s := args[0]

	emojis := turtle.Search(s)

	if emojis == nil {
		return fmt.Errorf("cannot find emojis for search %q", s)
	}

	j := NewJSONWriter(os.Stdout)

	return j.WriteEmojis(emojis)
}
