/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/garciawell/go-full-pos/36-cli/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
		RunE: runCreate(categoryDb),
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			fmt.Println("Here")
		},
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().String("name", "", "Category name")
	createCmd.Flags().String("description", "", "Category description")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
