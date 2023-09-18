package skraper

import (
	"regexp"
)

// func GetConfigsFromChannel(channelName string, offset int) []Config {

// }

func FindLinks(text string) []string {
	re := regexp.MustCompile(`vless:\/\/[a-zA-Z0-9%\-]+@[a-zA-Z0-9\.\-]+:[0-9]+[^\s]+`)
	matches := re.FindAllString(text, -1)

	return matches
}

func FindVmessLinks(text string) []string {
	re := regexp.MustCompile(`vmess://[a-z0-9-]+\S*`)
	matches := re.FindAllString(text, -1)
	return matches
}

func FindSSLinks(text string) []string {
	re := regexp.MustCompile(`ss:\/\/[A-Za-z0-9\-_]+@[A-Za-z0-9\.\-]+:\d+#[^\n]+`)
	matches := re.FindAllString(text, -1)
	return matches
}

func FindTrojanLinks(text string) []string {
	re := regexp.MustCompile(`trojan:\/\/[a-zA-Z0-9%\-]+@[a-zA-Z0-9\.\-]+:[0-9]+[^\s]+`)
	matches := re.FindAllString(text, -1)
	return matches
}

func ExtractLinksFromText(text string) []string {
	links := make([]string, 0)

	vmessLinks := FindVmessLinks(text)
	vlessLinks := FindLinks(text)
	trojanLinks := FindTrojanLinks(text)
	// ssLinks := FindSSLinks(text)

	links = append(links, vmessLinks...)
	links = append(links, vlessLinks...)
	links = append(links, trojanLinks...)
	// links = append(links, ssLinks...)

	return links
}
