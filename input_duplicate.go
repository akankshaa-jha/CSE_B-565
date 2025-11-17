package main

import "fmt"

func main() {

    var n int
    fmt.Print("Enter size of array: ")
    fmt.Scan(&n)

    arr := make([]int, n)

    fmt.Println("Enter", n, "elements:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

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
