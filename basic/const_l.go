// niubi
package main

import "fmt"

type Weekday = int

// 相当于枚举
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

const (
    FlagNone = 1 << iota
    FlagRed
    FlagGreen
    FlagBlue
)


func main() {
    fmt.Println( Monday)
    fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
}