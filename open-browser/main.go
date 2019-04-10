package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	err := openBrowser("http://youtube.com")
	if err != nil {
		log.Println(err)
	}

	fmt.Scan()
}

func openBrowser(urls string) error {

	str, err := exec.LookPath("/opt/google/chrome/chrome")
	if err != nil {
		return err
	}
	err = exec.Command(str, " http://youtube.com", "--start-fullscreen").Start()
	if err != nil {
		return err
	}
	return nil
}
