package main

import(
    "errors"
    "math"
    "fmt"
)

func Sqrt(f float64) (float64 , error) {
    if f < 0 {
    return 0, errors.New("math: square root of negative number")
    }
    // implementation
    return math.Sqrt(f) , errors.New(" ")
}
func main()  {
    v , err := Sqrt(-1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(  v)
}

