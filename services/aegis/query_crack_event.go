package aegis

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

func (client *Client) QueryCrackEvent(request *QueryCrackEventRequest) (response *QueryCrackEventResponse, err error) {
	response = CreateQueryCrackEventResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryCrackEventWithChan(request *QueryCrackEventRequest) (<-chan *QueryCrackEventResponse, <-chan error) {
	responseChan := make(chan *QueryCrackEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCrackEvent(request)
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

func (client *Client) QueryCrackEventWithCallback(request *QueryCrackEventRequest, callback func(response *QueryCrackEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCrackEventResponse
		var err error
		defer close(result)
		response, err = client.QueryCrackEvent(request)
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

type QueryCrackEventRequest struct {
	*requests.RpcRequest
	EndTime     string           `position:"Query" name:"EndTime"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	StartTime   string           `position:"Query" name:"StartTime"`
	Uuid        string           `position:"Query" name:"Uuid"`
	Status      requests.Integer `position:"Query" name:"Status"`
}

type QueryCrackEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      struct {
		PageInfo struct {
			CurrentPage int `json:"CurrentPage" xml:"CurrentPage"`
			PageSize    int `json:"PageSize" xml:"PageSize"`
			TotalCount  int `json:"TotalCount" xml:"TotalCount"`
			Count       int `json:"Count" xml:"Count"`
		} `json:"PageInfo" xml:"PageInfo"`
		List struct {
			Entity []struct {
				Uuid           string `json:"Uuid" xml:"Uuid"`
				AttackTime     string `json:"AttackTime" xml:"AttackTime"`
				AttackType     int    `json:"AttackType" xml:"AttackType"`
				AttackTypeName string `json:"AttackTypeName" xml:"AttackTypeName"`
				BuyVersion     string `json:"BuyVersion" xml:"BuyVersion"`
				CrackSourceIp  string `json:"CrackSourceIp" xml:"CrackSourceIp"`
				CrackTimes     int    `json:"CrackTimes" xml:"CrackTimes"`
				GroupId        int    `json:"GroupId" xml:"GroupId"`
				InstanceName   string `json:"InstanceName" xml:"InstanceName"`
				InstanceId     string `json:"InstanceId" xml:"InstanceId"`
				Ip             string `json:"Ip" xml:"Ip"`
				Region         string `json:"Region" xml:"Region"`
				Status         int    `json:"Status" xml:"Status"`
				StatusName     string `json:"StatusName" xml:"StatusName"`
				Location       string `json:"Location" xml:"Location"`
				InWhite        int    `json:"InWhite" xml:"InWhite"`
				UserName       string `json:"UserName" xml:"UserName"`
			} `json:"Entity" xml:"Entity"`
		} `json:"List" xml:"List"`
	} `json:"Data" xml:"Data"`
}

func CreateQueryCrackEventRequest() (request *QueryCrackEventRequest) {
	request = &QueryCrackEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "QueryCrackEvent", "vipaegis", "openAPI")
	return
}

func CreateQueryCrackEventResponse() (response *QueryCrackEventResponse) {
	response = &QueryCrackEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
