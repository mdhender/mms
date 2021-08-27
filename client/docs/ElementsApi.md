# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateElements**](ElementsApi.md#CreateOrUpdateElements) | **Post** /projects/{projectId}/refs/{refId}/elements | 
[**DeleteElement**](ElementsApi.md#DeleteElement) | **Delete** /projects/{projectId}/refs/{refId}/elements/{elementId} | 
[**DeleteElements**](ElementsApi.md#DeleteElements) | **Delete** /projects/{projectId}/refs/{refId}/elements | 
[**GetAllElements**](ElementsApi.md#GetAllElements) | **Get** /projects/{projectId}/refs/{refId}/elements | 
[**GetArtifactForElement1**](ElementsApi.md#GetArtifactForElement1) | **Get** /projects/{projectId}/refs/{refId}/elements/{elementId} | 
[**GetElements**](ElementsApi.md#GetElements) | **Put** /projects/{projectId}/refs/{refId}/elements | 

# **CreateOrUpdateElements**
> ElementsCommitResponse CreateOrUpdateElements(ctx, body, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ElementsRequest**](ElementsRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***ElementsApiCreateOrUpdateElementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementsApiCreateOrUpdateElementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **overwrite** | **optional.**|  | 

### Return type

[**ElementsCommitResponse**](ElementsCommitResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteElement**
> ElementsCommitResponse DeleteElement(ctx, projectId, refId, elementId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **elementId** | **string**|  | 

### Return type

[**ElementsCommitResponse**](ElementsCommitResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteElements**
> ElementsResponse DeleteElements(ctx, projectId, refId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 

### Return type

[**ElementsResponse**](ElementsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllElements**
> ElementsResponse GetAllElements(ctx, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***ElementsApiGetAllElementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementsApiGetAllElementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commitId** | **optional.String**|  | 

### Return type

[**ElementsResponse**](ElementsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactForElement1**
> InlineResponse200 GetArtifactForElement1(ctx, projectId, refId, elementId, accept, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **elementId** | **string**|  | 
  **accept** | **string**|  | 
 **optional** | ***ElementsApiGetArtifactForElement1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementsApiGetArtifactForElement1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **commitId** | **optional.String**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetElements**
> ElementsResponse GetElements(ctx, body, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ElementsRequest**](ElementsRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***ElementsApiGetElementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementsApiGetElementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **commitId** | **optional.**|  | 

### Return type

[**ElementsResponse**](ElementsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

