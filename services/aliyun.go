package services

import (
	"vps-provider/config"
	"vps-provider/types"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

// ClientInit /**
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

func CreateInstance(regionId, instanceType, imageId, securityGroupId, periodUnit string, period int32) (*types.CreateInstanceResponse, *tea.SDKError) {
	var out *types.CreateInstanceResponse

	createInstanceRequest := &ecs20140526.CreateInstanceRequest{
		RegionId:           tea.String(regionId),
		InstanceType:       tea.String(instanceType),
		DryRun:             tea.Bool(config.Cfg.DryRun),
		ImageId:            tea.String(imageId),
		SecurityGroupId:    tea.String(securityGroupId),
		InstanceChargeType: tea.String("PrePaid"),
		PeriodUnit:         tea.String(periodUnit),
		Period:             tea.Int32(period),
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err := AliClient.CreateInstanceWithOptions(createInstanceRequest, runtime)
		if err != nil {
			return err
		}

		out = &types.CreateInstanceResponse{
			InstanceId: *result.Body.InstanceId,
			OrderId:    *result.Body.OrderId,
			RequestId:  *result.Body.RequestId,
			TradePrice: *result.Body.TradePrice,
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
		return out, errors
	}
	return out, nil
}

func DescribePriceWithOptions(regionId, instanceType, priceUnit string, period int32) (*types.DescribePriceResponse, *tea.SDKError) {
	var out *types.DescribePriceResponse
	describePriceRequest := &ecs20140526.DescribePriceRequest{
		RegionId:     tea.String(regionId),
		InstanceType: tea.String(instanceType),
		PriceUnit:    tea.String(priceUnit),
		Period:       tea.Int32(period),
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
			return err
		}
		price := result.Body.PriceInfo.Price
		out = &types.DescribePriceResponse{
			Currency:      *price.Currency,
			OriginalPrice: *price.OriginalPrice,
			TradePrice:    *price.TradePrice,
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
		return out, errors
	}
	return out, nil
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

func DescribeRecommendInstanceTypeWithOptions(regionId string, cores int32, memory float32) (*ecs20140526.DescribeRecommendInstanceTypeResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeRecommendInstanceTypeResponse
	var err error
	describeRecommendInstanceTypeRequest := &ecs20140526.DescribeRecommendInstanceTypeRequest{
		// NetworkType:        tea.String("vpc"),
		RegionId:           tea.String(regionId),
		Cores:              tea.Int32(cores),
		Memory:             tea.Float32(memory),
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

func CreateSecurityGroup(regionId string) (*ecs20140526.CreateSecurityGroupResponse, *tea.SDKError) {
	var result *ecs20140526.CreateSecurityGroupResponse
	var err error

	createSecurityGroupRequest := &ecs20140526.CreateSecurityGroupRequest{
		RegionId: tea.String(regionId),
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

func DescribeImagesWithOptions(regionId, instanceType string) (*ecs20140526.DescribeImagesResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeImagesResponse
	var err error

	createSecurityGroupRequest := &ecs20140526.DescribeImagesRequest{
		RegionId:     tea.String(regionId),
		InstanceType: tea.String(instanceType),
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

func DescribeAvailableResourceWithOptions(regionId string, cores int32, memory float32) (*ecs20140526.DescribeAvailableResourceResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeAvailableResourceResponse
	var err error
	describeAvailableResourceRequest := &ecs20140526.DescribeAvailableResourceRequest{
		RegionId:            tea.String(regionId),
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
