package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var myLogger *log.Logger

func main() {
	// 로그파일 오픈
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// out, err := exec.Command("cmd", "/c", "tasklist /v | findstr \"nginx\"").Output()
	// out, err := exec.Command("cmd", "tasklist /v | findstr \"nginx\"").Output()
	out, err := exec.Command("tasklist").Output()
	if err != nil {
		fmt.Printf("error!!!\n")
		myLogger.Println("error", err)
		log.Fatal(err)
	}
	fmt.Printf("The Result is %s\n", out)
	myLogger.Println("result", string(out))

	// nginx 실행중 여부
	var isNginxStarting string = "false"

	// tasklist 중에 nginx.exe 가 있는지 확인하기
	// var isExecuteStop bool = false

	if strings.Contains(string(out), "nginx.exe") {
		// nginx가 구동중이면
		isNginxStarting = "true"
	}

	fmt.Printf(isNginxStarting)
}
