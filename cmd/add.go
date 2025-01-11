package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/PushpinderDeswal/go_bmk/repository"
	"github.com/PushpinderDeswal/go_bmk/services"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new url to bookmark list",
	Long:  `Save new URL as a bookmark in bookmark database`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Required url as first argument.")
			os.Exit(1)
		}
		inputUrl := args[0]
		parsedUrl, err := url.ParseRequestURI(inputUrl)

		if err != nil {
			fmt.Println("Invalid url")
			os.Exit(1)
		}

		serv := services.NewBookmarkService(repository.MakeSQLiteBookmarkRepository(db))

		err = serv.AddBookmark(cmd.Context(), parsedUrl.String())

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Url saved as bookmark.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
