package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("python3", "./minemind_analysis/main.py", "3", "tester01@minemind.net", "CLM_kuC4X", "-1","-1","-1.00","09/01/2021 03:17","22/07/1996","2021-01-04","2021-01-09")
	//cmd := exec.Command("python3", "./minemind_analysis/main.py", "4", "1234", "CLM_kuC4X", "20", "19", "61.51", "31/01/2020 18:30", "22/7/1996", "2020-11-01", "2020-11-17")

	log.Println(cmd)
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	log.Println(err)
	//}
	//err = cmd.Start()
	//if err != nil {
	//	log.Println(err)
	//}
	//scanner := bufio.NewScanner(stdout)
	var result string
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//	result = scanner.Text()
	//}
	//if err := cmd.Wait(); err != nil {
	//	log.Println(err)
	//}
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	result = string(output)
	result = strings.TrimSpace(result)
	log.Println(result)
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
