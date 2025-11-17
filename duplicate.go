package main

import "fmt"

func main() {
    arr := []int{1, 2, 2, 3, 1, 4}

    uniq := []int{}

    for _, v := range arr {
        found := false

        for _, u := range uniq {
            if v == u {
                found = true
                break
            }
        }

        if !found {
            uniq = append(uniq, v)
        }
    }

    fmt.Println("Unique array:", uniq)
}
