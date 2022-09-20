package main

import (
	"fmt"
	"math/rand"
)

type Player struct{ name string }

func main() {
	randonMin := 1
	randomMax := 100
	breakPoint := 11
	hit := 0

	var friends = []string{"Clara", "Fiqri", "Medy", "Lutfi"}

	var players []*Player

	for _, v := range friends {
		var player = &Player{
			name: v,
		}

		players = append(players, player)
	}

	for {
		hit++

		for _, player := range players {
			random := rand.Intn(randomMax-randonMin) + randonMin

			fmt.Printf("korek ada di %s pada hit ke %d dan mempunyai nilai %d\n", player.name, hit, random)
			if random%breakPoint == 0 {
				fmt.Printf("%s kalah pada hit ke %d\n", player.name, hit)
				return
			}
		}

		fmt.Printf("\n")
	}
}
