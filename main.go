package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func Punctuations(s []string) []string {
    puncs := []string{",", ".", "!", "?", ":", ";"}

   // Move punctuation from beginning of a word to end if not already there
    for i, word := range s {
        for _, punc := range puncs {
            if string(word[0]) == punc && string(word[len(word)-1]) != punc {
                s[i-1] += punc
                s[i] = word[1:]
            }
        }
    }

   // Remove punctuation at end of string and concatenate to previous word
    for i, word := range s {
        for _, punc := range puncs {
            if string(word[0]) == punc && s[len(s)-1] == word {
                s[i-1] += word
                s = s[:len(s)-1]
            }
        }
    }

   // Remove punctuation at end of word and concatenate to previous word, if not already at end of string
    for i, word := range s {
        for _, punc := range puncs {
            if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
                s[i-1] += word
                s = append(s[:i], s[i+1:]...)
            }
        }
    }

    // Handling apostrophes
    count := 0
    for i, word := range s {
        if word == "'" && count == 0 {
            count++
            s[i+1] = word + s[i+1]
            s = append(s[:i], s[i+1:]...)
        }
    }

    // Handling second apostrophes
    for i, word := range s {
        if word == "'" {
            s[i-1] = s[i-1] + word
            s = append(s[:i], s[i+1:]...)
        }
    }

    return s
}

func main() {
    args := os.Args[1:]

    // reading first file
    givenText, _ := os.ReadFile(args[0])

    // array to push the words into
    words := strings.Split(string(givenText), " ")

    for i, word := range words {
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
            words[i-1] = HextoInt(words[i-1])
            words = append(words[:i], words[i+1:]...)
        case "(bin)":
            words[i-1] = BintoInt(words[i-1])
            words = append(words[:i], words[i+1:]...)
        case "(up,":
            b := strings.Trim(string(words[i+1]), words[i+1][1:])
            number, _ := strconv.Atoi(string(b))
            for j := 1; j <= number; j++ {
                words[i-j] = strings.ToUpper(words[i-j])
            }
            words = append(words[:i], words[i+2:]...)
        case "(low,":
            b := strings.Trim(string(words[i+1]), words[i+1][1:])
            number, _ := strconv.Atoi(string(b))
            for j := 1; j <= number; j++ {
                words[i-j] = strings.ToLower(words[i-j])
            }
            words = append(words[:i], words[i+2:]...)
        case "(cap,":
            b := strings.Trim(string(words[i+1]), words[i+1][1:])
            number, _ := strconv.Atoi(string(b))
            for j := 1; j <= number; j++ {
                words[i-j] = strings.Title(words[i-j])
            }
            words = append(words[:i], words[i+2:]...)
        }
    }

    words = Punctuations(words)

    // Apply the "ChangeA" function to the words slice
    ChangeA(words)

    // Join the slice with spaces and store the result in a new variable
    needed := strings.Join(words, " ")

    // write file, automatically updates manipulated file.
    man := os.WriteFile(args[1], []byte(needed), 0644)
    if man != nil {
        panic(man)
    }
}

// conv hex to int
func HextoInt(hex string) string {
    number, _ := strconv.ParseInt(hex, 16, 64)
    return fmt.Sprint(number)
}

// conv binary to int
func BintoInt(bin string) string {
    number, _ := strconv.ParseInt(bin, 2, 64)
    return fmt.Sprint(number)
}

func shouldReplaceWithAn(word, nextLetter string) bool {
    vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
    for _, letter := range vowels {
        if (word == "a" || word == "A") && string(nextLetter[0]) == letter {
            return true
        }
    }
    return false
}

// StringBuilder-like function
func buildString(parts []string) string {
    result := ""
    for _, part := range parts {
        result += part
    }
    return result
}

// Function to replace "a" with "an" if necessary
func ChangeA(s []string) {
    for i, word := range s {
        if i < len(s)-1 {
            nextLetter := string(s[i+1][0])
            if shouldReplaceWithAn(word, nextLetter) {
                s[i] = "an"
            }
        }
    }
}