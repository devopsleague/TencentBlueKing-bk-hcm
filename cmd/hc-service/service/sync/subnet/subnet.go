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

// Package subnet defines vpc service.
package subnet

import (
	"hcm/cmd/hc-service/service/capability"
	cloudadaptor "hcm/cmd/hc-service/service/cloud-adaptor"
	proto "hcm/pkg/api/hc-service"
	"hcm/pkg/client"
	dataservice "hcm/pkg/client/data-service"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/rest"
)

// InitSyncSubnetService initial the subnet service
func InitSyncSubnetService(cap *capability.Capability) {
	v := &syncSubnetSvc{
		ad:      cap.CloudAdaptor,
		cs:      cap.ClientSet,
		dataCli: cap.ClientSet.DataService(),
	}

	h := rest.NewHandler()

	// vpc sync
	h.Add("TCloudVpcSync", "POST", "/vendors/tcloud/subnets/sync", v.SyncTCloudSubnet)
	h.Add("HuaWeiVpcSync", "POST", "/vendors/huawei/subnets/sync", v.SyncHuaWeiSubnet)
	h.Add("AwsVpcSync", "POST", "/vendors/aws/subnets/sync", v.SyncAwsSubnet)
	h.Add("AzureVpcSync", "POST", "/vendors/azure/subnets/sync", v.SyncAzureSubnet)
	h.Add("GcpVpcSync", "POST", "/vendors/gcp/subnets/sync", v.SyncGcpSubnet)

	h.Load(cap.WebService)
}

type syncSubnetSvc struct {
	ad      *cloudadaptor.CloudAdaptorClient
	cs      *client.ClientSet
	dataCli *dataservice.Client
}

func decodeSubnetSyncReq(cts *rest.Contexts) (*proto.ResourceSyncReq, error) {
	req := new(proto.ResourceSyncReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	return req, nil
}