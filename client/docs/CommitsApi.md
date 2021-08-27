# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommit**](CommitsApi.md#GetCommit) | **Get** /projects/{projectId}/commits/{commitId} | 
[**GetCommits**](CommitsApi.md#GetCommits) | **Put** /projects/{projectId}/commits | 
[**GetElementCommits**](CommitsApi.md#GetElementCommits) | **Get** /projects/{projectId}/refs/{refId}/elements/{elementId}/commits | 
[**GetRefCommits**](CommitsApi.md#GetRefCommits) | **Get** /projects/{projectId}/refs/{refId}/commits | 

# **GetCommit**
> CommitsResponse GetCommit(ctx, projectId, commitId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **commitId** | **string**|  | 

### Return type

[**CommitsResponse**](CommitsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommits**
> CommitsResponse GetCommits(ctx, body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommitsRequest**](CommitsRequest.md)|  | 
  **projectId** | **string**|  | 

### Return type

[**CommitsResponse**](CommitsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetElementCommits**
> CommitsResponse GetElementCommits(ctx, projectId, refId, elementId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **elementId** | **string**|  | 

### Return type

[**CommitsResponse**](CommitsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRefCommits**
> CommitsResponse GetRefCommits(ctx, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***CommitsApiGetRefCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetRefCommitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.String**|  | 
 **maxTimestamp** | **optional.String**|  | 

### Return type

[**CommitsResponse**](CommitsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

