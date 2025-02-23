// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_health_checks "github.com/oracle/oci-go-sdk/healthchecks"
)

func HealthChecksHttpMonitorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createHealthChecksHttpMonitor,
		Read:     readHealthChecksHttpMonitor,
		Update:   updateHealthChecksHttpMonitor,
		Delete:   deleteHealthChecksHttpMonitor,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"interval_in_seconds": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"headers": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vantage_point_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"results_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createHealthChecksHttpMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksHttpMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return CreateResource(d, sync)
}

func readHealthChecksHttpMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksHttpMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return ReadResource(sync)
}

func updateHealthChecksHttpMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksHttpMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return UpdateResource(d, sync)
}

func deleteHealthChecksHttpMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksHttpMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type HealthChecksHttpMonitorResourceCrud struct {
	BaseCrud
	Client                 *oci_health_checks.HealthChecksClient
	Res                    *oci_health_checks.HttpMonitor
	DisableNotFoundRetries bool
}

func (s *HealthChecksHttpMonitorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *HealthChecksHttpMonitorResourceCrud) Create() error {
	request := oci_health_checks.CreateHttpMonitorRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if headers, ok := s.D.GetOkExists("headers"); ok {
		request.Headers = objectMapToStringMap(headers.(map[string]interface{}))
	}

	if intervalInSeconds, ok := s.D.GetOkExists("interval_in_seconds"); ok {
		tmp := intervalInSeconds.(int)
		request.IntervalInSeconds = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if method, ok := s.D.GetOkExists("method"); ok {
		request.Method = oci_health_checks.CreateHttpMonitorDetailsMethodEnum(method.(string))
	}

	if path, ok := s.D.GetOkExists("path"); ok {
		tmp := path.(string)
		request.Path = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_health_checks.CreateHttpMonitorDetailsProtocolEnum(protocol.(string))
	}

	request.Targets = []string{}
	if targets, ok := s.D.GetOkExists("targets"); ok {
		interfaces := targets.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Targets = tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	request.VantagePointNames = []string{}
	if vantagePointNames, ok := s.D.GetOkExists("vantage_point_names"); ok {
		interfaces := vantagePointNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.VantagePointNames = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.CreateHttpMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HttpMonitor
	return nil
}

func (s *HealthChecksHttpMonitorResourceCrud) Get() error {
	request := oci_health_checks.GetHttpMonitorRequest{}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.GetHttpMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HttpMonitor
	return nil
}

func (s *HealthChecksHttpMonitorResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_health_checks.UpdateHttpMonitorRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if headers, ok := s.D.GetOkExists("headers"); ok {
		request.Headers = objectMapToStringMap(headers.(map[string]interface{}))
	}

	if intervalInSeconds, ok := s.D.GetOkExists("interval_in_seconds"); ok {
		tmp := intervalInSeconds.(int)
		request.IntervalInSeconds = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if method, ok := s.D.GetOkExists("method"); ok {
		request.Method = oci_health_checks.UpdateHttpMonitorDetailsMethodEnum(method.(string))
	}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	if path, ok := s.D.GetOkExists("path"); ok {
		tmp := path.(string)
		request.Path = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_health_checks.UpdateHttpMonitorDetailsProtocolEnum(protocol.(string))
	}

	request.Targets = []string{}
	if targets, ok := s.D.GetOkExists("targets"); ok {
		interfaces := targets.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Targets = tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	request.VantagePointNames = []string{}
	if vantagePointNames, ok := s.D.GetOkExists("vantage_point_names"); ok {
		interfaces := vantagePointNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.VantagePointNames = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.UpdateHttpMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HttpMonitor
	return nil
}

func (s *HealthChecksHttpMonitorResourceCrud) Delete() error {
	request := oci_health_checks.DeleteHttpMonitorRequest{}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	_, err := s.Client.DeleteHttpMonitor(context.Background(), request)
	return err
}

func (s *HealthChecksHttpMonitorResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("headers", s.Res.Headers)

	if s.Res.IntervalInSeconds != nil {
		s.D.Set("interval_in_seconds", *s.Res.IntervalInSeconds)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("method", s.Res.Method)

	if s.Res.Path != nil {
		s.D.Set("path", *s.Res.Path)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	s.D.Set("protocol", s.Res.Protocol)

	if s.Res.ResultsUrl != nil {
		s.D.Set("results_url", *s.Res.ResultsUrl)
	}

	s.D.Set("targets", s.Res.Targets)

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	s.D.Set("vantage_point_names", s.Res.VantagePointNames)

	return nil
}

func (s *HealthChecksHttpMonitorResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_health_checks.ChangeHttpMonitorCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MonitorId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	_, err := s.Client.ChangeHttpMonitorCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
