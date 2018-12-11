package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	//"strings"
)

type Service struct {
	XMLName       xml.Name `xml:"service"`
	Id            string   `xml:"id"`
	Name          string   `xml:"name"`
	Description   string   `xml:"description"`
	Logpath       string   `xml:"logpath"`
	Synchronous   string   `xml:"synchronous"`
	Auto          string   `xml:"auto"`
	Timeout       string   `xml:"timeout"`
	Startargument string   `xml:"startargument"`
	Stopargument  string   `xml:"stopargument"`
}

func ReadConfig(confFile string) Service {
	confFile = "E:\\svn\\tw6\\products\\apache-tomee\\target\\tongweb-webprofile-1.5.0\\service\\conf\\twnt.xml"
	content, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}
	var result Service
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetStartargument(confFile string) string {
	service := ReadConfig(confFile)
	return service.Startargument
}

func GetStopargument(confFile string) string {
	service := ReadConfig(confFile)
	return service.Stopargument
}

func GetTimeout(confFile string) int {
	service := ReadConfig(confFile)
	timeout, err := strconv.Atoi(service.Timeout)
	if err == nil {
		return timeout
	}
	return -1
}

func GetId(confFile string) string {
	service := ReadConfig(confFile)
	return service.Id
}

func GetLogpath(confFile string) string {
	service := ReadConfig(confFile)
	return service.Logpath
}

func GetDescription(confFile string) string {
	service := ReadConfig(confFile)
	return service.Description
}

func GetSynchronous(confFile string) string {
	service := ReadConfig(confFile)
	return service.Synchronous
}

func GetAuto(confFile string) string {
	service := ReadConfig(confFile)
	return service.Auto
}

func GetName(confFile string) string {
	service := ReadConfig(confFile)
	return service.Name
}

var contentArray = make([]string, 0, 5)

func main() {
	//log.Println("-------"+GetId(""))
	//log.Println("-------"+GetName(""))
	//log.Println("-------"+GetDescription(""))
	//log.Println("-------"+GetLogpath(""))
	//log.Println("-------"+GetSynchronous(""))
	//log.Println("-------"+GetAuto(""))
	//log.Println(GetTimeout(""))
	//log.Println("-------"+GetStartargument(""))
	//log.Println("-------"+GetStopargument(""))

	//var startargArray []string;
	//startargArray=strings.Split(GetStartargument("")," ")
	//var javaline string
	//javaline = strings.Replace(startargArray[0],"\"","",-1)
	//	argslice := startargArray[1:len(startargArray)]
	//	argsline :=strings.Join(argslice," ")
	//argsline=strings.Replace(argsline,"\n","",-1)

	//command := flag.String("cmd /c", javaline,"Set the command.")
	//args := flag.String("args", argsline, "Set the args. (separated by spaces)")
	//args := flag.String("args", "-jar ProcessLaucher.jar", "Set the args. (separated by spaces)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-cmd <command>] [-args <the arguments (separated by spaces)>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	//fmt.Println("Command: ", *command)
	//fmt.Println("Arguments: ", *args)
	//var argArray []string
	//if *args != "" {
	//	argArray = strings.Split(*args, " ")
	//} else {
	//	argArray = make([]string, 0)
	//}

	//env := os.Environ()
	HOME := os.Getenv("HOMEPATH")
	fmt.Println("----home:" + HOME)
	fmt.Println("-----java home:" + os.Getenv("JAVA_HOME"))
	contentArray = contentArray[0:0]

	//fmt.Println("env: ", env)
	//cmd := exec.Command(*command, argArray...)
	os.Chdir("E:\\svn\\tw6\\products\\apache-tomee\\target\\tongweb-webprofile-1.5.0\\bin\\")
	cmd := exec.Command("cmd", "/c", "E:\\svn\\tw6\\products\\apache-tomee\\target\\tongweb-webprofile-1.5.0\\bin\\startserver.bat")
	//cmd := exec.Command("cmd","/c","tasklist")
	stdout, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
		return
	}
	stderr, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
		return
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	readererr := bufio.NewReader(stderr)

	var index int
	//实时循环读取输出流中的一行内容
	go func() {
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println("stdout----" + line)
			index++
			//contentArray = append(contentArray, line)
		}
	}()

	var index2 int
	go func() {
		for {
			line, err2 := readererr.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println("stderr-----" + line)
			index2++
			//contentArray = append(contentArray, line)
		}
	}()

	//stdoutBuf, err := cmd.Output()

	//fmt.Println("---------"+string(stdoutBuf))
	cmd.Wait()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, *cmd, cmd.Args)
	//	return
	//}
	//enc := mahonia.NewEncoder("UTF8")
	//fmt.Fprintf(os.Stdout, "Result: %s", enc.ConvertString(string(stdoutBuf)))
}
