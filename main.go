package main

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "float",
	Short: "CSFloat price check",
	Long:  `A CSFloat price checker CLI`,
	Run:   run,
}

func init() {
	rootCmd.Flags().Bool("cron", false, "Enable cron mode")
	rootCmd.Flags().BoolP("auctions", "a", false, "Also check auctions")
	rootCmd.Flags().IntP("max", "m", 0, "Max price in cents")
	rootCmd.Flags().IntP("min", "n", 0, "Min price in cents")
	rootCmd.Flags().Float64P("discount", "d", 5.00, "Min discount percentage")
	rootCmd.Flags().IntP("category", "c", 1, "Item category - [0: Any, 1: Normal, 2: Stattrak, 3: Souvenir]")
	rootCmd.Flags().BoolP("stickers", "s", false, "Show stickers? (Default off)")
	rootCmd.Flags().IntP("top", "t", 10, "Top List")
	rootCmd.Flags().StringP("keyfile", "f", "", "The location of your API key file")
	rootCmd.Flags().StringP("keyword", "k", "", "The keyword. e.g a Skin Name like 'Asiimov' or 'Dragon Lore'")
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

	if flags.Cron {
		c := make(chan string)
		RunCronSchedule(flags, c)
		_ = <-c
	}
	FindSkins(flags)
}
