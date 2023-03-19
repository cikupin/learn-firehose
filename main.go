package main

import (
	"log"
	"time"
	_ "time/tzdata"

	"github.com/cikupin/learn-firehose/cmd"
)

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalln(err.Error())
	}
	time.Local = loc

	cmd.Execute()
}
