package cmd

import (
	"log"
	"os"
	"path"

	"github.com/leaanthony/mewn"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a shortcut repository in current working directory",
	Run: func(cmd *cobra.Command, args []string) {
		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal("Error on getting current path: ", err)
			return
		}

		metaFile, err := os.Create(path.Join(currentDirectory, "meta.yml"))
		if err != nil {
			log.Fatal("Error on creating file: ", err)
			return
		}
		defer metaFile.Close()

		meta := mewn.String("./assets/meta.yml.tpl")
		metaFile.WriteString(meta)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
