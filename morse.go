package main

import (
	"bufio"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"strings"
	"time"
)

var (
	pin     = rpio.Pin(10)
	library = map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		"0": "-----",
		" ": " ",
	}
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Alphanumeric Text (\"!exit\" to quit): ")
		scanner.Scan()
		text := scanner.Text()
		if text == "!exit" {
			break
		} else {
			fmt.Print("\n")
			output := convert(text)
			blinker(output)
		}

	}

}

func convert(input string) (output string) {

	fmt.Println("Input string: ", input)

	input = strings.ToUpper(input)
	for _, c := range input {
		if _, ok := library[string(c)]; !ok {
			fmt.Println("Please enter alphanumeric characters only!  Invalid character:  ", string(c))
			return
		}
	}
	for _, c := range input {
		output += library[string(c)]
	}

	fmt.Println("Output morse: ", output)

	return output
}
func blink(duration time.Duration) {
	pin.High()
	//fmt.Println("pin.High()")
	time.Sleep(duration)
	pin.Low()
	//fmt.Println("pin.Low()")
	time.Sleep(time.Second / 5)
}

func blinker(input string) {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		if err := rpio.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	pin.Output()

	for _, c := range input {

		switch string(c) {
		case ".":
			fmt.Println("dot")
			blink(time.Second / 5)
		case "-":
			fmt.Println("dash")
			blink(time.Second / 2)
		default:
			fmt.Println("space")
			time.Sleep(time.Second / 5)
		}
	}
}
