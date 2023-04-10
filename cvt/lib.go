package cvt

import (
	"log"
	"net/url"
	"ytcvt/utils"
)

func validateUrl(myUrl string) string {
	_, err := url.ParseRequestURI(myUrl)
	utils.PanicOnErr(err)
	return myUrl
}

func Convert(url string, output string) utils.ResultError {

	validateUrl(url)

	var ytdlCmd string = " youtube-dl -i -x --audio-format mp3 -o '{output}%(title)s.%(ext)s' {url} "
	args := map[string]string{
		"output": output,
		"url":    url,
	}
	myNewType := *utils.NewOneLineCmd(ytdlCmd, args)
	out, err := myNewType.Run()

	log.Printf("Successfully download %s", url)
	myResErr := utils.ResultError{Res: out, Err: err}
	return myResErr
}
