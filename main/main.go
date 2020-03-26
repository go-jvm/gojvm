package main

import "fmt"

func main() {
	cmd := parserCmd()
	if cmd.versionFlag {
		fmt.Println("version 1.8.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath: %s class: %s, args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}