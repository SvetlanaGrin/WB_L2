package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//folder, _ := os.UserHomeDir()
	folder, _ := os.Getwd()
	//folder, _ := os.Executable()
	for {
		fmt.Print(folder)
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if str, flag, err := execInput(folder, input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else if flag == 0 {
			folder = str
		} else if flag == 1 || flag == 2 {
			fmt.Println(str)
		}

	}
}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

func Regexp(folder, arg string) (string, int, error) {
	if m, _ := regexp.MatchString(".*:\\.*", arg); m {
		return arg, 0, os.Chdir(arg)
	} else if e, _ := regexp.MatchString(`^..$`, arg); e {
		strs := strings.Split(folder, "\\")
		return strings.Join(strs[:(len(strs)-1)], "\\"), 0, os.Chdir(strings.Join(strs[:(len(strs)-1)], "\\"))
	} else {
		return folder + "\\" + arg, 0, os.Chdir(folder + "\\" + arg)
	}
}
func KillProcess(name string) error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			fmt.Println("Not name", " ", p.Pid)
		}
		if n == name {
			return p.Kill()
		}
	}
	return fmt.Errorf("process not found")
}
func Path(folder string) (string, int, error) {
	str, err := os.Getwd()
	return str, 1, err
}
func PSProcess() error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			continue
		}
		time1, err := p.CreateTime()
		if time1 == 0 {
			continue
		}
		t := time.Unix(0, time1*int64(time.Millisecond))
		tOnly := t.Format(time.TimeOnly)
		if err != nil { // Always check errors even if they should not happen.
			panic(err)
		}
		if err != nil {
			fmt.Print("Not Times", " ")
		}
		cmd, err := p.Cmdline()
		if err != nil {

		}
		fmt.Println(p.Pid, " ", n, " ", tOnly, " ", cmd, " ")
	}
	return nil
}
func execInput(folder, input string) (string, int, error) {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\r\n")

	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return "", 0, ErrNoPath
		}
		return Regexp(folder, args[1])
	case "pwd":
		return Path(folder)
		// Change the directory and return the  error.
		//return os.Chdir(args[1]), args[1]
	case "echo":
		return args[1], 2, nil
	case "kill":
		if len(args) < 2 {
			return "", 0, io.ErrNoProgress
		}
		err := KillProcess(args[1])
		return "", 3, err
	case "ps":
		err := PSProcess()
		return "", 4, err
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return args[1], 0, cmd.Run()
}
