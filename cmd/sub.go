package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/talbx/csfloat/pkg"
	"github.com/talbx/csfloat/pkg/listing"
	"log"
	"strconv"
	"strings"
)

func generateCommand(title string, defIndexFunc func() []string) *cobra.Command {
	command := &cobra.Command{
		Use:   strings.ToLower(title),
		Short: fmt.Sprintf("Look for %s", title),
		Run: func(cmd *cobra.Command, args []string) {
			flags, err := pkg.ParseFlags(cmd.Flags())
			if err != nil {
				log.Default().Fatal(err)
			}

			if flags.Cron {
				c := make(chan string)
				pkg.RunCronSchedule(flags, c)
				<-c
			}

			flags.DefIndex = defIndexFunc()
			pkg.FindSkins(flags)
		},
	}
	return command
}

func GenerateFindCommand(title string, tuples []listing.Tuple) *cobra.Command {
	command := generateCommand(title, func() []string {
		var r []string
		for _, tuple := range tuples {
			r = append(r, strconv.Itoa(tuple.Index))
		}
		return r
	})

	for _, tuple := range tuples {
		subCmd := generateCommand(tuple.Name, func() []string {
			return []string{strconv.Itoa(tuple.Index)}
		})
		command.AddCommand(subCmd)
	}
	return command
}
