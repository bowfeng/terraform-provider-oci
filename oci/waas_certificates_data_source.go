// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_waas "github.com/oracle/oci-go-sdk/waas"
)

func WaasCertificatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWaasCertificates,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(WaasCertificateResource()),
			},
		},
	}
}

func readWaasCertificates(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCertificatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient

	return ReadResource(sync)
}

type WaasCertificatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.ListCertificatesResponse
}

func (s *WaasCertificatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasCertificatesDataSourceCrud) Get() error {
	request := oci_waas.ListCertificatesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.DisplayName = []string{}
	if displayNames, ok := s.D.GetOkExists("display_names"); ok {
		interfaces := displayNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.DisplayName = tmp
	}

	request.Id = []string{}
	if ids, ok := s.D.GetOkExists("ids"); ok {
		interfaces := ids.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Id = tmp
	}

	request.LifecycleState = []oci_waas.ListCertificatesLifecycleStateEnum{}
	if states, ok := s.D.GetOkExists("states"); ok {
		interfaces := states.([]interface{})
		tmp := make([]oci_waas.ListCertificatesLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waas.ListCertificatesLifecycleStateEnum(interfaces[i].(string))
			}
		}
		request.LifecycleState = tmp
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "waas")

	response, err := s.Client.ListCertificates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCertificates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WaasCertificatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		certificate := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			certificate["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			certificate["display_name"] = *r.DisplayName
		}

		certificate["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			certificate["id"] = *r.Id
		}

		certificate["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			certificate["time_created"] = r.TimeCreated.String()
		}

		if r.TimeNotValidAfter != nil {
			certificate["time_not_valid_after"] = r.TimeNotValidAfter.String()
		}

		resources = append(resources, certificate)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, WaasCertificatesDataSource().Schema["certificates"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("certificates", resources); err != nil {
		return err
	}

	return nil
}
