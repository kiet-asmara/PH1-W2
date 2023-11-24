package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func downloadFile(url string, destination string) error {
	// get data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create file
	out, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer out.Close()

	// write response to file
	_, err = io.Copy(out, resp.Body)

	defer fmt.Println("Downloaded file from", url, "to", destination)
	return err
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	urls := []string{
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_2mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_5mb.mp4",
	}

	for i, u := range urls {
		destination := "video" + strconv.Itoa(i+1) + ".mp4"
		wg.Add(1)
		go func(u string, destination string) {
			defer wg.Done()
			err := downloadFile(u, destination)
			if err != nil {
				fmt.Println("Error downloading file:", err)
				return
			}
		}(u, destination)
	}

	wg.Wait()

	fmt.Println("All files downloaded")
	fmt.Println(time.Since(start))
}
