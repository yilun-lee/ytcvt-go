package main

import (
	"flag"
	"ytcvt/cvt"
	"ytcvt/utils"
)

var (
	url      string
	filePath string
	output   string
)

func init() {
	flag.StringVar(&url, "u", "", "The youtube url link")
	flag.StringVar(&filePath, "f", "", "File list containing youtube link")
	flag.StringVar(&output, "o", "/mnt/e/Youtube/", "The output prefix")
}

func getArg() []string {
	flag.Parse()

	var urlList []string
	if url == "" && filePath == "" {
		panic("At least one of -i (url) and -f (fileList) should be specified!")
	} else if url != "" && filePath != "" {
		panic("Only one of -i (url) and -f (fileList) can be specified!")
	} else if url != "" {
		urlList = []string{url}
	} else if filePath != "" {
		urlList = utils.ReadFile(filePath)
	} else {
		panic("no condition for argument meet")
	}
	return urlList
}

func main() {
	urlList := getArg()
	cvt.CvtListThread(urlList, output)
}
