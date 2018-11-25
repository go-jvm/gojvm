package main

import (
	"jvm.com/parseCmd"
	"fmt"
)

func main() {
	cmd := parseCmd.ParseCmd()
	fmt.Println(cmd.VersionFlag)
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		parseCmd.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *parseCmd.Cmd){
	fmt.Printf("classpath: %s class:%s args:%v \n",cmd.CpOption, cmd.Class, cmd.Args)
}
