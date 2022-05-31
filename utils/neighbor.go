package utils

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Hostが通信可能状態で発見できるか
func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)

	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}
	return true
}

// 192.168.0.10:5002
// 192.168.0.10:5003
var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

// 周りのポートを検索する
func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	// マッチした文字列の一部分をスライスで取得
	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}

	prefixHost := m[1]                     // 192.168.0
	lastIp, _ := strconv.Atoi(m[len(m)-1]) //10
	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors
}

// 自信のIPアドレスを取得
func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "127.0.0.1"
	}
	address, err := net.LookupHost(hostname)
	if err != nil {
		return "127.0.0.1"
	}
	return address[0]
}
