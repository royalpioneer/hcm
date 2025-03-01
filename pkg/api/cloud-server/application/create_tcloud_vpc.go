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

package application

import (
	"hcm/pkg/criteria/validator"
)

// TCloudVpcCreateReq ...
type TCloudVpcCreateReq struct {
	BkBizID   int64  `json:"bk_biz_id" validate:"required,min=1"`
	AccountID string `json:"account_id" validate:"required"`
	Region    string `json:"region" validate:"required"`
	Name      string `json:"name" validate:"required,min=1,max=60"`
	IPv4Cidr  string `json:"ipv4_cidr" validate:"required,cidrv4"`
	BkCloudID int64  `json:"bk_cloud_id" validate:"required,min=1"`

	Subnet struct {
		Name     string `json:"name" validate:"required,min=1,max=60"`
		IPv4Cidr string `json:"ipv4_cidr" validate:"required,cidrv4"`
		Zone     string `json:"zone" validate:"required"`
	} `json:"subnet" validate:"required"`

	Memo *string `json:"memo" validate:"omitempty"`
}

// Validate ...
func (req *TCloudVpcCreateReq) Validate() error {
	if err := validator.Validate.Struct(req); err != nil {
		return err
	}

	return nil
}
