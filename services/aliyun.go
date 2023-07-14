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
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}

	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

func CreateInstanceWithOptions(RegionId, InstanceType, ImageId, SecurityGroupId, PeriodUnit string, Period int32) (*ecs20140526.CreateInstanceResponse, error) {
	var result *ecs20140526.CreateInstanceResponse

	client, err := CreateClient(tea.String(config.Cfg.AliyunAccessKeyID), tea.String(config.Cfg.AliyunAccessKeySecret))
	if err != nil {
		return result, err
	}

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
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err = client.CreateInstanceWithOptions(createInstanceRequest, runtime)
		if err != nil {
			fmt.Errorf("CreateInstanceWithOptions %v", err)
			return err
		}
		// fmt.Println(result)
		return nil
	}()

	if tryErr != nil {
		error := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}

		_err := util.AssertAsString(error.Message)
		if _err != nil {
			fmt.Print(*_err)
		}
	}
	return result, nil
}

func DescribePriceWithOptions(RegionId, InstanceType, PriceUnit string, Period int32) *ecs20140526.DescribePriceResponseBodyPriceInfoPrice {
	var price *ecs20140526.DescribePriceResponseBodyPriceInfoPrice
	client, _err := CreateClient(tea.String(config.Cfg.AliyunAccessKeyID), tea.String(config.Cfg.AliyunAccessKeySecret))
	if _err != nil {
		return price
	}

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
		result, _err := client.DescribePriceWithOptions(describePriceRequest, runtime)
		if _err != nil {
			return _err
		}
		price = result.Body.PriceInfo.Price
		return nil
	}()
	if tryErr != nil {
		error := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
	}
	return price
}

func DescribeRegionsWithOptions() *ecs20140526.DescribeRegionsResponse {
	var result *ecs20140526.DescribeRegionsResponse
	client, _err := CreateClient(tea.String(config.Cfg.AliyunAccessKeyID), tea.String(config.Cfg.AliyunAccessKeySecret))
	if _err != nil {
		return nil
	}

	describeRegionsRequest := &ecs20140526.DescribeRegionsRequest{}
	runtime := &util.RuntimeOptions{}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _err = client.DescribeRegionsWithOptions(describeRegionsRequest, runtime)
		if _err != nil {
			return _err
		}
		return nil
	}()
	if tryErr != nil {
		error := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
	}
	return result
}

func DescribeRecommendInstanceTypeWithOptions(RegionId string, Cores int32, Memory float32) *ecs20140526.DescribeRecommendInstanceTypeResponse {
	var result *ecs20140526.DescribeRecommendInstanceTypeResponse
	client, err := CreateClient(tea.String(config.Cfg.AliyunAccessKeyID), tea.String(config.Cfg.AliyunAccessKeySecret))
	if err != nil {
		fmt.Errorf("%v", err)
		return result
	}

	describeRecommendInstanceTypeRequest := &ecs20140526.DescribeRecommendInstanceTypeRequest{
		NetworkType: tea.String("vpc"),
		RegionId:    tea.String(RegionId),
		Cores:       tea.Int32(Cores),
		Memory:      tea.Float32(Memory),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err = client.DescribeRecommendInstanceTypeWithOptions(describeRecommendInstanceTypeRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		error := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		return result
	}
	return result
}

func CreateSecurityGroup(RegionId string) *ecs20140526.CreateSecurityGroupResponse {
	var result *ecs20140526.CreateSecurityGroupResponse
	client, err := CreateClient(tea.String(config.Cfg.AliyunAccessKeyID), tea.String(config.Cfg.AliyunAccessKeySecret))
	if err != nil {
		fmt.Errorf("%v", err)
		return result
	}

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
		result, err = client.CreateSecurityGroupWithOptions(createSecurityGroupRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		error := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		return result
	}
	return result
}

func DescribeImagesWithOptions(RegionId string) *ecs20140526.DescribeImagesResponse {
	var result *ecs20140526.DescribeImagesResponse
	client, err := CreateClient(tea.String(config.Cfg.AliyunAccessKeyID), tea.String(config.Cfg.AliyunAccessKeySecret))
	if err != nil {
		fmt.Errorf("%v", err)
		return result
	}

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
		result, err = client.DescribeImagesWithOptions(createSecurityGroupRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		error := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		return result
	}
	return result
}
