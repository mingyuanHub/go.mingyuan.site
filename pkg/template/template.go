package template

import (
	"html/template"
	"io"
	"mingyuanHub/mingyuan.site/pkg/logger"
)

type MyTemplate struct {
	*template.Template
}

func ParseFiles(file string) (t *MyTemplate, err error){
	t = new(MyTemplate)

	t.Template, err = template.ParseFiles("my-web/views/_header.html", "my-web/views/_footer.html", file)

	return
}

func (t MyTemplate) ExecuteTemplate(wr io.Writer, name string, data interface{}) {
	err := t.Template.ExecuteTemplate(wr, name, data)

	if err != nil {
		logger.Error("fail to ExecuteTemplate", err.Error())
	}
}
