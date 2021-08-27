# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateProjects**](ProjectsApi.md#CreateOrUpdateProjects) | **Post** /projects | 
[**DeleteProject**](ProjectsApi.md#DeleteProject) | **Delete** /projects/{projectId} | 
[**GetAllProjects**](ProjectsApi.md#GetAllProjects) | **Get** /projects | 
[**GetProject**](ProjectsApi.md#GetProject) | **Get** /projects/{projectId} | 
[**GetProjectSchemaOptions**](ProjectsApi.md#GetProjectSchemaOptions) | **Get** /schemas | 

# **CreateOrUpdateProjects**
> ProjectsResponse CreateOrUpdateProjects(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectsRequest**](ProjectsRequest.md)|  | 

### Return type

[**ProjectsResponse**](ProjectsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> ProjectsResponse DeleteProject(ctx, projectId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
 **optional** | ***ProjectsApiDeleteProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiDeleteProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hard** | **optional.Bool**|  | [default to false]

### Return type

[**ProjectsResponse**](ProjectsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjects**
> ProjectsResponse GetAllProjects(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiGetAllProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetAllProjectsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **optional.String**|  | 

### Return type

[**ProjectsResponse**](ProjectsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> ProjectsResponse GetProject(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**ProjectsResponse**](ProjectsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectSchemaOptions**
> SchemasResponse GetProjectSchemaOptions(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SchemasResponse**](SchemasResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

