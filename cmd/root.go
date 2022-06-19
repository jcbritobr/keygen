package cmd

import (
	"fmt"
	"os"

	"github.com/jcbritobr/keygen/uniengine"
	"github.com/spf13/cobra"
)

var (
	length    int
	symbol    bool
	lowercase bool
	uppercase bool
	number    bool
)

var rootCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Keygen is a very fast random key generator",
	Long: `Keygen is a very fast random key generator.
	Its possible to choose between lowercase, uppercase, number and symbol characters,
	and also the key's size`,
	Run: func(cmd *cobra.Command, args []string) {
		key := uniengine.GenerateRandomKey(lowercase, uppercase, number, symbol, length)
		fmt.Println(key)
	},
}

func init() {
	rootCmd.PersistentFlags().IntVar(&length, "length", 8, "inserts the key length")
	rootCmd.PersistentFlags().BoolVarP(&symbol, "symbol", "s", false, "inserts symbol characters in key")
	rootCmd.PersistentFlags().BoolVarP(&lowercase, "lowercase", "l", false, "inserts lowercase characters in key")
	rootCmd.PersistentFlags().BoolVarP(&uppercase, "uppercase", "u", false, "inserts uppercase characters in key")
	rootCmd.PersistentFlags().BoolVarP(&number, "number", "n", false, "inserts digit characters in key")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
