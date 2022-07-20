package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: waybard-nordvpn <exec|toggle>\n")
	os.Exit(1)
}

type status struct {
	isConnected bool
	output      string
}

func getStatus() status {
	cmd := exec.Command("nordvpn", "status")
	data, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
	out_str := string(data)
	out_str = out_str[strings.Index(out_str, "Status:"):]

	var isConnected bool
	if strings.Contains(out_str, "Status: Connected") {
		isConnected = true
	} else {
		isConnected = false
	}
	return status{isConnected: isConnected, output: out_str}
}

type ExecOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	// Class string `json::"class"`
}

func getExecOutput(status status) ExecOutput {
	var text string
	if status.isConnected {
		text = "  "
	} else {
		text = "  "
	}
	return ExecOutput{text, status.output}
}

func handleExec(status status) {
	execOutput := getExecOutput(status)
	res, err := json.Marshal(execOutput)
	if err != nil {
		log.Fatalf("json.Marshal() failed: %v", err)
	}
	fmt.Println(string(res))
}

func handleToggle(status status) {
	var cmdAction string
	if status.isConnected {
		cmdAction = "disconnect"
	} else {
		cmdAction = "connect"
	}
	cmd := exec.Command("nordvpn", cmdAction)
	cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	status := getStatus()
	switch os.Args[1] {
	case "exec":
		handleExec(status)
	case "toggle":
		handleToggle(status)
	default:
		usage()
	}
}
