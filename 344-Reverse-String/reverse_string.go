package main

import (
    "fmt"
)

func reverseString(s []byte)  {
    length := len(s)

    for i := 0; i < length / 2;  i++ {
        s[i], s[length-1-i] = s[length-1-i], s[i]
    }

}


func main() {
    arr := []byte{'h','e','l','l','o'}
    fmt.Println(arr)
    reverseString(arr)
    fmt.Println(arr)
    arr = []byte{'H','a','n','n','a','h'}
    fmt.Println(arr)
    reverseString(arr)
    fmt.Println(arr)
}
