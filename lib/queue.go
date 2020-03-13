package lib

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Job struct {
	TsUrl    string
	FilePath string
}

var Queue = make(chan Job, 1000)

func Worker() error {
	job, ok := <-Queue
	if !ok {
		return nil
	}
	resp, err := http.Get(job.TsUrl)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	list := strings.Split(job.TsUrl, "/")
	filename := list[len(list)-1]
	newFilepath := job.FilePath + "/" + filename
	err = ioutil.WriteFile(newFilepath, content, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
