package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", "./minemind_analysis/main.py", "1", "1234", "CLM_kuC4X", "20", "19", "61.51", "31/01/2020 18:30", "22/7/1996", "2020-11-01", "2020-11-17")
	log.Println(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	cmd.Wait()
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
