package main

import "fmt"

func main() {
    type car struct {
        Make string
        Model string
    }

    type truck struct {
        car // this is embedded struck (not the same as nested struct)
        bedSize int
    }

    type anotherTruck struct {
        anotherCar car
        bedSize int
    }

    newerTruck := anotherTruck {
        bedSize: 20,
        anotherCar: car {
            Make: "ford",
            Model: "2020",
        },
    }

    fmt.Printf("The anotherTruck: %s\n", newerTruck.anotherCar.Make)

    newTruck := truck {
        bedSize:  20,
        car: car {
            Make: "ford",
            Model: "2028",
        },
    }

    fmt.Printf("The truck: %s\n", newTruck.Make)

    newCar := car{}
    newCar.Make = "ford"

    fmt.Println("The car: ", newCar.Make)
}
