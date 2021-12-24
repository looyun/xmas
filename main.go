package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kyokomi/emoji/v2"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

const height = 30
const width = 150

var p = make([][]string, height)
var snow = make([][]string, height)

func initSnow() {
	for h := 0; h < height; h++ {
		snow[h] = make([]string, width)
		for w := 0; w < width; w++ {
			if rand.Intn(100) < 1 {
				snow[h][w] = "❄"
			}
		}
	}
}

func printTree() {
	fmt.Println()
	fmt.Println()

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if snow[h][w] != "" {
				fmt.Print(snow[h][w])
			} else {
				fmt.Print(p[h][w])
			}
		}
		fmt.Println()
	}

	newSnowLine := make([]string, width)
	for w := 0; w < width; w++ {
		if rand.Intn(100) < 2 {
			newSnowLine[w] = "❄"
		}
	}
	snow = append([][]string{newSnowLine}, snow[:height-1]...)
}

func main() {
	fmt.Println("merry chrismas")
	ticker := time.NewTicker(time.Second * 1)
	xmas()
	initSnow()
	for {
		select {
		case <-ticker.C:
			fmt.Print("\033[H\033[2J")
			printTree()
		}
	}

}

func xmas() {

	for h := 0; h < height; h++ {
		p[h] = make([]string, width)
		for w := 0; w < width; w++ {

			if h == 0 && w == width/2 {
				p[h][w] = Red + "⭑" + Reset

			} else if h == height-2 || h == height-3 {
				if w >= width/2-1 && w <= width/2+1 {
					p[h][w] = Red + "|" + Reset
				} else {
					p[h][w] = " "
				}
			} else if w >= width/2-int(float64(h)*0.75) && w <= width/2+int(float64(h)*0.75) {
				var color string
				if rand.Intn(100) < 5 {
					color = randomColor()
				} else {
					color = Green
				}
				p[h][w] = color + randomCharacter() + Reset
			} else {
				p[h][w] = " "
			}
		}
	}
}

func randomCharacter() string {
	strs := emoji.Sprint("~!@#$%^&*()_+`{}|[]:;<>?,.'Oo")
	a := "⭑"
	b := "🎁"
	c := "✨"

	i := rand.Intn(len(strs))
	if rand.Intn(100) < 10 {
		return a
	}
	if rand.Intn(100) < 1 {
		return b
	}
	if rand.Intn(100) < 1 {
		return c
	}
	return string(strs[i])

}

func randomColor() string {
	i := rand.Intn(8)
	switch i {
	case 1:
		return Red
	case 2:
		return Green
	case 3:
		return Gray
	case 4:
		return Purple
	case 5:
		return Yellow
	case 6:
		return Blue
	case 7:
		return White
	case 8:
		return Cyan
	default:
		return Green
	}
}
