# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateArtifacts**](ArtifactsApi.md#CreateOrUpdateArtifacts) | **Post** /projects/{projectId}/refs/{refId}/elements/{elementId} | 
[**DeleteArtifactByExtension**](ArtifactsApi.md#DeleteArtifactByExtension) | **Delete** /projects/{projectId}/refs/{refId}/elements/{elementId}/{extension} | 
[**GetArtifactByExtension**](ArtifactsApi.md#GetArtifactByExtension) | **Get** /projects/{projectId}/refs/{refId}/elements/{elementId}/{extension} | 

# **CreateOrUpdateArtifacts**
> ElementsResponse CreateOrUpdateArtifacts(ctx, projectId, refId, elementId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **elementId** | **string**|  | 
 **optional** | ***ArtifactsApiCreateOrUpdateArtifactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiCreateOrUpdateArtifactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**ElementsResponse**](ElementsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArtifactByExtension**
> ElementsResponse DeleteArtifactByExtension(ctx, projectId, refId, elementId, extension)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **elementId** | **string**|  | 
  **extension** | **string**|  | 

### Return type

[**ElementsResponse**](ElementsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactByExtension**
> string GetArtifactByExtension(ctx, projectId, refId, elementId, extension, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 
  **elementId** | **string**|  | 
  **extension** | **string**|  | 
 **optional** | ***ArtifactsApiGetArtifactByExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiGetArtifactByExtensionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **commitId** | **optional.String**|  | 

### Return type

**string**

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

