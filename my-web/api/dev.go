package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
	"os/exec"
)

func DevJson2GoStruct(c *gin.Context) {

	var tmplName = "dev_json2gostruct.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

type Request struct {
	Json string
}

func DevJson2GoStructApi(c *gin.Context) {
	var req = &Request{}

	c.ShouldBind(&req)

	command := fmt.Sprintf("echo '%s' | gojson -name=Request", req.Json)

	fmt.Println("[DevJson2GoStructApi] ", command)

	out, err := exec.Command("bash", "-c", command).Output()

	var result string
	if err != nil {
		result = err.Error()
	} else {
		result = string(out)
	}

	response := struct {
		Status int    `json:"status"`
		Result string `json:"result"`
	}{
		200,
		result,
	}
	c.JSONP(http.StatusOK, response)
}