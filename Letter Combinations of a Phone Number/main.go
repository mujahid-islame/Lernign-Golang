package main

import "fmt"

var phoneMap = map[byte]string{
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    result := []string{}
    backtrack(0, digits, "", &result)
    return result
}

func backtrack(index int, digits, path string, result *[]string) {
    if index == len(digits) {
        *result = append(*result, path)
        return
    }

    letters := phoneMap[digits[index]]
    for i := 0; i < len(letters); i++ {
        backtrack(index+1, digits, path+string(letters[i]), result)
    }
}

func main() {
    fmt.Println(letterCombinations("23"))
    fmt.Println(letterCombinations("2"))
}
