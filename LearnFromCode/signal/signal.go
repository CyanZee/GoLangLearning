package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"errors"
)

func main() {
	go func() {
		sendSignal()
	}()
	
	handler()
}

func sendSignal() {
	cmds := []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "signal"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("grep", "-v", "go run"),
		exec.Command("awk", "{print $2}"),
	}

	output, err := runCmds(cmds)
	if err != nil {
		fmt.Printf("--- Error: failure command execution: %s\n", err)
		return
	}

	pids, err := getPids(output)
	if err != nil {
		fmt.Printf("--- Error: fail to get pids: %s\n", err)
		return
	} else {
		fmt.Printf("+++ Printf pids: %v\n", pids)
	}

	for _, pid := range pids {
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("--- Error: fail to find process: %s\n", err)
			return
		}
		sig := syscall.SIGQUIT
		err = proc.Signal(sig)
		if err != nil {
			fmt.Printf("--- Error: fail to send signal: %s\n", err)
				return
		}
	}
}

func getPids(strs []string) ([]int, error) {
	var pids []int
	for _, str := range strs {
		pid, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}
	return pids, nil
}

func runCmds(cmds []*exec.Cmd) ([]string, error) {
	if cmds == nil || len(cmds) == 0 {
		return nil, errors.New("The cmd slice is invalid!")
	}
	first := true
	var output []byte
	var err error
	for _, cmd := range cmds {
		fmt.Printf("+++ Run command:%v\n",getCmdtext(cmd))
		if !first {
			var stdinBuf bytes.Buffer
			stdinBuf.Write(output)
			cmd.Stdin = &stdinBuf
		}
		var stdoutBuf bytes.Buffer
		cmd.Stdout = &stdoutBuf
		if err = cmd.Start(); err != nil {
			fmt.Printf("--- Error: fail to start the command:%s\n", err)
			return nil, err
		}
		if err = cmd.Wait(); err != nil {
			fmt.Printf("--- Error: Couldn't wait for the first command: %s\n", err)
			return nil, err
		}
		output = stdoutBuf.Bytes()
		if first {
			first = false
		}
	}
	var lines []string
	var outputBuf bytes.Buffer
	outputBuf.Write(output)
	for {
		line, err := outputBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil,err
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func handler() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv1,sigs1...)
	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	signal.Notify(sigRecv2,sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("+++ Received a signal from sigRecv1:%s\n", sig)
		}
		fmt.Printf("+++ End.[sigRecv1]\n")
		wg.Done()
	}()
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("+++ Received a signal from sigRecv2:%s\n", sig)
		}
		fmt.Printf("+++ End.[sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("+++ Wait for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Printf("+++ Stop motification.\n")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Printf("+++ Done.[sigRecv1]\n")
	signal.Stop(sigRecv2)
	close(sigRecv2)
	fmt.Printf("+++ Done.[sigRecv2]\n")
	wg.Wait()
}

func getCmdtext(cmd *exec.Cmd) string {
	var buf bytes.Buffer
	buf.WriteString(cmd.Path)
	for _, arg := range cmd.Args[1:] {
		buf.WriteRune(' ')
		buf.WriteString(arg)
	}
	return buf.String()
}
