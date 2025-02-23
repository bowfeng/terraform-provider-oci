// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseDbNodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbNodes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseDbNodeDataSource()),
			},
		},
	}
}

func readDatabaseDbNodes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseDbNodesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbNodesResponse
}

func (s *DatabaseDbNodesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodesDataSourceCrud) Get() error {
	request := oci_database.ListDbNodesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DbNodeSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbNodes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbNodes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbNodesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbNode := map[string]interface{}{
			"db_system_id": *r.DbSystemId,
		}

		if r.BackupVnicId != nil {
			dbNode["backup_vnic_id"] = *r.BackupVnicId
		}

		if r.FaultDomain != nil {
			dbNode["fault_domain"] = *r.FaultDomain
		}

		if r.Hostname != nil {
			dbNode["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			dbNode["id"] = *r.Id
			dbNode["db_node_id"] = *r.Id // maintain legacy vanity id
		}

		if r.SoftwareStorageSizeInGB != nil {
			dbNode["software_storage_size_in_gb"] = *r.SoftwareStorageSizeInGB
		}

		dbNode["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dbNode["time_created"] = r.TimeCreated.String()
		}

		if r.VnicId != nil {
			dbNode["vnic_id"] = *r.VnicId
		}

		resources = append(resources, dbNode)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseDbNodesDataSource().Schema["db_nodes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_nodes", resources); err != nil {
		return err
	}

	return nil
}
