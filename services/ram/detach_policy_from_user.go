package ram

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

func (client *Client) DetachPolicyFromUser(request *DetachPolicyFromUserRequest) (response *DetachPolicyFromUserResponse, err error) {
	response = CreateDetachPolicyFromUserResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DetachPolicyFromUserWithChan(request *DetachPolicyFromUserRequest) (<-chan *DetachPolicyFromUserResponse, <-chan error) {
	responseChan := make(chan *DetachPolicyFromUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachPolicyFromUser(request)
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

func (client *Client) DetachPolicyFromUserWithCallback(request *DetachPolicyFromUserRequest, callback func(response *DetachPolicyFromUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachPolicyFromUserResponse
		var err error
		defer close(result)
		response, err = client.DetachPolicyFromUser(request)
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

type DetachPolicyFromUserRequest struct {
	*requests.RpcRequest
	UserName   string `position:"Query" name:"UserName"`
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

type DetachPolicyFromUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDetachPolicyFromUserRequest() (request *DetachPolicyFromUserRequest) {
	request = &DetachPolicyFromUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromUser", "", "")
	return
}

func CreateDetachPolicyFromUserResponse() (response *DetachPolicyFromUserResponse) {
	response = &DetachPolicyFromUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
