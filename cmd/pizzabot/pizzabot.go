package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andrewcopp/pizzabot"
)

var config *pizzabot.Config

func init() {
	config = &pizzabot.Config{}
	err := config.Parse(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	grid := pizzabot.NewGrid(config)
	bot := pizzabot.NewSimple(config)
	out, err := bot.Solve(grid)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(out)
}
