package main 

import (
    "fmt"
    "sort"
)

func main() {
    var a *[]string
    var c []string
    c = []string{"he", "he2", "he3"}
    a = &c

    b := "he22"
    fmt.Println(*a, b)

    sort.Strings(*a)
    i := sort.SearchStrings(*a, b)
    fmt.Println(i < len(*a) && (*a)[i] == b)




}
