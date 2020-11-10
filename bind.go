package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"html"
	"net"
)

func insecureIgnoreHostKey() {
	_ = &ssh.ClientConfig{
		User:            "username",
		Auth:            []ssh.AuthMethod{nil},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func insecureHostKeyCallback() {
	_ = &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{nil},
		HostKeyCallback: ssh.HostKeyCallback(
			func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			}),
	}
}

func htmlMustEscape() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
}
