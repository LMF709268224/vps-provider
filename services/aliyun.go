package services

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func New() {
	client2, err := ecs.NewClientWithAccessKey("cn-hangzhou", "LTAI5tLczyBS9h4APo9u9c1X", "xHF9hl7qesCYqv8Wp1TLjhDU6EPPjQ")
	if err != nil {
		fmt.Printf("NewClientWithAccessKey err:%s", err.Error())
		return
	}

	request := ecs.CreateDescribePriceRequest()
	request.RegionId = "cn-hangzhou"
	request.InstanceType = "ecs.g5.large"
	// request.ImageId = "alinux_17_01_64_20G_cloudinit_20170818.raw"

	// req := &ecs.DescribePriceRequest{
	// 	RpcRequest: &requests.RpcRequest{},
	// }

	rsp, err := client2.DescribePrice(request)
	if err != nil {
		fmt.Printf("DescribePrice err:%s", err.Error())
		return
	}

	fmt.Println(rsp)
}

func DescribeRegions() []string {
	// 获取区域
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "LTAI5tLczyBS9h4APo9u9c1X", "xHF9hl7qesCYqv8Wp1TLjhDU6EPPjQ")
	if err != nil {
		fmt.Printf("NewClientWithAccessKey err:%s", err.Error())
		return nil
	}

	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https" // optional, set "https" for https access

	response, err := client.DescribeRegions(request)
	if err != nil {
		fmt.Printf("DescribeRegions err:%s", err.Error())
		return nil
	}

	list := make([]string, 0)
	// fmt.Printf("Response: %+v\n", response)
	for _, region := range response.Regions.Region {
		// fmt.Printf("Region ID: %s\n", region.RegionId)
		list = append(list, region.RegionId)
	}

	fmt.Printf("Size: %d\n", len(response.Regions.Region))
	return list
}

func DescribeZones() {
	// 获取区域下的可用区
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "LTAI5tLczyBS9h4APo9u9c1X", "xHF9hl7qesCYqv8Wp1TLjhDU6EPPjQ")
	if err != nil {
		fmt.Printf("NewClientWithAccessKey err:%s", err.Error())
		return
	}

	request := ecs.CreateDescribeZonesRequest()
	request.RegionId = "ap-southeast-3"

	// 发送请求并获取响应
	response, err := client.DescribeZones(request)
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印可用区信息
		for _, zone := range response.Zones.Zone {
			fmt.Println("ZoneId:", zone.ZoneId)
		}
	}
}

func DescribeAvailableResource(regionId string) {
	// 提供的所有实例规格的信息
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "LTAI5tLczyBS9h4APo9u9c1X", "xHF9hl7qesCYqv8Wp1TLjhDU6EPPjQ")
	if err != nil {
		fmt.Printf("NewClientWithAccessKey err:%s", err.Error())
		return
	}

	fmt.Printf("regionId :%s \n", regionId)

	request := ecs.CreateDescribeAvailableResourceRequest()
	request.RegionId = regionId
	request.DestinationResource = "InstanceType"

	response, err := client.DescribeAvailableResource(request)
	if err != nil {
		panic(err)
	}

	for _, zone := range response.AvailableZones.AvailableZone {
		for _, resource := range zone.AvailableResources.AvailableResource {
			// for _, sr := range resource.SupportedResources.SupportedResource {
			// 	fmt.Printf("Value: %s, Status: %v\n", sr.Value, sr)
			// }
			fmt.Printf("Type: %s , SupportedResource size:%d \n", resource.Type, len(resource.SupportedResources.SupportedResource))
		}
		fmt.Printf("ZoneId: %s, Status: %s \n", zone.ZoneId, zone.Status)
	}

	fmt.Println("size:", len(response.AvailableZones.AvailableZone))
}

func DescribeInstanceTypes(regionId string) {
	// 提供的所有实例规格的信息
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "LTAI5tLczyBS9h4APo9u9c1X", "xHF9hl7qesCYqv8Wp1TLjhDU6EPPjQ")
	if err != nil {
		fmt.Printf("NewClientWithAccessKey err:%s", err.Error())
		return
	}

	fmt.Printf("regionId :%s \n", regionId)

	request := ecs.CreateDescribeInstanceTypesRequest()
	request.Scheme = "https" // optional, set "https" for https access
	request.RegionId = regionId

	response, err := client.DescribeInstanceTypes(request)
	if err != nil {
		panic(err)
	}

	for _, instanceType := range response.InstanceTypes.InstanceType {

		fmt.Printf("InstanceType: %s, CPU Core Count: %d, Memory Size: %.fGB\n", instanceType.InstanceTypeId, instanceType.CpuCoreCount, instanceType.MemorySize)
		break
	}

	fmt.Println("size:", len(response.InstanceTypes.InstanceType))
}
