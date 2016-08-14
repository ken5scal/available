package main

import (
	"net"
	"bufio"
	"strings"
)

func exists(domain string) (bool, error) {
	const whoisServer = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43");
	if err != nil {
		return false, err
	}
	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}
	return true, nil
}

func main() {

}
