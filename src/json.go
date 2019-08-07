package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	str := `{"servers":
	[{"serverName":"Guangzhou_Base","serverIP":"127.0.0.1"},
	{"serverName":"Beijing_Base","serverIP":"127.0.0.2"}]}`
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerName, s.Servers[1].ServerIp)
}
