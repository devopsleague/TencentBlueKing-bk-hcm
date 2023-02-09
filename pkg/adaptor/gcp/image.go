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

package gcp

import (
	"strconv"

	"hcm/pkg/adaptor/types/image"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
	"hcm/pkg/runtime/filter"
)

// ListImage ...
// reference: https://cloud.google.com/compute/docs/reference/rest/v1/images/list
func (g *Gcp) ListImage(
	kt *kit.Kit,
	opt *image.GcpImageListOption,
) (*image.GcpImageListResult, string, error) {

	client, err := g.clientSet.computeClient(kt)
	if err != nil {
		return nil, "", err
	}

	req := client.Images.List(opt.ProjectID).Context(kt.Ctx)
	req.MaxResults(int64(filter.DefaultMaxInLimit))

	images := make([]image.GcpImage, 0)

	resp, err := req.PageToken(opt.NextPageToken).Do()
	if err != nil {
		logs.Errorf("list images failed, err: %v, rid: %s", err, kt.Rid)
		return nil, "", err
	}
	for _, pImage := range resp.Items {
		images = append(images, image.GcpImage{
			CloudID:      strconv.FormatUint(pImage.Id, 10),
			Name:         pImage.Name,
			Architecture: pImage.Architecture,
			State:        pImage.Status,
			Type:         "public",
		})
	}

	return &image.GcpImageListResult{Details: images}, resp.NextPageToken, nil
}