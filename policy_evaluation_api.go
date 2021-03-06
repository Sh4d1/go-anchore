/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.5
 * Contact: nurmi@anchore.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type PolicyEvaluationApiService service


/* PolicyEvaluationApiService Check policy evaluation status for image
 Get the policy evaluation for the given image
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param imageDigest 
 @param tag 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "policyId" (string) 
     @param "detail" (bool) 
     @param "history" (bool) 
 @return PolicyEvaluation*/
func (a *PolicyEvaluationApiService) GetImagePolicyCheck(ctx context.Context, imageDigest string, tag string, localVarOptionals map[string]interface{}) (PolicyEvaluation,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolicyEvaluation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/{imageDigest}/check"
	localVarPath = strings.Replace(localVarPath, "{"+"imageDigest"+"}", fmt.Sprintf("%v", imageDigest), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["policyId"], "string", "policyId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["detail"], "bool", "detail"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["history"], "bool", "history"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["policyId"].(string); localVarOk {
		localVarQueryParams.Add("policyId", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("tag", parameterToString(tag, ""))
	if localVarTempParam, localVarOk := localVarOptionals["detail"].(bool); localVarOk {
		localVarQueryParams.Add("detail", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["history"].(bool); localVarOk {
		localVarQueryParams.Add("history", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PolicyEvaluationApiService Check policy evaluation status for image
 Get the policy evaluation for the given image
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param imageId 
 @param tag 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "policyId" (string) 
     @param "detail" (bool) 
     @param "history" (bool) 
 @return PolicyEvaluation*/
func (a *PolicyEvaluationApiService) GetImagePolicyCheckByImageId(ctx context.Context, imageId string, tag string, localVarOptionals map[string]interface{}) (PolicyEvaluation,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolicyEvaluation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/by_id/{imageId}/check"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["policyId"], "string", "policyId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["detail"], "bool", "detail"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["history"], "bool", "history"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["policyId"].(string); localVarOk {
		localVarQueryParams.Add("policyId", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("tag", parameterToString(tag, ""))
	if localVarTempParam, localVarOk := localVarOptionals["detail"].(bool); localVarOk {
		localVarQueryParams.Add("detail", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["history"].(bool); localVarOk {
		localVarQueryParams.Add("history", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

