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

func describePrice(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	priceUnit := c.Query("priceUnit")
	period := str2Int32(c.Query("period"))
	price := services.DescribePriceWithOptions(regionID, instanceType, priceUnit, period)

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"price": price,
	}))
}

func createInstance(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	imageID := c.Query("imageId")
	securityGroupID := c.Query("securityGroupId")
	periodUnit := c.Query("priceUnit")
	period := str2Int32(c.Query("period"))
	result, err := services.CreateInstanceWithOptions(regionID, instanceType, imageID, securityGroupID, periodUnit, period)
	if err != nil {
		c.JSON(http.StatusOK, respError(err))
		return
	}

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"data": result,
	}))
}

func describeRecommendInstanceType(c *gin.Context) {
	cores := str2Int32(c.Query("cores"))
	regionID := c.Query("regionId")
	memory := str2Float32(c.Query("memory"))
	rsp := services.DescribeRecommendInstanceTypeWithOptions(regionID, cores, memory)

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

func describeImages(c *gin.Context) {
	regionID := c.Query("regionId")
	fmt.Println("RegionId:", regionID)
	rsp := services.DescribeImagesWithOptions(regionID)

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

func createSecurityGroup(c *gin.Context) {
	regionID := c.Query("regionId")
	rsp := services.CreateSecurityGroup(regionID)

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"security_group_id": rsp.Body.SecurityGroupId,
	}))
}

func str2Int32(s string) int32 {
	n, _ := strconv.Atoi(s)
	num := int32(n)
	return num
}

func str2Float32(s string) float32 {
	ret, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Error(err.Error())
		return 0.00
	}
	return float32(ret)
}

func describeRegions(c *gin.Context) {
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
	tmpl := template.Must(template.ParseFiles("homepage.html"))
	tmpl.Execute(c.Writer, nil)
}
