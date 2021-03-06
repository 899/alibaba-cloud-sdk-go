package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CreatePolicyWithSpecifiedPolicy(request *CreatePolicyWithSpecifiedPolicyRequest) (response *CreatePolicyWithSpecifiedPolicyResponse, err error) {
	response = CreateCreatePolicyWithSpecifiedPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreatePolicyWithSpecifiedPolicyWithChan(request *CreatePolicyWithSpecifiedPolicyRequest) (<-chan *CreatePolicyWithSpecifiedPolicyResponse, <-chan error) {
	responseChan := make(chan *CreatePolicyWithSpecifiedPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePolicyWithSpecifiedPolicy(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) CreatePolicyWithSpecifiedPolicyWithCallback(request *CreatePolicyWithSpecifiedPolicyRequest, callback func(response *CreatePolicyWithSpecifiedPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePolicyWithSpecifiedPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreatePolicyWithSpecifiedPolicy(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type CreatePolicyWithSpecifiedPolicyRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PolicyId             string           `position:"Query" name:"PolicyId"`
}

type CreatePolicyWithSpecifiedPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCreatePolicyWithSpecifiedPolicyRequest() (request *CreatePolicyWithSpecifiedPolicyRequest) {
	request = &CreatePolicyWithSpecifiedPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreatePolicyWithSpecifiedPolicy", "rds", "openAPI")
	return
}

func CreateCreatePolicyWithSpecifiedPolicyResponse() (response *CreatePolicyWithSpecifiedPolicyResponse) {
	response = &CreatePolicyWithSpecifiedPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
