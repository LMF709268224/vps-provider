package services

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

const AccessKeyId = "LTAI5tLczyBS9h4APo9u9c1X"
const AccessKeySecret = "xHF9hl7qesCYqv8Wp1TLjhDU6EPPjQ"

func CreateInstanceWithOptions(RegionId, InstanceType, ImageId, SecurityGroupId string) (err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, err := CreateClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
	if err != nil {
		fmt.Errorf("%v", err)
		return err
	}

	createInstanceRequest := &ecs20140526.CreateInstanceRequest{
		RegionId:        tea.String(RegionId),
		InstanceType:    tea.String(InstanceType),
		DryRun:          tea.Bool(true),
		ImageId:         tea.String(ImageId),
		SecurityGroupId: tea.String(SecurityGroupId),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		result, err := client.CreateInstanceWithOptions(createInstanceRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		fmt.Println(result)
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_err_ := util.AssertAsString(error.Message)
		if _err_ != nil {
			fmt.Print(*_err_)
		}
	}
	return nil
}

func DescribePriceWithOptions(RegionId, InstanceType string) *float32 {
	var price *float32
	client, _err := CreateClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
	if _err != nil {
		return price
	}

	describePriceRequest := &ecs20140526.DescribePriceRequest{
		RegionId:     tea.String(RegionId),
		InstanceType: tea.String(InstanceType),
	}
	runtime := &util.RuntimeOptions{}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		result, _err := client.DescribePriceWithOptions(describePriceRequest, runtime)
		if _err != nil {
			return _err
		}
		fmt.Println(result.String())
		fmt.Println(result.Body.PriceInfo.Price.TradePrice)
		price = result.Body.PriceInfo.Price.TradePrice
		return nil
	}()
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
	}
	return price
}

func DescribeRecommendInstanceTypeWithOptions(RegionId string, Cores int32, Memory float32) *ecs20140526.DescribeRecommendInstanceTypeResponse {
	var result *ecs20140526.DescribeRecommendInstanceTypeResponse
	client, err := CreateClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
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
		// 复制代码运行请自行打印 API 的返回值
		result, err = client.DescribeRecommendInstanceTypeWithOptions(describeRecommendInstanceTypeRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		fmt.Println(result)
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
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
	client, err := CreateClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
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
		// 复制代码运行请自行打印 API 的返回值
		result, err = client.CreateSecurityGroupWithOptions(createSecurityGroupRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		fmt.Println(result)
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
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
	client, err := CreateClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
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
		// 复制代码运行请自行打印 API 的返回值
		result, err = client.DescribeImagesWithOptions(createSecurityGroupRequest, runtime)
		if err != nil {
			fmt.Errorf("%v", err)
			return err
		}
		fmt.Println(result)
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		return result
	}
	return result
}
