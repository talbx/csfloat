package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "float",
	Short: "csfloat checker",
	Long:  `A csfloat checker`,
	Run:   run,
}

func init() {
	rootCmd.Flags().IntP("max", "m", 1000, "Max price")
	rootCmd.Flags().Float64P("discount", "d", 5.00, "Min Discount")
	rootCmd.Flags().IntP("discountValue", "v", 10, "Min Discount Value (cents)")
	rootCmd.Flags().IntP("category", "c", 1, "Item category - default normal (1)")
	rootCmd.Flags().BoolP("stickers", "s", false, "Show stickers?")
	rootCmd.Flags().IntP("top", "t", 5, "Top List")
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
