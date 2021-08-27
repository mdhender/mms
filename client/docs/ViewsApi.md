# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateViews**](ViewsApi.md#CreateOrUpdateViews) | **Post** /projects/{projectId}/refs/{refId}/views | 
[**GetDocuments**](ViewsApi.md#GetDocuments) | **Get** /projects/{projectId}/refs/{refId}/documents | 
[**GetGroups**](ViewsApi.md#GetGroups) | **Get** /projects/{projectId}/refs/{refId}/groups | 
[**GetMounts**](ViewsApi.md#GetMounts) | **Get** /projects/{projectId}/refs/{refId}/mounts | 
[**GetView**](ViewsApi.md#GetView) | **Get** /projects/{projectId}/refs/{refId}/views/{viewId} | 
[**GetViews**](ViewsApi.md#GetViews) | **Put** /projects/{projectId}/refs/{refId}/views | 

# **CreateOrUpdateViews**
> ElementsResponse CreateOrUpdateViews(ctx, body, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ElementsRequest**](ElementsRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***ViewsApiCreateOrUpdateViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiCreateOrUpdateViewsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **overwrite** | **optional.**|  | 

### Return type

[**ElementsResponse**](ElementsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocuments**
> DocumentsResponse GetDocuments(ctx, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***ViewsApiGetDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiGetDocumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commitId** | **optional.String**|  | 

### Return type

[**DocumentsResponse**](DocumentsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups**
> GroupsResponse GetGroups(ctx, projectId, refId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 

### Return type

[**GroupsResponse**](GroupsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMounts**
> MountsResponse GetMounts(ctx, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***ViewsApiGetMountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiGetMountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commitId** | **optional.String**|  | 

### Return type

[**MountsResponse**](MountsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetView**
> ElementsResponse GetView(ctx, projectId, refId, viewId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **viewId** | **string**|  | 
 **optional** | ***ViewsApiGetViewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiGetViewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **commitId** | **optional.String**|  | 

### Return type

[**ElementsResponse**](ElementsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetViews**
> ElementsResponse GetViews(ctx, body, projectId, refId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ElementsRequest**](ElementsRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
 **optional** | ***ViewsApiGetViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiGetViewsOpts struct
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

