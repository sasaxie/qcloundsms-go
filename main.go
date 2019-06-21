package main

import (
	"fmt"
	"github.com/sasaxie/qcloundsms-go/util"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {

		s, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		f := util.Sha1Sum(s)
		fmt.Println(f)
	})
	http.ListenAndServe("127.0.0.1:8000", nil)
}
