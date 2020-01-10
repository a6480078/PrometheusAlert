package vod

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

// ListAIJob invokes the vod.ListAIJob API synchronously
// api document: https://help.aliyun.com/api/vod/listaijob.html
func (client *Client) ListAIJob(request *ListAIJobRequest) (response *ListAIJobResponse, err error) {
	response = CreateListAIJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListAIJobWithChan invokes the vod.ListAIJob API asynchronously
// api document: https://help.aliyun.com/api/vod/listaijob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAIJobWithChan(request *ListAIJobRequest) (<-chan *ListAIJobResponse, <-chan error) {
	responseChan := make(chan *ListAIJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAIJob(request)
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

// ListAIJobWithCallback invokes the vod.ListAIJob API asynchronously
// api document: https://help.aliyun.com/api/vod/listaijob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAIJobWithCallback(request *ListAIJobRequest, callback func(response *ListAIJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAIJobResponse
		var err error
		defer close(result)
		response, err = client.ListAIJob(request)
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

// ListAIJobRequest is the request struct for api ListAIJob
type ListAIJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	JobIds               string `position:"Query" name:"JobIds"`
}

// ListAIJobResponse is the response struct for api ListAIJob
type ListAIJobResponse struct {
	*responses.BaseResponse
	RequestId        string               `json:"RequestId" xml:"RequestId"`
	NonExistAIJobIds NonExistAIJobIds     `json:"NonExistAIJobIds" xml:"NonExistAIJobIds"`
	AIJobList        AIJobListInListAIJob `json:"AIJobList" xml:"AIJobList"`
}

// CreateListAIJobRequest creates a request to invoke ListAIJob API
func CreateListAIJobRequest() (request *ListAIJobRequest) {
	request = &ListAIJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListAIJob", "vod", "openAPI")
	return
}

// CreateListAIJobResponse creates a response to parse from ListAIJob response
func CreateListAIJobResponse() (response *ListAIJobResponse) {
	response = &ListAIJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}