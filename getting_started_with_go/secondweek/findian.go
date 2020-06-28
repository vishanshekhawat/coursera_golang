
package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main(){

    var ian string

    fmt.Printf("Please enter one string: ")
    input := bufio.NewReader(os.Stdin)
    ian_untrimmed, err := input.ReadString('\n')

    if (err == nil){
        ian = strings.TrimSuffix(ian_untrimmed, "\n")
        ian_lower := strings.ToLower(ian)

        starts := strings.HasPrefix(ian_lower, "i")
        ends := strings.HasSuffix(ian_lower, "n")
        contains := strings.Contains(ian_lower, "a")

        if (starts && contains && ends){

            fmt.Printf("Found!\n")
        } else{
            fmt.Printf("Not Found!\n")
        }
    } else{
        fmt.Printf("Input Error\n")
    }
}
