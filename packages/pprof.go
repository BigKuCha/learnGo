package packages

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func RunPprof() {
	go func() {
		for {
			log.Println(add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("localhost:6061", nil)
}

var datas []string

func add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
