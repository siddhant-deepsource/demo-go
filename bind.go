package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}

func makeSSHTunnel(user string, signer ssh.Signer, host string) (*SSHTunnel, error) {
	config := ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}
