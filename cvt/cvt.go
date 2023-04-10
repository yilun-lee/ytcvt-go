package cvt

import (
	"log"
	"ytcvt/utils"
)

func CvtList(urlList []string, output string) {
	myResErrList := make([]utils.ResultError, len(urlList))
	for _, url := range urlList {
		myResErr := Convert(url, output)
		myResErrList = append(myResErrList, myResErr)
	}
	log.Printf("Successfully convert %d files", len(urlList))
}

func CvtListThread(urlList []string, output string) []utils.ResultError {

	// set output
	anonyFunc := func(url string) utils.ResultError {
		myResErr := Convert(url, output)
		return myResErr
	}

	myResErrList := utils.ThreadMap(urlList, anonyFunc)
	return myResErrList
}
