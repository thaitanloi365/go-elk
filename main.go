package main

import (
	"log"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/kjk/dailyrotate"
	logger "github.com/thaitanloi365/go-log"
)

const (
	logPath = "./logs/go.log"
)

func main() {
	logrotate, err := dailyrotate.NewFile(filepath.Join("logs/db", "2006-01-02.log"), func(path string, didRotate bool) {})
	if err != nil {
		panic(err)
	}

	var logger = logger.New(&logger.Config{
		BufferedSize: 50,
		Writer:       log.New(logrotate, "", 0),
	})

	i := 0
	for {
		i++
		time.Sleep(time.Second * 3)
		if rand.Intn(10) == 1 {
			logger.Errorf("error because test: %d", i)
		} else {
			logger.Infof("test log: %d", i)
		}
	}
}
