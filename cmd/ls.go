package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/PushpinderDeswal/go_bmk/repository"
	"github.com/PushpinderDeswal/go_bmk/services"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list bookmarks",
	Long:  `Show list of all bookmarks, can accept --id flag to filter based on id and --top to show only top x results`,
	Run: func(cmd *cobra.Command, args []string) {
		serv := services.NewBookmarkService(repository.MakeSQLiteBookmarkRepository(db))

		bmks, err := serv.GetAllBookmarks(cmd.Context())

		if err != nil {
			log.Fatal(err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

		fmt.Fprintln(writer, "ID\tURL\tCreated At")
		for _, bmk := range bmks {
			formattedCreatedAt := bmk.CreatedAt.Format("Jan 02, 2006 15:04")
			fmt.Fprintf(writer, "%v\t%v\t%v\n", bmk.ID, bmk.Url, formattedCreatedAt)
		}

		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
