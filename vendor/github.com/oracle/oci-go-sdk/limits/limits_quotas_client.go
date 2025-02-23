// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Quotas APIs
//
// APIs for managing Compartment Resource Quotas.
//

package limits

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//QuotasClient a client for Quotas
type QuotasClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewQuotasClientWithConfigurationProvider Creates a new default Quotas client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewQuotasClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client QuotasClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = QuotasClient{BaseClient: baseClient}
	client.BasePath = "20181025"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *QuotasClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("limits", "https://limits.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *QuotasClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *QuotasClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateQuota Creates a new quota with the details supplied.
func (client QuotasClient) CreateQuota(ctx context.Context, request CreateQuotaRequest) (response CreateQuotaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createQuota, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateQuotaResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateQuotaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateQuotaResponse")
	}
	return
}

// createQuota implements the OCIOperation interface (enables retrying operations)
func (client QuotasClient) createQuota(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/quotas/")
	if err != nil {
		return nil, err
	}

	var response CreateQuotaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteQuota Deletes the quota corresponding to the given OCID.
func (client QuotasClient) DeleteQuota(ctx context.Context, request DeleteQuotaRequest) (response DeleteQuotaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteQuota, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteQuotaResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteQuotaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteQuotaResponse")
	}
	return
}

// deleteQuota implements the OCIOperation interface (enables retrying operations)
func (client QuotasClient) deleteQuota(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/quotas/{quotaId}")
	if err != nil {
		return nil, err
	}

	var response DeleteQuotaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetQuota Gets the quota for the OCID specified.
func (client QuotasClient) GetQuota(ctx context.Context, request GetQuotaRequest) (response GetQuotaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQuota, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetQuotaResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQuotaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQuotaResponse")
	}
	return
}

// getQuota implements the OCIOperation interface (enables retrying operations)
func (client QuotasClient) getQuota(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/quotas/{quotaId}")
	if err != nil {
		return nil, err
	}

	var response GetQuotaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListQuotas Lists all quotas on resources from the given compartment
func (client QuotasClient) ListQuotas(ctx context.Context, request ListQuotasRequest) (response ListQuotasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQuotas, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListQuotasResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQuotasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQuotasResponse")
	}
	return
}

// listQuotas implements the OCIOperation interface (enables retrying operations)
func (client QuotasClient) listQuotas(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/quotas/")
	if err != nil {
		return nil, err
	}

	var response ListQuotasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateQuota Updates the quota corresponding to given OCID with the details supplied.
func (client QuotasClient) UpdateQuota(ctx context.Context, request UpdateQuotaRequest) (response UpdateQuotaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateQuota, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateQuotaResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateQuotaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateQuotaResponse")
	}
	return
}

// updateQuota implements the OCIOperation interface (enables retrying operations)
func (client QuotasClient) updateQuota(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/quotas/{quotaId}")
	if err != nil {
		return nil, err
	}

	var response UpdateQuotaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
