package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "www.google.com",
		"AWS":    "http://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)
}
