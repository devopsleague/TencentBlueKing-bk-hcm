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

package lblogic

import (
	"encoding/json"

	dataservice "hcm/pkg/client/data-service"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
)

var _ ImportValidator = (*CreateLayer4ListenerValidator)(nil)

func newCreateLayer4ListenerValidator(cli *dataservice.Client, vendor enumor.Vendor, bkBizID int64,
	accountID string, regionIDs []string) *CreateLayer4ListenerValidator {

	return &CreateLayer4ListenerValidator{
		newBasePreviewExecutor(cli, vendor, bkBizID, accountID, regionIDs),
	}
}

// CreateLayer4ListenerValidator ...
type CreateLayer4ListenerValidator struct {
	*basePreviewExecutor
}

// Validate ...
func (c *CreateLayer4ListenerValidator) Validate(kt *kit.Kit, rawData json.RawMessage) (interface{}, error) {
	executor := &CreateLayer4ListenerPreviewExecutor{
		basePreviewExecutor: c.basePreviewExecutor,
	}
	err := json.Unmarshal(rawData, &executor.details)
	if err != nil {
		return nil, err
	}

	// reset status and validateResult
	for _, detail := range executor.details {
		detail.Status = ""
		detail.ValidateResult = make([]string, 0)
	}

	if err = executor.validate(kt); err != nil {
		logs.Errorf("validate failed,operationType: %s, err: %v, rid: %s", CreateLayer4Listener, err, kt.Rid)
		return nil, err
	}
	return executor.details, nil
}
