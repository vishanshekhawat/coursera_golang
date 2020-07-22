/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.
As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.
A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
Submit your Go program source code.
*/

package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func swap (slice [] int, i int){

    one := slice[i+1]
    two := slice[i]

    slice[i] = one
    slice[i+1] = two
}

func BubbleSort (slice [] int){

    sorted := false

    for sorted == false{

        swapped := false

        for i:=0; i < (len(slice)-1); i++{

            if slice[i] > slice[i+1]{
                swap(slice, i)
                swapped = true
            }
        }

        if swapped == false{
            sorted = true
        }
    }
} 

func main(){

    readInput := bufio.NewScanner(os.Stdin)
    fmt.Println("\nPlease enter a sequence of up to ten integers separated by spaces.")
    fmt.Println("Hit enter when finished:")
    readInput.Scan()
    line := readInput.Text()

    userInts := strings.Split(line, " ")

    var userSLice []int

    for i:=0; i < len(userInts); i++ {
        user_int, err := strconv.Atoi(userInts[i])
        if err == nil {
            userSLice = append(userSLice, userInts)
        } else {
            fmt.Println("Invalid input. Please run the program again with ")
            fmt.Println("a sequence of up to ten integers separated by spaces. ")
            fmt.Print("Exiting program ")
            os.Exit(1)
        }

    }

    BubbleSort(userSLice)

    fmt.Println("Sorted slice of integers:")

    for i:=0; i < len(userSLice); i++{
        fmt.Printf("%d ", userSLice[i])
    }

    fmt.Println()
}
