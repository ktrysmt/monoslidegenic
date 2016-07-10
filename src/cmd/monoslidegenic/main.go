package main

import (
	"github.com/urfave/cli"
	"math/rand"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "OMIKUJI"
	app.Usage = "Get today's fortune!"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "level",
			Value: "normal",
			Usage: "Fortune level, it is hard, easy and normal",
		},
	}
}

func shuffle(arr []string) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
