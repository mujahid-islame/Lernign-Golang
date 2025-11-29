package main

import "fmt"

func findSubstring(s string, words []string) []int {
    if len(words) == 0 || len(s) == 0 {
        return []int{}
    }

    wordLen := len(words[0])
    numWords := len(words)
    totalLen := wordLen * numWords
    result := []int{}

    wordCount := map[string]int{}
    for _, w := range words {
        wordCount[w]++
    }

    for i := 0; i <= len(s)-totalLen; i++ {
        seen := map[string]int{}
        j := 0
        for ; j < numWords; j++ {
            word := s[i+j*wordLen : i+(j+1)*wordLen]
            if _, ok := wordCount[word]; !ok {
                break
            }
            seen[word]++
            if seen[word] > wordCount[word] {
                break
            }
        }
        if j == numWords {
            result = append(result, i)
        }
    }

    return result
}

func main() {
    fmt.Println(findSubstring("barfoothefoobarman", []string{"foo","bar"})) // [0 9]
    fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"})) // []
    fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar","foo","the"})) // [6 9 12]
}
