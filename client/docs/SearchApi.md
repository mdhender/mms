# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBasicSearch**](SearchApi.md#GetBasicSearch) | **Get** /projects/{projectId}/refs/{refId}/search | 
[**PostBasicSearch**](SearchApi.md#PostBasicSearch) | **Post** /projects/{projectId}/refs/{refId}/search | 

# **GetBasicSearch**
> ElementsSearchResponse GetBasicSearch(ctx, projectId, refId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 

### Return type

[**ElementsSearchResponse**](ElementsSearchResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBasicSearch**
> ElementsSearchResponse PostBasicSearch(ctx, body, projectId, refId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BasicSearchRequest**](BasicSearchRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 

### Return type

[**ElementsSearchResponse**](ElementsSearchResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

