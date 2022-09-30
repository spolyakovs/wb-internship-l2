package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("enter url to download")
	}

	if len(os.Args) > 3 {
		log.Fatalln("too many arguments")
	}

	urlPath := os.Args[1]
	var filename string
	var err error

	if len(os.Args) == 2 {
		filename, err = getFilename(urlPath)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		filename = os.Args[2]
	}

	resp, err := getResponse(urlPath)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Downloaded a file %s with size %v\n", filename, size)
}

func getFilename(urlPath string) (string, error) {
	urlParsed, err := url.Parse(urlPath)
	if err != nil {
		return "", err
	}
	segments := strings.Split(urlParsed.Path, "/")
	var last string

	for i := len(segments) - 1; i >= 0; i-- {
		last = segments[i]
		if last != "" {
			break
		}
	}
	if last == "" {
		last = urlParsed.Host
	}

	last = strings.Split(last, ".")[0]

	return "wget_" + last, nil
}

func getResponse(urlPath string) (*http.Response, error) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	res, err := client.Get(urlPath)
	if err != nil {
		return nil, err
	}

	return res, nil
}
