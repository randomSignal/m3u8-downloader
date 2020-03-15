package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Job struct {
	TsUrl    string
	FilePath string
}

var Queue = make(chan Job, 1000)

func Worker() error {
	for {
		if len(Queue) == 0 {
			break
		}

		job := <-Queue
		_ = TsDownloader(job.TsUrl, job.FilePath)
	}
	return nil
}

func TsDownloader(tsUrl string, filepath string) error {
	fmt.Println("Worker:", tsUrl)
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := httpClient.Get(tsUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	list := strings.Split(tsUrl, "/")
	filename := list[len(list)-1]
	newFilepath := filepath + "/" + filename
	err = ioutil.WriteFile(newFilepath, content, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
