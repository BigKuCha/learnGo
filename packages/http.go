package packages

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func TestHttp() {
	Get()
}

func Get() {
	resp, err := http.Get("http://local.yii.dev/site/json")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var body map[string]interface{}
	json.Unmarshal(b, &body)
	io.Copy(os.Stdout, resp.Body)
	fmt.Printf("姓名:%v, 年龄:%v, 工资: %v \n", body["name"], body["age"], body["salary"])
}
