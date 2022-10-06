package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	InValid = "Neither"
	IPV4    = "IPv4"
	IPV6    = "IPv6"
)

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
	fmt.Println(validIPAddress("172.16.0254.1"))
	fmt.Println(validIPAddress("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(validIPAddress("02001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(validIPAddress("2001:0db8:85a3::8A2E:037j:7334"))
	fmt.Println(validIPAddress("2001:db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("1e1.4.5.6"))
	fmt.Println(validIPAddress("20EE:FGb8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("192.0.0.1"))
}

func validIPAddress(queryIP string) string {
	if isValidIPv4(queryIP) {
		return IPV4
	} else if isValidIPv6(queryIP) {
		return IPV6
	} else {
		return InValid
	}

}

func isValidIPv4(queryIP string) bool {
	fragm := strings.Split(queryIP, ".")
	if len(fragm) != 4 {
		return false
	}
	for i := 0; i < len(fragm); i++ {
		fragInt, err := strconv.Atoi(fragm[i])
		if err != nil {
			return false
		}
		digits := []byte(fragm[i])
		if fragInt < 0 || fragInt > 255 || (string(digits[0]) == "0" && len(digits) > 1) {
			return false
		}
	}
	return true
}

func isValidIPv6(queryIP string) bool {
	fragm := strings.Split(queryIP, ":")
	if len(fragm) != 8 {
		return false
	}
	for i := 0; i < len(fragm); i++ {
		if len(fragm[i]) < 1 || len(fragm[i]) > 4 {
			return false
		}
		for _, v := range []byte(fragm[i]) {
			if (v >= 65 && v <= 70) || (v >= 97 && v <= 102) || (v >= 48 && v <= 57) {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
