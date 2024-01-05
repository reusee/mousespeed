package main

import (
	"fmt"
	"os/exec"
)

var (
	pt = fmt.Printf
)

func main() {
	left := float64(-1)
	right := float64(1)
	for {
		m := (left + right) / 2
		_, err := exec.Command("gsettings", "set", "org.gnome.desktop.peripherals.mouse", "speed", fmt.Sprintf("%v", m)).CombinedOutput()
		if err != nil {
			panic(err)
		}
		for {
			var input string
			pt("+: faster, -: slower, =: ok\n")
			fmt.Scanf("%v", &input)
			switch input {
			case "+":
				left = m
			case "-":
				right = m
			case "=":
				pt("%v\n", m)
				return
			default:
				continue
			}
			break
		}
	}
}
