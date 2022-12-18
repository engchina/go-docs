package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	fmt.Println("main start")

	ip := "127.0.0.1:22"
	username := "oracle"
	//exmaple-1: use password
	//password := "oracle"

	//example-2: user ssh key
	keyString, err := os.ReadFile("/home/oracle/.ssh/id_rsa")
	privateKey, err := ssh.ParsePrivateKey(keyString)
	if err != nil {
		panic(err)
	}
	publicKey := ssh.PublicKeys(privateKey)

	//example-1: use password
	//client, err := ssh.Dial("tcp", ip, &ssh.ClientConfig{User: username, Auth: []ssh.AuthMethod{ssh.Password(password)}, HostKeyCallback: ssh.InsecureIgnoreHostKey()})
	//example-2: user ssh key
	client, err := ssh.Dial("tcp", ip, &ssh.ClientConfig{User: username, Auth: []ssh.AuthMethod{publicKey}, HostKeyCallback: ssh.InsecureIgnoreHostKey()})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	session.RequestPty("linux", 32, 160, modes)
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	session.Shell()
	session.Wait()
}
