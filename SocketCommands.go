package main

import (
	"bufio"
	"fmt"
	cm "github.com/daviddengcn/go-colortext"
	"os"
	"strings"
)
var commandmode = true

func FirstStartListening() {
	commandmode = true
	cm.ChangeColor(cm.Green,true,cm.Black,false)
	fmt.Println("----------GoGambling------------")
	fmt.Println("use the command connect {ip} {port} to connect to a lobby server")
	ReadCommand()
}

func ReadCommand() {
	reader := bufio.NewReader(os.Stdin)
	if commandmode {
		_, _ = fmt.Print("- ->")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)
		args := strings.Split(text, " ")
		cmdname := args[0]
		copy(args[0:], args[0+1:]) // Shift a[i+1:] left one index.
		args[len(args)-1] = ""     // Erase last element (write zero value).
		args = args[:len(args)-1]     // Truncate slice.
		cmdHandling(cmdname,args)
	}
}
func StopListening() {
	commandmode = false
}
func RestartListening()  {
	commandmode = true
}
func cmdHandling(name string, args []string)  {
	if name == "connect"{
		ip := args[0]
		port := args[1]
		StartClient(ip,port)
	}
}