package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	// Github
	"github.com/eiannone/keyboard"
)

const WIDTH = 20
const HEIGHT = 50

type Point struct {
	x, y int
}

var defeat bool
var tail [100]Point
var snake, apple Point
var camp [WIDTH][HEIGHT]string
var dir_change = make(chan int)
var score, nTail, i, j, k, flag int

func setup() {
	score = 0
	snake = Point{
		x: WIDTH / 2,
		y: HEIGHT / 2,
	}
	defeat = false
	apple = Point{
		x: (rand.Intn(WIDTH-2) + 2),
		y: (rand.Intn(HEIGHT-2) + 2),
	}
}

func resetTerm() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("reset")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clean() {
	for i = 0; i < WIDTH; i++ {
		for j = 0; j < HEIGHT; j++ {
			if i == 0 || i == WIDTH-1 || j == 0 || j == HEIGHT-1 {
				camp[i][j] = "\033[0;34m#"
			} else {

				camp[i][j] = "\033[0m "
			}

			if j == snake.y && i == snake.x {
				camp[i][j] = "\033[0;32m0"
			} else if j == apple.y && i == apple.x {
				camp[i][j] = "\033[0;31m@"
			} else {
				for k = 0; k < nTail; k++ {
					if tail[k].y == j && tail[k].x == i {
						camp[i][j] = "\033[0;32m0"
					}
				}
			}

		}
	}
}

func draw() {
	fmt.Println("\033c")
	fmt.Println("Cobrinha em ASCII - Golang")
	for i = 0; i < WIDTH; i++ {
		for j = 0; j < HEIGHT; j++ {
			fmt.Print(camp[i][j])
		}
		fmt.Println("\033[0m")
	}
	fmt.Println("\nPontos: ", score)
	fmt.Printf("Clique em X para Sair\n")
}

func input() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetKey()
		switch char {
		case 'a':
			dir_change <- 1
		case 's':
			dir_change <- 2
		case 'd':
			dir_change <- 3
		case 'w':
			dir_change <- 4
		case 'x':
			defeat = true
		default:
			break
		}
	}
}

func logic() {
	prev1 := Point{
		x: tail[0].x,
		y: tail[0].y,
	}
	var prev2 Point
	tail[0] = Point{
		x: snake.x,
		y: snake.y,
	}

	for k = 1; k < nTail; k++ {
		prev2 = Point{
			x: tail[k].x,
			y: tail[k].y,
		}
		tail[k] = Point{
			x: prev1.x,
			y: prev1.y,
		}
		prev1 = prev2
	}

	switch flag {
	case 1:
		snake.y--
	case 2:
		snake.x++
	case 3:
		snake.y++
	case 4:
		snake.x--
	}

	if snake.x < 1 || snake.x > WIDTH-2 || snake.y < 1 || snake.y > HEIGHT-2 {
		defeat = true
	}

	if apple.x < 1 || apple.x > WIDTH-2 || apple.y < 1 || apple.y > HEIGHT-2 {
		apple = Point{
			x: (rand.Intn(WIDTH-2) + 2),
			y: (rand.Intn(HEIGHT-2) + 2),
		}
	}

	for k = 0; k < nTail; k++ {
		if tail[k].x == snake.x && tail[k].y == snake.y {
			defeat = true
		}
	}

	if snake.x == apple.x && snake.y == apple.y {
		score += 10
		apple = Point{
			x: (rand.Intn(WIDTH-2) + 2),
			y: (rand.Intn(HEIGHT-2) + 2),
		}
		nTail++
	}
}

func main() {
	defer resetTerm()
	setup()
	go input()
	for !defeat {
		clean()
		draw()
		select {
		case dir := <-dir_change:
			flag = dir
		default:
		}
		logic()
		time.Sleep(80 * time.Millisecond)
	}
}
