package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Start to run unit tests")

	execute("pwd")
	execute("go test ./... -v")

	fmt.Println("Finish run unit tests")
}

func execute(command string) {
	if len(command) == 0 {
		return
	}

	cmd := exec.Command("powershell", "-c", command)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
