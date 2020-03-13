package main

import (
	"flag"
	"fmt"
	"github.com/randomSignal/m3u8-downloader/lib"
	"os"
	"sync"
	"time"
)

// 实际中应该用更好的变量名
var (
	h        bool
	url      string
	filepath string
	thread   int
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&url, "url", "", "the m3u8 url")
	flag.StringVar(&filepath, "filepath", "./", "the m3u8 write filepath")
	flag.IntVar(&thread, "thread", 10, "num of download thread")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h || url == "" {
		flag.Usage()
		return
	}

	err := lib.WriteQueue(url, filepath)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("There are ", len(lib.Queue), " left to download ")
		}
	}()

	var wg sync.WaitGroup

	for i := 0; i < thread; i++ {
		go func() {
			err = lib.Worker()
			wg.Done()
		}()
		wg.Add(1)
	}
	wg.Wait()

	if err != nil {
		panic(err)
	}
}

func usage() {

	fmt.Fprintf(os.Stderr, `Usage: m3u8-downloader  [-url m3u8_url] [-filepath write_path] 

Options:
`)
	flag.PrintDefaults()

}
