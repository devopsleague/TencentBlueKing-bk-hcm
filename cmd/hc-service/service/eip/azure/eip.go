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

package azure

import (
	cloudclient "hcm/cmd/hc-service/service/cloud-adaptor"
	"hcm/cmd/hc-service/service/eip/datasvc"
	"hcm/pkg/adaptor/azure"
	"hcm/pkg/adaptor/types/core"
	"hcm/pkg/adaptor/types/eip"
	proto "hcm/pkg/api/hc-service/eip"
	dataservice "hcm/pkg/client/data-service"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/kit"
	"hcm/pkg/rest"
)

// EipSvc ...
type EipSvc struct {
	Adaptor *cloudclient.CloudAdaptorClient
	DataCli *dataservice.Client
}

// DeleteEip ...
func (svc *EipSvc) DeleteEip(cts *rest.Contexts) (interface{}, error) {
	req := new(proto.EipDeleteReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	opt, err := svc.makeEipDeleteOption(cts.Kit, req)
	if err != nil {
		return nil, err
	}

	client, err := svc.Adaptor.Azure(cts.Kit, req.AccountID)
	if err != nil {
		return nil, err
	}

	err = client.DeleteEip(cts.Kit, opt)
	if err != nil {
		return nil, err
	}

	manager := datasvc.EipManager{DataCli: svc.DataCli}
	return nil, manager.Delete(cts.Kit, []string{req.EipID})
}

// AssociateEip ...
func (svc *EipSvc) AssociateEip(cts *rest.Contexts) (interface{}, error) {
	req := new(proto.AzureEipAssociateReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	client, err := svc.Adaptor.Azure(cts.Kit, req.AccountID)
	if err != nil {
		return nil, err
	}

	opt, err := svc.makeEipAssociateOption(cts.Kit, req, client)
	if err != nil {
		return nil, err
	}

	err = client.AssociateEip(cts.Kit, opt)
	if err != nil {
		return nil, err
	}

	manager := datasvc.EipCvmRelManager{CvmID: req.CvmID, EipID: req.EipID, DataCli: svc.DataCli}
	return nil, manager.Create(cts.Kit)
}

// DisassociateEip ...
func (svc *EipSvc) DisassociateEip(cts *rest.Contexts) (interface{}, error) {
	req := new(proto.AzureEipDisassociateReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	client, err := svc.Adaptor.Azure(cts.Kit, req.AccountID)
	if err != nil {
		return nil, err
	}

	opt, err := svc.makeEipDisassociateOption(cts.Kit, req, client)
	if err != nil {
		return nil, err
	}

	err = client.DisassociateEip(cts.Kit, opt)
	if err != nil {
		return nil, err
	}

	manager := datasvc.EipCvmRelManager{CvmID: req.CvmID, EipID: req.EipID, DataCli: svc.DataCli}
	return nil, manager.Delete(cts.Kit)
}

func (svc *EipSvc) makeEipDeleteOption(
	kt *kit.Kit,
	req *proto.EipDeleteReq,
) (*eip.AzureEipDeleteOption, error) {
	eipData, err := svc.DataCli.Azure.RetrieveEip(kt.Ctx, kt.Header(), req.EipID)
	if err != nil {
		return nil, err
	}
	return &eip.AzureEipDeleteOption{
		ResourceGroupName: eipData.Extension.ResourceGroupName,
		EipName:           *eipData.Name,
	}, nil
}

func (svc *EipSvc) makeEipAssociateOption(
	kt *kit.Kit,
	req *proto.AzureEipAssociateReq,
	cli *azure.Azure,
) (*eip.AzureEipAssociateOption, error) {
	dataCli := svc.DataCli.Azure

	eipData, err := dataCli.RetrieveEip(kt.Ctx, kt.Header(), req.EipID)
	if err != nil {
		return nil, err
	}

	networkInterface, err := dataCli.NetworkInterface.Get(kt.Ctx, kt.Header(), req.NetworkInterfaceID)
	if err != nil {
		return nil, err
	}
	networkInterfaces, err := cli.ListRawNetworkInterfaceByIDs(
		kt,
		&core.AzureListByIDOption{CloudIDs: []string{networkInterface.CloudID}},
	)
	if err != nil {
		return nil, err
	}

	return &eip.AzureEipAssociateOption{
		ResourceGroupName: eipData.Extension.ResourceGroupName,
		CloudEipID:        eipData.CloudID,
		NetworkInterface:  networkInterfaces[0],
	}, nil
}

func (svc *EipSvc) makeEipDisassociateOption(
	kt *kit.Kit,
	req *proto.AzureEipDisassociateReq,
	cli *azure.Azure,
) (*eip.AzureEipDisassociateOption, error) {
	dataCli := svc.DataCli.Azure

	eipData, err := dataCli.RetrieveEip(kt.Ctx, kt.Header(), req.EipID)
	if err != nil {
		return nil, err
	}

	networkInterface, err := dataCli.NetworkInterface.Get(kt.Ctx, kt.Header(), req.NetworkInterfaceID)
	if err != nil {
		return nil, err
	}
	networkInterfaces, err := cli.ListRawNetworkInterfaceByIDs(
		kt,
		&core.AzureListByIDOption{CloudIDs: []string{networkInterface.CloudID}},
	)
	if err != nil {
		return nil, err
	}

	return &eip.AzureEipDisassociateOption{
		ResourceGroupName: eipData.Extension.ResourceGroupName,
		CloudEipID:        eipData.CloudID,
		NetworkInterface:  networkInterfaces[0],
	}, nil
}