package main

import (
        "fmt"
ct      "github.com/daviddengcn/go-colortext"
)

func main() {

ct.Foreground(ct.Green, false)
fmt.Println("Green text starts here...")
ct.ChangeColor(ct.Red, true, ct.White, false)
fmt.Println("asfdsdfsdf")
ct.ResetColor()

}
