package main

import (
    "fmt"
    "os/exec"
    //"io/ioutil"
    //"bytes"
    "os"

)

func exec_command(program string, args ...string) {
    cmd := exec.Command(program, args...)
    
    cmd.Stdout = os.Stdout
    cmd.Stdin = os.Stdin
    cmd.Stderr = os.Stderr


    err := cmd.Start() 
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    
	err = cmd.Wait()

}

func main() {
        fmt.Printf("Start Emulation!")
        arg := os.Args[1]
        exec_command("sudo", "./fat.py", "../../../efi4st/public/uploads/firmware/"+arg)
}