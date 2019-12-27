package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {

	//Open pid file
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Can't read file")
	}

	if err := os.Remove(pidFile); err != nil {
		//We can continue if this fails
		log.Printf("Warning!! can't remove file - %s", err)
	}

	//convert pid to int
	strPID := strings.TrimSpace(string(data))
	pidInt, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "Can't convert to int, bad process ID")
	}

	fmt.Printf("Killing Pid %d", pidInt)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
