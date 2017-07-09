package main

import (
    "os"
    "fmt"
    "bufio"
    "bytes"
    "io"
)

func main() {
    file,err3 := os.Open("d:/code/1.java")
    if err3 != nil {
        if pe , ok := err3.(*os.PathError);ok {
            fmt.Printf("Path Error: %s (op = %s , path = %s)\n" , pe.Err , pe.Op , pe.Path)
        }else {
            fmt.Printf("uknown Error : %s\n" , err3)
        }
    }
    r := bufio.NewReader(file)
    var buf bytes.Buffer
    for {
        byteArray, _, err4 := r.ReadLine()
        if err4 != nil {
            if err4 != io.EOF {
                break
            } else {
                fmt.Printf("Read Error: %s\n", err4)
                break
            }
        } else {
            buf.Write(byteArray)
        }
    }
}