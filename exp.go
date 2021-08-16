package main

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net"
	"path"
)

func Pallete() {
	var (
		blue    = "0x0000ff"
		cyan    = "0x00FFff"
		crimson = "0xdC143c"
	)
	fmt.Printf("blue = %s, cyan = %s, crimson = %s\n", blue, cyan, crimson)
}

func IPAppender() {
	ips := make([]net.IP, 0)

	ips = append(ips, net.IPv4bcast)
	ips = append(ips, net.IPv4allsys)
	ips = append(ips, net.IPv4allrouter)
	ips = append(ips, net.IPv4zero)

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func GetFileExtensionList() {
	var s []string
	s = append(s, ".go")
	s = append(s, ".py")
	s = append(s, ".js")
}

func IsEmpty(path string) bool {
	return len(path) == 0
}

func WriteB64Password(filename string, password []byte) {
	data := []byte(base64.StdEncoding.EncodeToString(password))
	perm := fs.FileMode(0o644)
	if err := ioutil.WriteFile(path.Join("tmp", filename), data, perm); err != nil {
		panic(err)
	}
}
