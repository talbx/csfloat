package main

import (
	"github.com/spf13/cobra"
	"github.com/talbx/csfloat/pkg"
	"github.com/talbx/csfloat/pkg/listing"
	"log"
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
	subCmds := generateSubCommands()
	findCmd.AddCommand(subCmds...)
	rootCmd.AddCommand(findCmd)
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

	// doesnt make sense to filter defIndex when you e.g use float find pistols
	rootCmd.Flags().StringSlice("defIndex", nil, "The keyword. e.g a Skin Name like 'Asiimov' or 'Dragon Lore'")
}

func generateSubCommands() []*cobra.Command {
	knives := GenerateFindCommand("knife", listing.KNIFES)
	smgs := GenerateFindCommand("smg", listing.SMGs)
	rifles := GenerateFindCommand("rifle", listing.RIFLES)
	gloves := GenerateFindCommand("gloves", listing.GLOVES)
	pistol := GenerateFindCommand("pistol", listing.PISTOLS)
	heavy := GenerateFindCommand("heavy", listing.MGs)
	pumps := GenerateFindCommand("pumps", listing.PUMPS)
	misc := GenerateFindCommand("misc", listing.MISC)

	return []*cobra.Command{
		knives, smgs, rifles, gloves, pistol, heavy, pumps, misc,
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Default().Fatal(err)
	}
}

func run(cmd *cobra.Command, _ []string) {
	flags, err := pkg.ParseFlags(cmd.PersistentFlags())

	if err != nil {
		log.Default().Fatal(err)
	}
	defIndex, err := cmd.Flags().GetStringSlice("defIndex")
	if err != nil {
		log.Default().Fatal(err)
	}

	flags.DefIndex = defIndex

	if flags.Cron {
		c := make(chan string)
		pkg.RunCronSchedule(flags, c)
		<-c
	}
	pkg.FindSkins(flags)
}
