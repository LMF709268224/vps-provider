package services

import (
	"fmt"

	"vps-provider/config"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param config.Cfg.AliyunAccessKeyID
 * @param config.Cfg.AliyunAccessKeySecret
 * @return Client
 * @throws Exception
 */
func ClientInit() error {
	configClient := &openapi.Config{
		AccessKeyId:     tea.String(config.Cfg.AliyunAccessKeyID),
		AccessKeySecret: tea.String(config.Cfg.AliyunAccessKeySecret),
	}
	configClient.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	result := &ecs20140526.Client{}
	result, err := ecs20140526.NewClient(configClient)
	AliClient = result
	return err
}

type Client struct {
	ecs20140526.Client
}

var AliClient *ecs20140526.Client

func CreateInstance(RegionId, InstanceType, ImageId, SecurityGroupId, PeriodUnit string, Period int32) (*ecs20140526.CreateInstanceResponse, *tea.SDKError) {
	var result *ecs20140526.CreateInstanceResponse

	createInstanceRequest := &ecs20140526.CreateInstanceRequest{
		RegionId:           tea.String(RegionId),
		InstanceType:       tea.String(InstanceType),
		DryRun:             tea.Bool(true),
		ImageId:            tea.String(ImageId),
		SecurityGroupId:    tea.String(SecurityGroupId),
		InstanceChargeType: tea.String("PrePaid"),
		PeriodUnit:         tea.String(PeriodUnit),
		Period:             tea.Int32(Period),
	}
	var err error
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err = AliClient.CreateInstanceWithOptions(createInstanceRequest, runtime)
		if err != nil {
			fmt.Errorf("CreateInstanceWithOptions %v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return result, errors
	}
	return result, nil
}

func DescribePriceWithOptions(RegionId, InstanceType, PriceUnit string, Period int32) (*ecs20140526.DescribePriceResponseBodyPriceInfoPrice, *tea.SDKError) {
	var price *ecs20140526.DescribePriceResponseBodyPriceInfoPrice
	describePriceRequest := &ecs20140526.DescribePriceRequest{
		RegionId:     tea.String(RegionId),
		InstanceType: tea.String(InstanceType),
		PriceUnit:    tea.String(PriceUnit),
		Period:       tea.Int32(Period),
	}
	runtime := &util.RuntimeOptions{}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err := AliClient.DescribePriceWithOptions(describePriceRequest, runtime)
		if err != nil {
			fmt.Errorf("DescribePriceWithOptions:%v", err)
			return err
		}
		price = result.Body.PriceInfo.Price
		return nil
	}()
	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return price, errors
	}
	return price, nil
}

func DescribeRegionsWithOptions() (*ecs20140526.DescribeRegionsResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeRegionsResponse
	describeRegionsRequest := &ecs20140526.DescribeRegionsRequest{}
	runtime := &util.RuntimeOptions{}
	var err error
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err = AliClient.DescribeRegionsWithOptions(describeRegionsRequest, runtime)
		if err != nil {
			return err
		}
		return nil
	}()
	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return result, errors
	}
	return result, nil
}

func DescribeRecommendInstanceTypeWithOptions(RegionId string, Cores int32, Memory float32) (*ecs20140526.DescribeRecommendInstanceTypeResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeRecommendInstanceTypeResponse
	var err error
	describeRecommendInstanceTypeRequest := &ecs20140526.DescribeRecommendInstanceTypeRequest{
		// NetworkType:        tea.String("vpc"),
		RegionId:           tea.String(RegionId),
		Cores:              tea.Int32(Cores),
		Memory:             tea.Float32(Memory),
		InstanceChargeType: tea.String("PrePaid"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err = AliClient.DescribeRecommendInstanceTypeWithOptions(describeRecommendInstanceTypeRequest, runtime)
		if err != nil {
			fmt.Errorf("DescribeRecommendInstanceTypeWithOptions %v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		errString := util.AssertAsString(errors.Message)
		if errString != nil {
			fmt.Println(*errString)
		}
		return result, errors
	}
	return result, nil
}

func CreateSecurityGroup(RegionId string) (*ecs20140526.CreateSecurityGroupResponse, *tea.SDKError) {
	var result *ecs20140526.CreateSecurityGroupResponse
	var err error

	createSecurityGroupRequest := &ecs20140526.CreateSecurityGroupRequest{
		RegionId: tea.String(RegionId),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err = AliClient.CreateSecurityGroupWithOptions(createSecurityGroupRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return result, errors
	}
	return result, nil
}

func DescribeImagesWithOptions(RegionId string) (*ecs20140526.DescribeImagesResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeImagesResponse
	var err error

	createSecurityGroupRequest := &ecs20140526.DescribeImagesRequest{
		RegionId: tea.String(RegionId),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err = AliClient.DescribeImagesWithOptions(createSecurityGroupRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return result, errors
	}
	return result, nil
}

func DescribeAvailableResourceWithOptions(RegionId string, cores int32, memory float32) (*ecs20140526.DescribeAvailableResourceResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeAvailableResourceResponse
	var err error
	describeAvailableResourceRequest := &ecs20140526.DescribeAvailableResourceRequest{
		RegionId:            tea.String(RegionId),
		DestinationResource: tea.String("InstanceType"),
		InstanceChargeType:  tea.String("PrePaid"),
		Cores:               tea.Int32(cores),
		Memory:              tea.Float32(memory),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err = AliClient.DescribeAvailableResourceWithOptions(describeAvailableResourceRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return result, errors
	}
	return result, nil
}
