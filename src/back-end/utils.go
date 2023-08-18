package backend

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Util struct{}

// Used to read body from a http.Request and return as a string
func (u *Util) ReadBodyToString(r *http.Request) string {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Read Error  ", err)
	}
	log.Println(len(b), " bytes read from request body.")
	return string(b)
}

func (u *Util) GetValueForKey(s string, k string) string {
	s = s + "&"
	idx1 := strings.Index(s, k)
	updatedString := s[idx1:]

	idx2 := strings.Index(updatedString, "&")

	str := updatedString[len(k)+1 : idx2]
	str = strings.ReplaceAll(str, "+", " ")

	return str
}

func (u *Util) SaveToContentFile(s string) bool {
	fd, err := os.OpenFile("C:\\CACHE_BASEDIR\\todo-app\\todo\\content-store\\contents.json", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}

	n, err := fd.Write([]byte(s + "\n"))
	if err != nil {
		return false
	} else {
		log.Println(n, " bytes written into content file..")
		return true
	}
}
