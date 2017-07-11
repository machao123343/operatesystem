package main

import(
"errors"
"math"
    "fmt"
)

func Sqrt(f float64) (float64 , error) {
    if f < 0 {
	   return 0, fmt.Errorf("math: square root of negative number %g", f)
    }
    return math.Sqrt(f) , errors.New(" ")
}

func main() {
    f, err := Sqrt(-1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%f" , f)
}
