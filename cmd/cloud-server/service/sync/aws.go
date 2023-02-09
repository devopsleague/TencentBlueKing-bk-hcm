/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package sync

import (
	"net/http"

	"hcm/pkg/api/core"
	dataproto "hcm/pkg/api/data-service/cloud"
	protoregion "hcm/pkg/api/data-service/cloud/region"
	proto "hcm/pkg/api/hc-service"
	protodisk "hcm/pkg/api/hc-service/disk"
	"hcm/pkg/client"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/dal/dao/tools"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
	"hcm/pkg/runtime/filter"
)

// SyncAwsAll sync aws all resource
func SyncAwsAll(c *client.ClientSet, kit *kit.Kit, header http.Header, accountID string) error {

	regions, err := c.DataService().Aws.Region.ListRegion(
		kit.Ctx,
		header,
		&protoregion.AwsRegionListReq{
			Filter: tools.EqualExpression("vendor", enumor.Aws),
			Page:   &core.BasePage{Start: 0, Limit: core.DefaultMaxPageLimit},
		},
	)
	if err != nil {
		logs.Errorf("sync list aws region failed, err: %v, rid: %s", err, kit.Rid)
		return err
	}

	for _, region := range regions.Details {

		// sg
		err = c.HCService().Aws.SecurityGroup.SyncSecurityGroup(
			kit.Ctx,
			header,
			&proto.SecurityGroupSyncReq{
				AccountID: accountID,
				Region:    region.RegionID,
			},
		)
		if err != nil {
			logs.Errorf("sync do aws sync sg failed, err: %v, regionID: %s, rid: %s",
				err, region.RegionID, kit.Rid)
		}

		// sg rule
		err = SyncAwsSGRule(c, kit, header, region.RegionID, accountID)
		if err != nil {
			logs.Errorf("sync do aws sync sg rule failed, err: %v, regionID: %s,  rid: %s",
				err, region.RegionID, kit.Rid)
		}

		// disk
		err = c.HCService().Aws.Disk.SyncDisk(
			kit.Ctx,
			header,
			&protodisk.DiskSyncReq{
				AccountID: accountID,
				Region:    region.RegionID,
			},
		)
		if err != nil {
			logs.Errorf("sync do aws sync disk failed, err: %v, regionID: %s,  rid: %s",
				err, region.RegionID, kit.Rid)
		}

		// Vpc
		err = c.HCService().Aws.Vpc.SyncVpc(
			kit.Ctx,
			header,
			&proto.ResourceSyncReq{
				AccountID: accountID,
				Region:    region.RegionID,
			},
		)
		if err != nil {
			logs.Errorf("sync do aws sync vpc failed, err: %v, accountID: %s, regionID: %s, rid: %s",
				err, accountID, region.RegionID, kit.Rid)
		}

		// Subnet
		err = c.HCService().Aws.Subnet.SyncSubnet(
			kit.Ctx,
			header,
			&proto.ResourceSyncReq{
				AccountID: accountID,
				Region:    region.RegionID,
			},
		)
		if err != nil {
			logs.Errorf("sync do aws sync subnet failed, err: %v, accountID: %s, regionID: %s,  rid: %s",
				err, accountID, region.RegionID, kit.Rid)
		}
	}
	if err != nil {
		return err
	}

	return nil
}

// SyncAwsSGRule ...
func SyncAwsSGRule(c *client.ClientSet, kit *kit.Kit, header http.Header,
	region string, accountID string) error {

	start := 0
	for {
		results, err := c.DataService().Global.SecurityGroup.ListSecurityGroup(
			kit.Ctx,
			header,
			&dataproto.SecurityGroupListReq{
				Filter: &filter.Expression{
					Op: filter.And,
					Rules: []filter.RuleFactory{
						&filter.AtomRule{
							Field: "vendor",
							Op:    filter.Equal.Factory(),
							Value: enumor.Aws,
						},
						&filter.AtomRule{
							Field: "region",
							Op:    filter.Equal.Factory(),
							Value: region,
						},
						&filter.AtomRule{
							Field: "account_id",
							Op:    filter.Equal.Factory(),
							Value: accountID,
						},
					},
				},
				Page: &core.BasePage{
					Start: uint32(start),
					Limit: core.DefaultMaxPageLimit,
				},
			},
		)
		if err != nil {
			logs.Errorf("list aws security group failed, err: %v, rid: %s", err, kit.Rid)
			return err
		}

		if len(results.Details) == 0 {
			break
		}

		for _, v := range results.Details {
			err = c.HCService().Aws.SecurityGroup.SyncSecurityGroupRule(
				kit.Ctx,
				header,
				&proto.SecurityGroupSyncReq{
					AccountID: v.AccountID,
					Region:    v.Region,
				},
				v.ID,
			)
			if err != nil {
				logs.Errorf("sync do aws sync sg  rule failed, err: %v, regionID: %s, rid: %s",
					err, v.Region, kit.Rid)
			}
		}

		start += len(results.Details)
		if uint(len(results.Details)) < core.DefaultMaxPageLimit {
			break
		}
	}

	return nil
}

// SyncAwsImage ...
func SyncAwsImage(c *client.ClientSet, kit *kit.Kit, accountID string, header http.Header) error {
	regions, err := c.DataService().Aws.Region.ListRegion(
		kit.Ctx,
		header,
		&protoregion.AwsRegionListReq{
			Filter: tools.EqualExpression("vendor", enumor.Aws),
			Page:   &core.BasePage{Start: 0, Limit: core.DefaultMaxPageLimit},
		},
	)
	if err != nil {
		logs.Errorf("sync list aws region failed, err: %v, rid: %s", err, kit.Rid)
		return err
	}

	for _, region := range regions.Details {
		err = c.HCService().Aws.Image.SyncImage(
			kit.Ctx,
			header,
			&protodisk.DiskSyncReq{
				AccountID: accountID,
				Region:    region.RegionID,
			},
		)
		// sync only one time
		if err == nil {
			break
		} else {
			logs.Errorf("sync aws image failed, err: %v, rid: %s", err, kit.Rid)
		}
	}

	return err
}