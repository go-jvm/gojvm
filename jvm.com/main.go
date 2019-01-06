package main

import (
	"jvm.com/parseCmd"
	"fmt"
	"jvm.com/classpath"
	"strings"
	"jvm.com/classfile"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.Class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile{
	classData, _ , err := cp.ReadClass(className)

	if err!= nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)

	if err!= nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version：%v.%v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constants count：%v\n",len(cf.ConstantPoll()))
	fmt.Printf("access flags：0x%x\n",cf.AccessFlags())
	fmt.Printf("this class：%v\n",cf.ClassName())
	fmt.Printf("super class：%v\n",cf.SuperClassName())
	fmt.Printf("interfaces：%v\n",len(cf.Fields()))
	for _, f := range cf.Fields(){
		fmt.Printf("%s\n",f.Name())
	}

	fmt.Printf("methods count: %v\n",len(cf.Methods()))

	for _,m := range cf.Methods() {
		fmt.Printf("  %s\n",m.Name())
	}
}
