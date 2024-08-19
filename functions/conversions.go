package functions

import (
    "fmt"
    "strconv"
)

func HextoInt(hex string) string {
    number, err := strconv.ParseInt(hex, 16, 64)
    if err != nil {
        fmt.Println("Error converting hex:", err)
        return ""
    }
    return fmt.Sprint(number)
}

func BintoInt(bin string) string {
    number, err := strconv.ParseInt(bin, 2, 64)
    if err != nil {
        fmt.Println("Error converting bin:", err)
        return ""
    }
    return fmt.Sprint(number)
}
