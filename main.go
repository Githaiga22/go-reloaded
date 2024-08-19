package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-reloaded/functions"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: program <sample.txt> <result.txt>")
		return
	}

	givenText, _ := os.ReadFile(args[1])
	words := strings.Split(string(givenText), " ")

	for i, word := range words {
		if i == 0 {
			continue
		}
		switch word {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap)":
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(hex)":
			words[i-1] = functions.HextoInt(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(bin)":
			words[i-1] = functions.BintoInt(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(up,", "(low,", "(cap,":
			if i+1 >= len(words) || i+2 >= len(words) {
				continue
			}
			number, _ := strconv.Atoi(strings.Trim(words[i+1], " )"))
			for j := 1; j <= number; j++ {
				switch word {
				case "(up,":
					words[i-j] = strings.ToUpper(words[i-j])
				case "(low,":
					words[i-j] = strings.ToLower(words[i-j])
				case "(cap,":
					words[i-j] = strings.Title(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
		}
	}

	words = functions.Punctuations(words)
	functions.ChangeA(words)

	needed := strings.Join(words, " ")
	err := os.WriteFile(args[1], []byte(needed), 0o644)
	if err != nil {
		panic(err)
	}
}
