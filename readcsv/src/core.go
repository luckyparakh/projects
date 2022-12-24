package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	ques string
	ans  int
}

type Core struct {
	path     string
	timer    int
	problems []problem
}

func NewCore(path string, timer int) *Core {
	return &Core{
		path:     path,
		timer:    timer,
		problems: []problem{},
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "Error while opening file:%v\n", err)
	os.Exit(1)
}

func (c *Core) ReadCSV() {
	problems := []problem{}
	// Can use CSV file
	file, err := os.Open(c.path)
	if err != nil {
		exit(err)
	}
	defer file.Close()
	fmt.Println("File read successfully")
	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines) // Not needed as NewScanner will split by default by lines
	for scanner.Scan() {
		quesAns := scanner.Text()
		ques := strings.Split(quesAns, ",")[0]
		ans, err := strconv.Atoi(strings.Split(quesAns, ",")[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while answer from CSV: %v\n", err)
			os.Exit(1)
		}
		p := problem{
			ques,
			ans,
		}
		problems = append(problems, p)
	}
	c.problems = problems
}

func (c *Core) ReadUI() {
	uiC := make(chan int)
	errC := make(chan error)
	correctAns := 0
	for cnt, p := range c.problems {
		fmt.Printf("Q%v %v: ", cnt+1, p.ques)
		go userInput(uiC, errC)
		select {
		case ui := <-uiC:
			if ui == p.ans {
				correctAns++
			}
		case errV := <-errC:
			exit(errV)
		case <-time.After(time.Duration(c.timer) * time.Second):
			fmt.Println("Timeout")
		}
	}
	fmt.Printf("Total Question %v, correct answer %v\n", len(c.problems), correctAns)
}

func userInput(c chan int, errC chan error) {
	var userInput int
	if _, err := fmt.Scanln(&userInput); err != nil {
		errC <- err
	}
	c <- userInput
}
