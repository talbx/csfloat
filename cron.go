package main

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/talbx/csfloat/types"
	"log"
	"time"
)

func RunCronSchedule(flags *types.InputConfig) {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Default().Panic(err)
	}

	jobCounter := 0
	_, err = s.NewJob(
		gocron.DurationJob(
			1*time.Minute,
		),
		gocron.NewTask(findSkinsTask(jobCounter), flags),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)
	if err != nil {
		log.Default().Fatal(err)
	}
	s.Start()
}

func findSkinsTask(jobCounter int) SkinFinder {
	return func(f *types.InputConfig) {
		FindSkins(f, jobCounter)
		jobCounter++
	}
}
