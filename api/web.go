package api

import (
	"html/template"
	"net/http"

	"vps-provider/utils"

	"vps-provider/services"

	"github.com/gin-gonic/gin"
)

func describePrice(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	priceUnit := c.Query("priceUnit")
	period := utils.Str2Int32(c.Query("period"))
	price, err := services.DescribePriceWithOptions(regionID, instanceType, priceUnit, period)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		return
	}

	c.JSON(http.StatusOK, respJSON(http.StatusOK, price))
}

func createInstance(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	imageID := c.Query("imageId")
	securityGroupID := c.Query("securityGroupId")
	periodUnit := c.Query("priceUnit")
	period := utils.Str2Int32(c.Query("period"))

	if periodUnit == "Year" {
		periodUnit = "Month"
		period = period * 12
	}

	result, err := services.CreateInstance(regionID, instanceType, imageID, securityGroupID, periodUnit, period)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
		return
	}

	c.JSON(http.StatusOK, respJSON(http.StatusOK, result))
}

func describeRecommendInstanceType(c *gin.Context) {
	cores := utils.Str2Int32(c.Query("cores"))
	regionID := c.Query("regionId")
	memory := utils.Str2Float32(c.Query("memory"))
	rsp, err := services.DescribeRecommendInstanceTypeWithOptions(regionID, cores, memory)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
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
	c.JSON(http.StatusOK, respJSON(http.StatusOK, rpsData))
}

func describeImages(c *gin.Context) {
	regionID := c.Query("regionId")
	instanceType := c.Query("instanceType")
	// fmt.Println("RegionId:", regionID)
	rsp, err := services.DescribeImagesWithOptions(regionID, instanceType)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
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

	c.JSON(http.StatusOK, respJSON(http.StatusOK, rsp.Body.SecurityGroupId))
}

func describeAvailableResource(c *gin.Context) {
	regionID := c.Query("regionId")
	cores := utils.Str2Int32(c.Query("cores"))
	memory := utils.Str2Float32(c.Query("memory"))
	rsp, err := services.DescribeAvailableResourceWithOptions(regionID, cores, memory)
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
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
	rsp, err := services.DescribeRegionsWithOptions()
	if err != nil {
		data := utils.StrToMap(*err.Data)
		c.JSON(http.StatusOK, respJSON(*err.StatusCode, jsonObject{
			"msg":     err.Code,
			"details": data["Message"],
		}))
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
