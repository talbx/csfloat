package main

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/talbx/csfloat/types"
	"log"
	"time"
)

func RunCronSchedule(flags *types.SearchConfig, ch chan string) {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Default().Panic(err)
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			1*time.Minute,
		),
		gocron.NewTask(FindSkins, flags),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)
	if err != nil {
		log.Default().Fatal(err)
	}
	s.Start()
}
