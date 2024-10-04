package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/talbx/csfloat/listing"
	"log"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "float",
	Short: "CSFloat price check",
	Long:  `A CSFloat price checker CLI`,
	Run:   run,
}

var findCmd = &cobra.Command{
	Use:     "find",
	Short:   "Find a specific weapon",
	Example: "float find talon // looks a talon knife",
}

func init() {
	rootCmd.PersistentFlags().Bool("cron", false, "Enable cron mode")
	rootCmd.PersistentFlags().BoolP("auctions", "a", false, "Also check auctions")
	rootCmd.PersistentFlags().IntP("max", "m", 0, "Max price in cents")
	rootCmd.PersistentFlags().IntP("min", "n", 0, "Min price in cents")
	rootCmd.PersistentFlags().Float64P("discount", "d", 5.00, "Min discount percentage")
	rootCmd.PersistentFlags().IntP("category", "c", 1, "Item category - [0: Any, 1: Normal, 2: Stattrak, 3: Souvenir]")
	rootCmd.PersistentFlags().BoolP("stickers", "s", false, "Show stickers? (Default off)")
	rootCmd.PersistentFlags().IntP("top", "t", 10, "Top List")
	rootCmd.PersistentFlags().StringP("keyfile", "f", "", "The location of your API key file")
	rootCmd.PersistentFlags().StringP("keyword", "k", "", "The keyword. e.g a Skin Name like 'Asiimov' or 'Dragon Lore'")
	rootCmd.PersistentFlags().StringSlice("defIndex", nil, "The keyword. e.g a Skin Name like 'Asiimov' or 'Dragon Lore'")
	var subCmds []*cobra.Command
	for _, tuple := range listing.CreateIndices() {
		command := &cobra.Command{
			Use:   strings.ToLower(tuple.Name),
			Short: fmt.Sprintf("Look for %s", tuple.Name),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(cmd)
				flags, err := ParseFlags(cmd.PersistentFlags())
				if err != nil {
					log.Default().Fatal(err)
				}

				if flags.Cron {
					c := make(chan string)
					RunCronSchedule(flags, c)
					<-c
				}

				flags.DefIndex = []string{string(rune(tuple.Index))}
				FindSkins(flags)
			},
		}
		subCmds = append(subCmds, command)
	}
	findCmd.AddCommand(subCmds...)
	rootCmd.AddCommand(findCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Default().Fatal(err)
	}
}

func run(cmd *cobra.Command, _ []string) {
	fmt.Println(cmd)
	flags, err := ParseFlags(cmd.PersistentFlags())
	if err != nil {
		log.Default().Fatal(err)
	}

	if flags.Cron {
		c := make(chan string)
		RunCronSchedule(flags, c)
		<-c
	}
	FindSkins(flags)
}
