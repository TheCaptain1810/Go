package main

import "fmt"

func main() {
    type car struct {
        Make string
        Model string
    }

    newCar := car{}
    newCar.Make = "ford"

    fmt.Println("The car: ", newCar.Make)
}
