package skraper

import (
	"regexp"
)

// func GetConfigsFromChannel(channelName string, offset int) []Config {

// }

func FindLinksInText(text string) []string {
	re := regexp.MustCompile(`vless://[a-z0-9-]+@[a-z0-9.-@]+:\d+\S*`)
	matches := re.FindAllString(text, -1)

	return matches
}

func FindVmessLinkInText(text string) []string {
	re := regexp.MustCompile(`vmess://[a-z0-9-]+\S*`)
	matches := re.FindAllString(text, -1)
	return matches
}

func CheckPreviousMessages(channels []string) {

}
