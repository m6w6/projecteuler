package main

import (
	"fmt"
	"os"
)

func grid(filename string) [20][20] int {
	var grid [20][20]int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for line := 0; line < 20; line++ {
		for row := 0; row < 20; row++ {
			_, err := fmt.Fscanf(file, "%d", &grid[line][row])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
	}

	return grid
}

func add(i *int) {
	*i++
}
func sub(i *int) {
	*i--
}
func addX(x, y *int) {
	add(x)
}
func addY(x, y *int) {
	add(y)
}
func addXY(x, y *int) {
	add(x)
	add(y)
}
func addXsubY(x, y *int) {
	add(x)
	sub(y)
}

type next func(x,y *int)

func main() {
	grid := grid("src/11.txt")
	nexts := []next{addX, addY, addXY, addXsubY}
	max := 0
	for X := 0; X < 20; X++ {
		for Y := 0; Y < 20; Y++ {
			NEXT:
			for fn := range nexts {
				product := 1
				x := X
				y := Y
				for i := 0; i < 4; i++ {
					if x < 0 || x >= 20 || y < 0 || y >= 20 {
						continue NEXT
					}
					product *= grid[x][y]
					nexts[fn](&x, &y)
				}
				if product > max {
					max = product
				}
			}
		}
	}
	fmt.Println(max)
}
