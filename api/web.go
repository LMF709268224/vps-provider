package api

import (
	"fmt"
	"html/template"
	"net/http"

	"vps-provider/services"

	"github.com/gin-gonic/gin"
)

func someAction(c *gin.Context) {
	// username := c.Query("username")
	// passStr := c.Query("password")

	// 这里处理按钮点击后的逻辑
	//...

	fmt.Println("someAction-----------")
	http.Redirect(c.Writer, c.Request, "/someAction2", http.StatusOK)
}

func someAction2(c *gin.Context) {
	// username := c.Query("username")
	// passStr := c.Query("password")

	// 这里处理按钮点击后的逻辑
	//...

	fmt.Println("someAction2-----------")
}

func homePage(c *gin.Context) {
	options := services.DescribeRegions()

	myVars := PageVariables{
		RegionIds: options,
	}

	tmpl := template.Must(template.ParseFiles("homepage.html"))
	tmpl.Execute(c.Writer, myVars)
}

type PageVariables struct {
	RegionIds []string
}
