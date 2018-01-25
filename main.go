package main

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	err := openBrowser("https:fenwickelliott.io")
	if err != nil {
		fmt.Println(err)
	}

}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "darwin":
		exec.Command("open", url).Start()
	default:
		return errors.New("Unrecogniszed os: " + runtime.GOOS)
	}
	return nil
}
