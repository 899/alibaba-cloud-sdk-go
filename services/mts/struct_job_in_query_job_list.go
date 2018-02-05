package mts

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

type JobInQueryJobList struct {
	JobId            string           `json:"JobId" xml:"JobId"`
	State            string           `json:"State" xml:"State"`
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	Percent          int              `json:"Percent" xml:"Percent"`
	PipelineId       string           `json:"PipelineId" xml:"PipelineId"`
	CreationTime     string           `json:"CreationTime" xml:"CreationTime"`
	FinishTime       string           `json:"FinishTime" xml:"FinishTime"`
	Input            Input            `json:"Input" xml:"Input"`
	Output           Output           `json:"Output" xml:"Output"`
	MNSMessageResult MNSMessageResult `json:"MNSMessageResult" xml:"MNSMessageResult"`
}