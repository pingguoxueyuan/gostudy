package main

import (
	"fmt"
	"strings"
)

var (
	ipWhiteList        []string
	usernameWhiteList  []string
	userIdWhiteList    []string
	useragentWhiteList []string
)

const (
	WhiteListTypeIP        = "ip_whitelist"
	WhiteListTypeUserName  = "username_whitelist"
	WhiteListTypeUserID    = "userid_whitelist"
	WhiteListTypeUserAgent = "useragent_whitelist"
)

func init() {

	count := 150000
	for i := 0; i < count; i++ {
		ip := fmt.Sprintf("IP.IP.IP.%d", i+1)
		ipWhiteList = append(ipWhiteList, ip)
	}

	for i := 0; i < count; i++ {
		userId := fmt.Sprintf("KDKDLWLWW%d", i)
		userIdWhiteList = append(userIdWhiteList, userId)
	}
	for i := 0; i < count; i++ {
		username := fmt.Sprintf("WHITELSLSKDI%d", i)
		usernameWhiteList = append(usernameWhiteList, username)
	}

	for i := 0; i < count; i++ {
		ua := fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36%d", i)
		useragentWhiteList = append(useragentWhiteList, ua)
	}
}

func getWhiteListType(data string) string {

	for _, v := range ipWhiteList {
		if strings.ToLower(v) == strings.ToLower(data) {
			return WhiteListTypeIP
		}
	}

	for _, v := range userIdWhiteList {
		if strings.ToLower(v) == strings.ToLower(data) {
			return WhiteListTypeUserID
		}
	}

	for _, v := range usernameWhiteList {
		if strings.ToLower(v) == strings.ToLower(data) {
			return WhiteListTypeUserName
		}
	}

	for _, v := range useragentWhiteList {
		if strings.ToLower(v) == strings.ToLower(data) {
			return WhiteListTypeUserAgent
		}
	}

	return "unknown white_list"
}

func main() {
	data := "10.237.36.2"
	whitelistType := getWhiteListType(data)

	fmt.Printf("white type:%s\n", whitelistType)
}
