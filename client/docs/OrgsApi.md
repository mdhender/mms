# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateOrgs**](OrgsApi.md#CreateOrUpdateOrgs) | **Post** /orgs | 
[**DeleteOrg**](OrgsApi.md#DeleteOrg) | **Delete** /orgs/{orgId} | 
[**GetAllOrgs**](OrgsApi.md#GetAllOrgs) | **Get** /orgs | 
[**GetOrg**](OrgsApi.md#GetOrg) | **Get** /orgs/{orgId} | 

# **CreateOrUpdateOrgs**
> OrganizationsResponse CreateOrUpdateOrgs(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrganizationsRequest**](OrganizationsRequest.md)|  | 

### Return type

[**OrganizationsResponse**](OrganizationsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrg**
> OrganizationsResponse DeleteOrg(ctx, orgId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 

### Return type

[**OrganizationsResponse**](OrganizationsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllOrgs**
> OrganizationsResponse GetAllOrgs(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrganizationsResponse**](OrganizationsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrg**
> OrganizationsResponse GetOrg(ctx, orgId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 

### Return type

[**OrganizationsResponse**](OrganizationsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

