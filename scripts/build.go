package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Start to build project")

	execute("pwd")

	logMessage("Start formatting whole project")
	execute("gofmt -w -s .")

	logMessage("Start building project")
	execute("go build")

	logMessage("Start app")
	start("./go-rest-api.exe")

	logMessage("Go rest api server started")
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

func start(app string) {
	cmd := exec.Command(app)
	err := cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func logMessage(message string) {
	fmt.Println(message)
}
