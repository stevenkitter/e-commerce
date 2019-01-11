package server

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Open() {
	file, err := os.Open("./freeUrl.txt")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	result := string(data)
	rows := strings.Split(result, "\n")
	for i, row := range rows {
		if i < 26 {
			continue
		}
		urls := strings.Split(row, "是:")
		if len(urls) > 1 {
			cmd := exec.Command("open", urls[1])
			err = cmd.Start()
			if err != nil {
				log.Panic(err)
			}
			time.Sleep(5 * time.Second)
		}
	}
}
