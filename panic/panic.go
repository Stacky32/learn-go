package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const magicWord string = "goblin"
const sectionDelimeter = "=========="

func main() {
	fmt.Println("Enter lines of text, but don't say a certain word.")
	play()
	fmt.Println("Game over.")
}

func play() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("You failed!", r)
		}
	}()

	inputs := captureInput()

	fmt.Println("You entered:")
	fmt.Println(sectionDelimeter)
	for _, l := range inputs {
		fmt.Println(l)
	}
	fmt.Println(sectionDelimeter)
	fmt.Println("Well done!")
}

func captureInput() []string {
	inputs := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("-> ")
		if !scanner.Scan() {
			break
		}

		txt := scanner.Text()
		if txt == "" {
			break
		}

		if err := checkSafe(txt); err != nil {
			panic(err)
		}

		inputs = append(inputs, txt)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return inputs
}

func checkSafe(txt string) error {
	if strings.Contains(strings.ToLower(txt), magicWord) {
		return fmt.Errorf("You can't mention %s", magicWord)
	}

	return nil
}
