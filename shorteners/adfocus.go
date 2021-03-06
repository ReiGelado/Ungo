package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Adfocus(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "adfoc.us"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`<a href=\"(.*)" class="skip"`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
