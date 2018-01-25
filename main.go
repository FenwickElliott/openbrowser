package main

import "os/exec"

func main() {
	url := "https:fenwickelliott.io"

	exec.Command("open", url).Start()
}
