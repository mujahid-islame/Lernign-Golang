package main

import "fmt"

func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    rows := make([][]rune, numRows)
    for i := range rows {
        rows[i] = []rune{}
    }

    index, step := 0, 1
    for _, ch := range s {
        rows[index] = append(rows[index], ch)
        if index == 0 {
            step = 1
        } else if index == numRows-1 {
            step = -1
        }
        index += step
    }

    result := ""
    for _, row := range rows {
        result += string(row)
    }
    return result
}

func main() {
    fmt.Println(convert("PAYPALISHIRING", 3)) // Output: PAHNAPLSIIGYIR
    fmt.Println(convert("PAYPALISHIRING", 4)) // Output: PINALSIGYAHRPI
    fmt.Println(convert("A", 1))              // Output: A
}
