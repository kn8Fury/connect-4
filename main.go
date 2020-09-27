package main

import (
	"fmt"

	"github.com/kn8Fury/connect4/game/colors"

	"github.com/kn8Fury/connect4/game/results"

	"github.com/kn8Fury/connect4/game"
)

func main() {
	fmt.Println("Hello, world!")
	request := ""
	for request != "START" {
		fmt.Println("Enter 'START' to start the game")
		fmt.Scanf("%s\n", &request)
	}
	gb := game.NewBoard()
	result := results.Invalid
	var err error
	var player colors.Color
	var column int
	fmt.Println("READY")
	for result == results.Invalid || result == results.Continue {
		fmt.Println("Current status of game:")
		fmt.Println(gb.Draw())
		player = gb.CurrentPlayer()
		fmt.Println("Current Player: ", player)
		fmt.Print("Requesting column... ")
		_, err = fmt.Scanf("%d\n", &column)
		for err != nil {
			fmt.Print("Invalid request recieved. Please try again. Requesting column... ")
			_, err = fmt.Scanf("%d", &column)
		}
		fmt.Printf("Dropping %s coin in column %d\n", player, column)
		result, err = gb.Drop(player, column)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("GAME OVER!")
	fmt.Println(gb.Draw())
	if result == results.Win {
		fmt.Println("Winner: ", gb.CurrentPlayer())
	} else {
		fmt.Println("It is a draw.")
	}
}
