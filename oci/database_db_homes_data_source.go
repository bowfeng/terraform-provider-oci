// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseDbHomesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbHomes,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_homes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseDbHomeDataSource()),
			},
		},
	}
}

func readDatabaseDbHomes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseDbHomesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbHomesResponse
}

func (s *DatabaseDbHomesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbHomesDataSourceCrud) Get() error {
	request := oci_database.ListDbHomesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DbHomeSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbHomes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbHomes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbHomesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbHome := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"db_system_id":   *r.DbSystemId,
		}

		if r.DbVersion != nil {
			dbHome["db_version"] = *r.DbVersion
		}

		if r.DisplayName != nil {
			dbHome["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			dbHome["id"] = *r.Id
			dbHome["db_home_id"] = *r.Id
		}

		if r.LastPatchHistoryEntryId != nil {
			dbHome["last_patch_history_entry_id"] = *r.LastPatchHistoryEntryId
		}

		dbHome["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dbHome["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, dbHome)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseDbHomesDataSource().Schema["db_homes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_homes", resources); err != nil {
		return err
	}

	return nil
}
