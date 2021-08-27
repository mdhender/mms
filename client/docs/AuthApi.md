# {{classname}}

All URIs are relative to *https://mms.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAuthenticationToken**](AuthApi.md#CheckAuthenticationToken) | **Get** /checkAuth | 
[**CreateAuthenticationToken**](AuthApi.md#CreateAuthenticationToken) | **Post** /authentication | 
[**CreateUser**](AuthApi.md#CreateUser) | **Post** /user | 
[**GetAuthenticationToken**](AuthApi.md#GetAuthenticationToken) | **Get** /authentication | 
[**GetBranchPermissions**](AuthApi.md#GetBranchPermissions) | **Get** /projects/{projectId}/refs/{refId}/permissions | 
[**GetOrgPermissions**](AuthApi.md#GetOrgPermissions) | **Get** /orgs/{orgId}/permissions | 
[**GetProjectPermissions**](AuthApi.md#GetProjectPermissions) | **Get** /projects/{projectId}/permissions | 
[**GetUsers**](AuthApi.md#GetUsers) | **Get** /users | 
[**LookupPermissions**](AuthApi.md#LookupPermissions) | **Put** /permissions | 
[**UpdateBranchPermissions**](AuthApi.md#UpdateBranchPermissions) | **Post** /projects/{projectId}/refs/{refId}/permissions | 
[**UpdateOrgPermissions**](AuthApi.md#UpdateOrgPermissions) | **Post** /orgs/{orgId}/permissions | 
[**UpdatePassword**](AuthApi.md#UpdatePassword) | **Post** /password | 
[**UpdateProjectPermissions**](AuthApi.md#UpdateProjectPermissions) | **Post** /projects/{projectId}/permissions | 

# **CheckAuthenticationToken**
> JwtTokenValidationResponse CheckAuthenticationToken(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JwtTokenValidationResponse**](JwtTokenValidationResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthenticationToken**
> JwtAuthenticationResponse CreateAuthenticationToken(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JwtAuthenticationRequest**](JwtAuthenticationRequest.md)|  | 

### Return type

[**JwtAuthenticationResponse**](JwtAuthenticationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> UserCreateRequest CreateUser(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserCreateRequest**](UserCreateRequest.md)|  | 

### Return type

[**UserCreateRequest**](UserCreateRequest.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthenticationToken**
> JwtAuthenticationResponse GetAuthenticationToken(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JwtAuthenticationResponse**](JwtAuthenticationResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranchPermissions**
> PermissionsResponse GetBranchPermissions(ctx, projectId, refId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **refId** | **string**|  | 

### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgPermissions**
> PermissionsResponse GetOrgPermissions(ctx, orgId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 

### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectPermissions**
> PermissionsResponse GetProjectPermissions(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> UsersResponse GetUsers(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupPermissions**
> PermissionLookupResponse LookupPermissions(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PermissionLookupRequest**](PermissionLookupRequest.md)|  | 

### Return type

[**PermissionLookupResponse**](PermissionLookupResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBranchPermissions**
> PermissionUpdatesResponse UpdateBranchPermissions(ctx, body, projectId, refId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PermissionsRequest**](PermissionsRequest.md)|  | 
  **projectId** | **string**|  | 
  **refId** | **string**|  | 

### Return type

[**PermissionUpdatesResponse**](PermissionUpdatesResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgPermissions**
> PermissionUpdatesResponse UpdateOrgPermissions(ctx, body, orgId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PermissionsRequest**](PermissionsRequest.md)|  | 
  **orgId** | **string**|  | 

### Return type

[**PermissionUpdatesResponse**](PermissionUpdatesResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePassword**
> interface{} UpdatePassword(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserCreateRequest**](UserCreateRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectPermissions**
> PermissionUpdatesResponse UpdateProjectPermissions(ctx, body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PermissionsRequest**](PermissionsRequest.md)|  | 
  **projectId** | **string**|  | 

### Return type

[**PermissionUpdatesResponse**](PermissionUpdatesResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

