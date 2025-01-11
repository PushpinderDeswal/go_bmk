package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/PushpinderDeswal/go_bmk/repository"
	"github.com/PushpinderDeswal/go_bmk/services"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a bookmark",
	Long: `rm remove a bookmark using bookmark id: 
		USAGE: bmk rm <bookmark_id>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide bookmark id")
			os.Exit(1)
		}
		if uuid.Validate(args[0]) != nil {
			fmt.Println("Invalid bookmark id")
			os.Exit(1)
		}
		serv := services.NewBookmarkService(repository.MakeSQLiteBookmarkRepository(db))

		if err := serv.DeleteBookmark(cmd.Context(), args[0]); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Deleted bookmark with id %v\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
