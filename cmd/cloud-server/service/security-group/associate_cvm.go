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

package securitygroup

import (
	"hcm/cmd/cloud-server/service/cvm"
	proto "hcm/pkg/api/cloud-server"
	protoaudit "hcm/pkg/api/data-service/audit"
	hcproto "hcm/pkg/api/hc-service"
	"hcm/pkg/criteria/constant"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/iam/meta"
	"hcm/pkg/logs"
	"hcm/pkg/rest"
	"hcm/pkg/runtime/filter"
)

// AssociateCvm ...
func (svc *securityGroupSvc) AssociateCvm(cts *rest.Contexts) (interface{}, error) {
	req, vendor, err := svc.decodeAndValidateAssocCvmReq(cts, meta.Associate)
	if err != nil {
		return nil, err
	}

	// create operation audit.
	audit := protoaudit.CloudResourceOperationInfo{
		ResType:           enumor.SecurityGroupRuleAuditResType,
		ResID:             req.SecurityGroupID,
		Action:            protoaudit.Associate,
		AssociatedResType: enumor.CvmAuditResType,
		AssociatedResID:   req.CvmID,
	}
	if err = svc.audit.ResOperationAudit(cts.Kit, audit); err != nil {
		logs.Errorf("create operation audit failed, err: %v, rid: %s", err, cts.Kit.Rid)
		return nil, err
	}

	switch vendor {
	case enumor.TCloud:
		associateReq := &hcproto.SecurityGroupAssociateCvmReq{
			SecurityGroupID: req.SecurityGroupID,
			CvmID:           req.CvmID,
		}
		err = svc.client.HCService().TCloud.SecurityGroup.AssociateCvm(cts.Kit.Ctx, cts.Kit.Header(),
			associateReq)

	case enumor.HuaWei:
		associateReq := &hcproto.SecurityGroupAssociateCvmReq{
			SecurityGroupID: req.SecurityGroupID,
			CvmID:           req.CvmID,
		}
		err = svc.client.HCService().HuaWei.SecurityGroup.AssociateCvm(cts.Kit.Ctx, cts.Kit.Header(),
			associateReq)

	case enumor.Aws:
		associateReq := &hcproto.SecurityGroupAssociateCvmReq{
			SecurityGroupID: req.SecurityGroupID,
			CvmID:           req.CvmID,
		}
		err = svc.client.HCService().Aws.SecurityGroup.AssociateCvm(cts.Kit.Ctx, cts.Kit.Header(),
			associateReq)

	default:
		return nil, errf.Newf(errf.Unknown, "vendor: %s not support", vendor)
	}

	if err != nil {
		logs.Errorf("security group associate cvm failed, err: %v, req: %+v, rid: %s", err, req, cts.Kit.Rid)
		return nil, err
	}

	return nil, nil
}

// DisassociateCvm ...
func (svc *securityGroupSvc) DisassociateCvm(cts *rest.Contexts) (interface{}, error) {
	req, vendor, err := svc.decodeAndValidateAssocCvmReq(cts, meta.Disassociate)
	if err != nil {
		return nil, err
	}

	// create operation audit.
	audit := protoaudit.CloudResourceOperationInfo{
		ResType:           enumor.SecurityGroupRuleAuditResType,
		ResID:             req.SecurityGroupID,
		Action:            protoaudit.Disassociate,
		AssociatedResType: enumor.CvmAuditResType,
		AssociatedResID:   req.CvmID,
	}
	if err = svc.audit.ResOperationAudit(cts.Kit, audit); err != nil {
		logs.Errorf("create operation audit failed, err: %v, rid: %s", err, cts.Kit.Rid)
		return nil, err
	}

	switch vendor {
	case enumor.TCloud:
		associateReq := &hcproto.SecurityGroupAssociateCvmReq{
			SecurityGroupID: req.SecurityGroupID,
			CvmID:           req.CvmID,
		}
		err = svc.client.HCService().TCloud.SecurityGroup.DisassociateCvm(cts.Kit.Ctx, cts.Kit.Header(),
			associateReq)

	case enumor.HuaWei:
		associateReq := &hcproto.SecurityGroupAssociateCvmReq{
			SecurityGroupID: req.SecurityGroupID,
			CvmID:           req.CvmID,
		}
		err = svc.client.HCService().HuaWei.SecurityGroup.DisassociateCvm(cts.Kit.Ctx, cts.Kit.Header(),
			associateReq)

	case enumor.Aws:
		associateReq := &hcproto.SecurityGroupAssociateCvmReq{
			SecurityGroupID: req.SecurityGroupID,
			CvmID:           req.CvmID,
		}
		err = svc.client.HCService().Aws.SecurityGroup.DisassociateCvm(cts.Kit.Ctx, cts.Kit.Header(),
			associateReq)

	default:
		return nil, errf.Newf(errf.Unknown, "vendor: %s not support", vendor)
	}

	if err != nil {
		logs.Errorf("security group disassociate cvm failed, err: %v, req: %+v, rid: %s", err, req, cts.Kit.Rid)
		return nil, err
	}

	return nil, nil
}

func (svc *securityGroupSvc) decodeAndValidateAssocCvmReq(cts *rest.Contexts, action meta.Action) (
	*proto.SecurityGroupAssociateCvmReq, enumor.Vendor, error) {

	req := new(proto.SecurityGroupAssociateCvmReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, "", errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	if err := req.Validate(); err != nil {
		return nil, "", errf.NewFromErr(errf.InvalidParameter, err)
	}

	basicInfo, err := svc.client.DataService().Global.Cloud.GetResourceBasicInfo(cts.Kit.Ctx, cts.Kit.Header(),
		enumor.SecurityGroupCloudResType, req.SecurityGroupID)
	if err != nil {
		logs.Errorf("get resource vendor failed, id: %s, err: %s, rid: %s", basicInfo, err, cts.Kit.Rid)
		return nil, "", err
	}

	// authorize
	authRes := meta.ResourceAttribute{Basic: &meta.Basic{Type: meta.SecurityGroup, Action: action,
		ResourceID: basicInfo.AccountID}}
	err = svc.authorizer.AuthorizeWithPerm(cts.Kit, authRes)
	if err != nil {
		return nil, "", err
	}

	// 已分配业务的资源，不允许操作
	flt := &filter.AtomRule{Field: "id", Op: filter.Equal.Factory(), Value: req.SecurityGroupID}
	err = CheckSecurityGroupsInBiz(cts.Kit, svc.client, flt, constant.UnassignedBiz)
	if err != nil {
		return nil, "", err
	}

	flt = &filter.AtomRule{Field: "id", Op: filter.Equal.Factory(), Value: req.CvmID}
	err = cvm.CheckCvmsInBiz(cts.Kit, svc.client, flt, constant.UnassignedBiz)
	if err != nil {
		return nil, "", err
	}

	return req, basicInfo.Vendor, nil
}