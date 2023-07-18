package api

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"vps-provider/utils"

	"vps-provider/services"

	"github.com/gin-gonic/gin"
)

func describePrice(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	priceUnit := c.Query("priceUnit")
	period := utils.Str2Int32(c.Query("period"))
	imageID := c.Query("imageId")
	price, err := services.DescribePrice(regionID, instanceType, priceUnit, imageID, period)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		fmt.Println("DescribePrice err:", err.Error())
		return
	}

	c.JSON(http.StatusOK, respJSON(http.StatusOK, price))
}

func createInstance(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	imageID := c.Query("imageId")
	periodUnit := c.Query("priceUnit")
	password := c.Query("password")
	period := utils.Str2Int32(c.Query("period"))

	if periodUnit == "Year" {
		periodUnit = "Month"
		period = period * 12
	}

	var securityGroupID string

	group, err := services.DescribeSecurityGroups(regionID)
	if err == nil && len(group) > 0 {
		securityGroupID = group[0]
	} else {
		securityGroupID, err = services.CreateSecurityGroup(regionID)
		if err != nil {
			data := utils.StrToMap(*err.Data)
			c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
				"msg":     err.Code,
				"details": data["Message"],
			}))
			fmt.Println("CreateSecurityGroup err:", err.Error())
			return
		}
	}

	fmt.Println("securityGroupID : ", securityGroupID)

	result, err := services.CreateInstance(regionID, instanceType, imageID, password, securityGroupID, periodUnit, period)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		fmt.Println("CreateInstance err:", err.Error())
		return
	}

	address, err := services.AllocatePublicIpAddress(regionID, result.InstanceId)
	if err != nil {
		fmt.Println("AllocatePublicIpAddress err:", err.Error())
	} else {
		result.PublicIpAddress = address
	}

	err = services.AuthorizeSecurityGroup(regionID, securityGroupID)
	if err != nil {
		fmt.Println("AuthorizeSecurityGroup err:", err.Error())
	}

	// 一分钟后调用
	go func() {
		time.Sleep(1 * time.Minute)

		err := services.StartInstance(regionID, result.InstanceId)
		fmt.Println("StartInstance err:", err)
	}()

	c.JSON(http.StatusOK, respJSON(http.StatusOK, result))
}

func describeRecommendInstanceType(c *gin.Context) {
	cores := utils.Str2Int32(c.Query("cores"))
	regionID := c.Query("regionId")
	memory := utils.Str2Float32(c.Query("memory"))
	rsp, err := services.DescribeRecommendInstanceType(regionID, cores, memory)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		fmt.Println("DescribeRecommendInstanceType err:", err.Error())
		return
	}

	fmt.Println("DescribeRecommendInstanceType rsp:", rsp)
	resources := make(map[string]string)
	for _, data := range rsp.Body.Data.RecommendInstanceType {
		instanceType := data.InstanceType.InstanceType
		if *instanceType == "" {
			continue
		}
		resources[*instanceType] = *instanceType
	}

	var rpsData []string
	for value := range resources {
		rpsData = append(rpsData, value)
	}

	c.JSON(http.StatusOK, respJSON(http.StatusOK, rpsData))
}

func describeImages(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	// fmt.Println("RegionId:", regionID)
	rsp, err := services.DescribeImages(regionID, instanceType)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		fmt.Println("DescribeImages err:", err.Error())
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
	c.JSON(http.StatusOK, respJSON(http.StatusOK, rpsData))
}

func createSecurityGroup(c *gin.Context) {
	regionID := c.Query("regionId")
	rsp, err := services.CreateSecurityGroup(regionID)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		return
	}

	c.JSON(http.StatusOK, respJSON(http.StatusOK, rsp))
}

func describeAvailableResource(c *gin.Context) {
	regionID := c.Query("regionId")
	cores := utils.Str2Int32(c.Query("cores"))
	memory := utils.Str2Float32(c.Query("memory"))
	rsp, err := services.DescribeAvailableResource(regionID, cores, memory)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		fmt.Println("DescribeAvailableResource err:", err.Error())
		return
	}

	resources := make(map[string]string)
	for _, data := range rsp.Body.AvailableZones.AvailableZone {
		for _, resource := range data.AvailableResources.AvailableResource {
			for _, sr := range resource.SupportedResources.SupportedResource {
				resources[*sr.Value] = *sr.Status
			}
		}
	}

	var rpsData []string
	for value := range resources {
		rpsData = append(rpsData, value)
	}

	c.JSON(http.StatusOK, respJSON(http.StatusOK, rpsData))
}

func describeRegions(c *gin.Context) {
	rsp, err := services.DescribeRegions()
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		fmt.Println("DescribeRegions err:", err.Error())
		return
	}

	var rpsData []string
	// fmt.Printf("Response: %+v\n", response)
	for _, region := range rsp.Body.Regions.Region {
		// fmt.Printf("Region ID: %s\n", region.RegionId)
		rpsData = append(rpsData, *region.RegionId)
	}

	c.JSON(http.StatusOK, respJSON(http.StatusOK, rpsData))
}

func homePage(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("homepage.html"))
	tmpl.Execute(c.Writer, nil)
}
