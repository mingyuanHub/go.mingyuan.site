package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
	"os/exec"
	"strings"
)

func DataDiff(c *gin.Context) {

	var tmplName = "data_diff.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func DataFilter(c *gin.Context) {

	var tmplName = "data_filter.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func DataCalculcator(c *gin.Context) {

	var tmplName = "data_calculator.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

type Calc struct {
	Command string
}

func DataCalc(c *gin.Context) {
	var ca = &Calc{}

	c.ShouldBind(&ca)

	ca.Command = strings.Replace(ca.Command, "x", "*", -1)
	ca.Command = strings.Replace(ca.Command, "X", "*", -1)
	ca.Command = strings.Replace(ca.Command, "÷", "/", -1)

	command := fmt.Sprintf("echo 'scale=6; %s' | bc", ca.Command)

	fmt.Println("[DataCalc] ", command)

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