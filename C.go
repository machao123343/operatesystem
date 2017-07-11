stdout0 , err := cmd.StdoutPipe()
if err := cmd).Start(); err != nil {
    fmt.Printf("Error : the command No.0 can not be startup: %s\n" , err)
    return
}

stdout0 , err := cmd0.StdoutPipe()
if err != nil {
    fmt.Printf("Error : Can not obtain the stdout pipe for command No.0 : %s\n" , err)
    return
}

if err := cmd.Start() ; err != nil {
    fmt.Printf("Error : Can not obtain the stdout pipe for command No.0 : %s\n" , err)
    return
}

output0 := make([]byte , 30)
n , err = stdout0.Read(output0)
if err != nil {
    fmt.Printf("Error : Can not read data form the pipe : %s\n" , err)
    return
}
fmt.Printf("%s\n" , output0[:n])

var outputBuf0 bytes.Buffer

for {
    tempOutput := make([]byte , 5)
    n , err := stdout0.Read(tempOutput)
    if err != nil {
        if err == io.EOF {
            break
        } else {
            fmt.Printf("Error: Can not read data from the pipe : %s\n" , err)
            return
        }
    }
    if n > 0 {
        outputBuf0.Write(tempOutput[:n])
    }
}
fmt.Printf("%s\n" , outputBuf0.String())
cmd1 := exec.Command("ps" , "aux")
cmd2 := exec.Command("grep" , "apipe")

stdout1 , err := cmd1.StdoutPipe()
if err != nil {
    fmt.Printf("Error: Can not obtain the stdout pipe for command: %s\n" , err)
    return
}

outputBuf1 := bufio.NewReader(stdout1)
stdin2 , err := cmd2.StdinPipe()
if err != nil {
    fmt.Printf("Error: Can not obtain the stdin pipe for command: %s\n" , err)
    return
}
outputBuf1.WriteTo(stdin2)
var outputBuf2 bytes.buffer
cmd2.Stdout = &outputBuf2

//关闭输入管道
if err := cmd2.Start(); err != nil {
    fmt.Printf("Error: The command can not be startup: %s\n" , err)
    return
}
err := stdin2.Close()
if err != nil {
    fmt.Printf("Error : Can not close the stdio pipe: %s\n" , err)
}

if err := cmd2.Wait(); err != nil {
    fmt.Printf("Error : Can not wait for the command: %s\n" , err)
    return
}

//     up
reader , writer , err := os.Pipe()
n , err := writer.Write(input)
if err != nil {
    fmt.Printf("Error: Can not write data to the named pipe: %s\n " , err)
}
fmt.Printf("Written %d byte(s) . [file-based pipe]\n" , n)
out := make([]byte , 100)
n , err := reader.Read(output)
if err != nil {
    fmt.Printf("Read %d byte(s) . [file-based pipe]\n" , n)
}

