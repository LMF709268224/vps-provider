package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"vps-provider/services"

	"github.com/gin-gonic/gin"
)

type JsonObject map[string]interface{}

func someAction(c *gin.Context) {
	// username := c.Query("username")
	// passStr := c.Query("password")

	// 这里处理按钮点击后的逻辑
	//...

	fmt.Println("someAction-----------")
	http.Redirect(c.Writer, c.Request, "/someAction2", http.StatusOK)
}

func DescribePrice(c *gin.Context) {
	RegionId := c.Query("regionId")
	InstanceType := c.Query("instanceType")
	price := services.DescribePriceWithOptions(RegionId, InstanceType)
	// 这里处理按钮点击后的逻辑
	//...
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"price": price,
	}))
}

func CreateInstance(c *gin.Context) {
	RegionId := c.Query("regionId")
	InstanceType := c.Query("instanceType")
	ImageId := c.Query("imageId")
	SecurityGroupId := c.Query("securityGroupId")
	err := services.CreateInstanceWithOptions(RegionId, InstanceType, ImageId, SecurityGroupId)
	// 这里处理按钮点击后的逻辑
	//...
	if err != nil {
		c.JSON(http.StatusOK, respJSON(JsonObject{
			"data": "create failed",
		}))
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"price": "success",
	}))
}

func DescribeRecommendInstanceType(c *gin.Context) {
	Cores := Str2Int32(c.Query("cores"))
	RegionId := c.Query("regionId")
	Memory := Str2Float32(c.Query("memory"))
	rsp := services.DescribeRecommendInstanceTypeWithOptions(RegionId, Cores, Memory)
	// 这里处理按钮点击后的逻辑
	//...
	if rsp == nil {
		c.JSON(http.StatusOK, respJSON(JsonObject{
			"data": nil,
		}))
		return
	}
	var rpsData []string
	for _, data := range rsp.Body.Data.RecommendInstanceType {
		instanceType := data.InstanceType.InstanceType
		if *instanceType == "" {
			continue
		}
		rpsData = append(rpsData, *instanceType)
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"data": rpsData,
	}))
}

func DescribeImages(c *gin.Context) {
	regionId := c.Query("regionId")
	fmt.Println("RegionId:", regionId)
	rsp := services.DescribeImagesWithOptions(regionId)

	if rsp == nil {
		c.JSON(http.StatusOK, respJSON(JsonObject{
			"data": nil,
		}))
		return
	}
	var rpsData []string
	for _, data := range rsp.Body.Images.Image {
		instanceType := data.ImageId
		if *instanceType == "" {
			continue
		}
		rpsData = append(rpsData, *instanceType)
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"images": rpsData,
	}))
}

func CreateSecurityGroup(c *gin.Context) {
	RegionId := c.Query("regionId")
	rsp := services.CreateSecurityGroup(RegionId)
	// 这里处理按钮点击后的逻辑
	//...
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"security_group_id": rsp.Body.SecurityGroupId,
	}))
}

func Str2Int32(s string) int32 {
	n, _ := strconv.Atoi(s)
	num := int32(n)
	return num
}

func Str2Float32(s string) float32 {
	ret, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Error(err.Error())
		return 0.00
	}
	return float32(ret)
}

func DescribeRegions(c *gin.Context) {
	list := services.DescribeRegions()

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"images": list,
	}))
}

func homePage(c *gin.Context) {
	// options := services.DescribeRegions()

	// myVars := PageVariables{
	// 	RegionIds: options,
	// }

	tmpl := template.Must(template.ParseFiles("homepage2.html"))
	tmpl.Execute(c.Writer, nil)
}

type PageVariables struct {
	RegionIds []string
}
