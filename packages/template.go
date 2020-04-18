package packages

import (
	"fmt"
	"github.com/urfave/cli"
	"html/template"
	"os"
)

type Body struct {
	Content string
}
type Footer struct {
	Content string
}

func Tpl(ctx *cli.Context) {
	t, err := template.ParseFiles("./packages/template.html")
	if err != nil {
		panic(err)
	}
	b := Body{Content: "这是body"}
	f := Footer{Content: "这是footer"}
	Data := struct {
		Body   Body
		Footer Footer
	}{
		Body:   b,
		Footer: f,
	}

	err = t.Execute(os.Stdout, Data)
	if err != nil {
		fmt.Println(err)
	}
}
