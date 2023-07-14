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
	regionId := c.Query("regionId")
	instanceType := c.Query("instanceType")
	priceUnit := c.Query("priceUnit")
	period := Str2Int32(c.Query("period"))
	priceUnit = "Week"
	period = 1
	price := services.DescribePriceWithOptions(regionId, instanceType, priceUnit, period)
	// 这里处理按钮点击后的逻辑
	//...
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"price": price,
	}))
}

func CreateInstance(c *gin.Context) {
	regionId := c.Query("regionId")
	instanceType := c.Query("instanceType")
	imageId := c.Query("imageId")
	securityGroupId := c.Query("securityGroupId")
	periodUnit := c.Query("priceUnit")
	period := Str2Int32(c.Query("period"))
	periodUnit = "Week"
	period = 1
	result, err := services.CreateInstanceWithOptions(regionId, instanceType, imageId, securityGroupId, periodUnit, period)
	// 这里处理按钮点击后的逻辑
	//...
	// fmt.Println("-------------------------------")
	// fmt.Println(result)
	if err != nil {
		c.JSON(http.StatusOK, respError(err))
	}

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"data": result,
	}))
}

func DescribeRecommendInstanceType(c *gin.Context) {
	cores := Str2Int32(c.Query("cores"))
	regionId := c.Query("regionId")
	memory := Str2Float32(c.Query("memory"))
	rsp := services.DescribeRecommendInstanceTypeWithOptions(regionId, cores, memory)
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
			"images": nil,
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
	regionId := c.Query("regionId")
	rsp := services.CreateSecurityGroup(regionId)
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
	rsp := services.DescribeRegionsWithOptions()

	list := make([]string, 0)
	// fmt.Printf("Response: %+v\n", response)
	for _, region := range rsp.Body.Regions.Region {
		// fmt.Printf("Region ID: %s\n", region.RegionId)
		list = append(list, *region.RegionId)
	}

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"images": list,
	}))
}

func homePage(c *gin.Context) {
	// options := services.DescribeRegions()

	// myVars := PageVariables{
	// 	RegionIds: options,
	// }

	tmpl := template.Must(template.ParseFiles("homepage.html"))
	tmpl.Execute(c.Writer, nil)
}

type PageVariables struct {
	RegionIds []string
}
