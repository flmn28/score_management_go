package management

import (
	"bufio"
	"fmt"
	"github.com/flmn28/score_management/src/score"
	"os"
	"strconv"
)

func Create() {
	fmt.Println(">please enter the score")
	stdin := bufio.NewScanner(os.Stdin)
  	stdin.Scan()
	text := stdin.Text()
	value, _ := strconv.Atoi(text)
	score.Create(value)
	fmt.Println(">Successfully created score")
}

func Read() {
	fmt.Println(">Show all score")
	for i, score := range score.All() {
		fmt.Printf("%d.%d\n", i + 1, score.Value)
	}
}

func Delete() {
	fmt.Println(">Please enter delete line number")
	stdin := bufio.NewScanner(os.Stdin)
  	stdin.Scan()
	text := stdin.Text()
	num, _ := strconv.Atoi(text)
	score.Delete(num)
	fmt.Println(">Successflly deleted score")
}

func Average() {
	fmt.Println(">Score Average:", score.Average())
}
