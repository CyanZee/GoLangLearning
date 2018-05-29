package main 

import (
	"os/exec"
	"fmt"
	"bytes"
	"io"
	"bufio"
)

func main() {
	cmd0 := exec.Command("echo","-n","My first command comes from golang.") //as same as "echo -n "My first command comes from golang."

	stdout0, er := cmd0.StdoutPipe()  //create a output pipe to get the command
	if er != nil {                    // "stdout0" 's type is io.ReadCloser
		fmt.Printf("Error: Couldn't obtain the stdout pipe for command No.0:%s\n", er)
	}

	if err := cmd0.Start(); err != nil {   // cmd start
		fmt.Printf("Error:The command No.0 can not be startup:%s\n", err)
	} else {
		fmt.Println("cmd0.Start success.")
	}

	output0 := make([]byte,30)
	n, e := stdout0.Read(output0)   // call stdout0.Read to get output of the command.
	if e != nil {
		fmt.Printf("Error: Couldn't read data from the pipe:%s\n", e)
	}
	fmt.Printf("%s\n", output0[:n])  //

	fmt.Println("=============================")

	var outputBuf0 bytes.Buffer
	for {
		tempOutput := make([]byte, 5)
		n, err := stdout0.Read(tempOutput)
		if err != nil {
			if err == io.EOF {  //if there is nothing to be read in pipe, return err=io.EOF
				break
			} else {
				fmt.Printf("Error: Couldn't read data from the pipe:%s\n", err)
			}
		}
		if n > 0 {
			outputBuf0.Write(tempOutput[:n])  // put "tempOutput" to the buffer "outputBuf0"
		}
	}
	fmt.Printf("%s\n", outputBuf0.String())

	fmt.Println("================================")

	outputBuf1 := bufio.NewReader(stdout0)   // create a reader with buffer.
	output1, _, err := outputBuf1.ReadLine()
	if err != nil {
		fmt.Printf("Error: Couldn't read data from the pipe:%s\n", err)
	}
	fmt.Printf("%s\n",string(output1))
	
	fmt.Println("=====================================")

	cmd1 := exec.Command("ps", "aux")

	var outputBuf2 bytes.Buffer
	cmd1.Stdout = &outputBuf2
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The first command can not be startup %s\n", err)
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the first command: %s\n", err)
	}
//	fmt.Printf("%s\n", outputBuf2.Bytes())


	cmd2 := exec.Command("grep", "bash")
	cmd2.Stdin = &outputBuf2   // as same as "ps aux | grep bash"
	var outputBuf3 bytes.Buffer
	cmd2.Stdout = &outputBuf3
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The second command can not be startup: %s\n", err)
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the second command:%s\n", err)
	}
	fmt.Printf("%s\n", outputBuf3.Bytes())
}
