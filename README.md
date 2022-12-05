## Pwdgen

Password generator written in Go

By default, 16 digit password is generated consiting of the following:
* 4 Uppercase alphabetical characters
* 4 Lowercase alphabetical characters
* 4 Digits
* 4 Spcial characters

It is possible to change this by setting the config structure:

    pwdgen.Config.LowerAlphas = 3
    pwdgen.Config.UpperAlphas = 5
    pwdgen.Config.Digits = 1
    pwdgen.Config.Specials = 2

Basic usage:

    package main

    import (
        "fmt"
        "pwdgen/pwdgen"
    )

    func main() {
        fmt.Println(pwdgen.Password())
    }
