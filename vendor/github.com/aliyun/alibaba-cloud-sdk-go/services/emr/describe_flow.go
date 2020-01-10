package emr

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

// DescribeFlow invokes the emr.DescribeFlow API synchronously
// api document: https://help.aliyun.com/api/emr/describeflow.html
func (client *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
	response = CreateDescribeFlowResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowWithChan invokes the emr.DescribeFlow API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowWithChan(request *DescribeFlowRequest) (<-chan *DescribeFlowResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlow(request)
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

// DescribeFlowWithCallback invokes the emr.DescribeFlow API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowWithCallback(request *DescribeFlowRequest, callback func(response *DescribeFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlow(request)
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

// DescribeFlowRequest is the request struct for api DescribeFlow
type DescribeFlowRequest struct {
	*requests.RpcRequest
	Id        string `position:"Query" name:"Id"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DescribeFlowResponse is the response struct for api DescribeFlow
type DescribeFlowResponse struct {
	*responses.BaseResponse
	RequestId               string         `json:"RequestId" xml:"RequestId"`
	Id                      string         `json:"Id" xml:"Id"`
	GmtCreate               int64          `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified             int64          `json:"GmtModified" xml:"GmtModified"`
	Name                    string         `json:"Name" xml:"Name"`
	Description             string         `json:"Description" xml:"Description"`
	Type                    string         `json:"Type" xml:"Type"`
	Status                  string         `json:"Status" xml:"Status"`
	Periodic                bool           `json:"Periodic" xml:"Periodic"`
	StartSchedule           int64          `json:"StartSchedule" xml:"StartSchedule"`
	EndSchedule             int64          `json:"EndSchedule" xml:"EndSchedule"`
	CronExpr                string         `json:"CronExpr" xml:"CronExpr"`
	CreateCluster           bool           `json:"CreateCluster" xml:"CreateCluster"`
	ClusterId               string         `json:"ClusterId" xml:"ClusterId"`
	HostName                string         `json:"HostName" xml:"HostName"`
	Graph                   string         `json:"Graph" xml:"Graph"`
	CategoryId              string         `json:"CategoryId" xml:"CategoryId"`
	AlertConf               string         `json:"AlertConf" xml:"AlertConf"`
	AlertUserGroupBizId     string         `json:"AlertUserGroupBizId" xml:"AlertUserGroupBizId"`
	AlertDingDingGroupBizId string         `json:"AlertDingDingGroupBizId" xml:"AlertDingDingGroupBizId"`
	Application             string         `json:"Application" xml:"Application"`
	EditLockDetail          string         `json:"EditLockDetail" xml:"EditLockDetail"`
	ParentFlowList          ParentFlowList `json:"ParentFlowList" xml:"ParentFlowList"`
}

// CreateDescribeFlowRequest creates a request to invoke DescribeFlow API
func CreateDescribeFlowRequest() (request *DescribeFlowRequest) {
	request = &DescribeFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlow", "emr", "openAPI")
	return
}

// CreateDescribeFlowResponse creates a response to parse from DescribeFlow response
func CreateDescribeFlowResponse() (response *DescribeFlowResponse) {
	response = &DescribeFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}