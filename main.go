package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "float",
	Short: "CSFloat price check",
	Long:  `A CSFloat price checker`,
	Run:   run,
}

func init() {
	rootCmd.Flags().IntP("max", "m", 0, "Max price in cents")
	rootCmd.Flags().Float64P("discount", "d", 5.00, "Min discount percentage")
	rootCmd.Flags().IntP("discountValue", "v", 10, "Min discount in cents")
	rootCmd.Flags().IntP("category", "c", 1, "Item category - [0: Any, 1: Normal, 2: Stattrak, 3: Souvenir]")
	rootCmd.Flags().BoolP("stickers", "s", false, "Show stickers? (Default off)")
	rootCmd.Flags().IntP("top", "t", 10, "Top List")
	rootCmd.Flags().StringP("keyfile", "k", "", "The location of your API key file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Default().Fatal(os.Stderr, err)
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Default().Fatal(err)
	}
}

func run(cmd *cobra.Command, _ []string) {
	flags, err := ParseFlags(cmd.Flags())
	if err != nil {
		log.Default().Fatal(err)
	}
	c := make(chan bool)
	RunCronSchedule(flags)
	<-c
}
