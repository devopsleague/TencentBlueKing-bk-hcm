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

package cloud

import (
	"hcm/cmd/hc-service/service/capability"
	"hcm/pkg/adaptor"
	"hcm/pkg/rest"
)

// InitDiskService initial the disk service
func InitDiskService(cap *capability.Capability) {
	d := &disk{
		ad: cap.Adaptor,
	}

	h := rest.NewHandler()

	h.Add("CreateDisks", "POST", "/vendors/{vendor}/disks/create", d.CreateDisks)

	h.Load(cap.WebService)
}

type disk struct {
	ad *adaptor.Adaptor
}

//type createFunc func()
//
//var CreateMethods map[enumor.Vendor]createFunc

// CreateDisks 创建云硬盘
func (d *disk) CreateDisks(cts *rest.Contexts) (interface{}, error) {
	// vendor := enumor.Vendor(cts.Request.PathParameter("vendor"))
	return nil, nil
}