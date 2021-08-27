# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateNotebooks**](NotebooksApi.md#CreateOrUpdateNotebooks) | **Post** /projects/{projectId}/refs/{refId}/notebooks | 
[**GetAllNotebooks**](NotebooksApi.md#GetAllNotebooks) | **Get** /projects/{projectId}/refs/{refId}/notebooks | 
[**GetNotebook**](NotebooksApi.md#GetNotebook) | **Get** /projects/{projectId}/refs/{refId}/notebooks/{notebookId} | 
[**GetNotebooks**](NotebooksApi.md#GetNotebooks) | **Put** /projects/{projectId}/refs/{refId}/notebooks | 

# **CreateOrUpdateNotebooks**
> NotebooksResponse CreateOrUpdateNotebooks(ctx, body, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotebooksRequest**](NotebooksRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***NotebooksApiCreateOrUpdateNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotebooksApiCreateOrUpdateNotebooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **overwrite** | **optional.**|  | 

### Return type

[**NotebooksResponse**](NotebooksResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllNotebooks**
> NotebooksResponse GetAllNotebooks(ctx, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***NotebooksApiGetAllNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotebooksApiGetAllNotebooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commitId** | **optional.String**|  | 

### Return type

[**NotebooksResponse**](NotebooksResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotebook**
> NotebooksResponse GetNotebook(ctx, projectId, refId, notebookId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **notebookId** | **string**|  | 
 **optional** | ***NotebooksApiGetNotebookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotebooksApiGetNotebookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **commitId** | **optional.String**|  | 

### Return type

[**NotebooksResponse**](NotebooksResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotebooks**
> NotebooksResponse GetNotebooks(ctx, body, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotebooksRequest**](NotebooksRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***NotebooksApiGetNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotebooksApiGetNotebooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **commitId** | **optional.**|  | 

### Return type

[**NotebooksResponse**](NotebooksResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

