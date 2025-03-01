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

package bill

import (
	"strings"

	typesBill "hcm/pkg/adaptor/types/bill"
	"hcm/pkg/adaptor/types/core"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/criteria/validator"
)

// -------------------------- List --------------------------

// AwsBillListReq define aws bill list req.
type AwsBillListReq struct {
	AccountID string `json:"account_id" validate:"required"`
	// 起始日期，格式为yyyy-mm-dd，不支持跨月查询
	BeginDate string `json:"begin_date" validate:"required"`
	// 截止日期，格式为yyyy-mm-dd，不支持跨月查询
	EndDate string           `json:"end_date" validate:"required"`
	Page    *AwsBillListPage `json:"page" validate:"omitempty"`
}

// Validate aws bill list req.
func (opt AwsBillListReq) Validate() error {
	if err := validator.Validate.Struct(opt); err != nil {
		return err
	}

	if opt.Page != nil {
		if err := opt.Page.Validate(); err != nil {
			return err
		}
	}

	if (opt.BeginDate == "" && opt.EndDate != "") || (opt.BeginDate != "" && opt.EndDate == "") {
		return errf.New(errf.InvalidParameter, "begin_date and end_date can not be empty")
	}

	beginArr := strings.Split(opt.BeginDate, "-")
	if len(beginArr) != 3 {
		return errf.New(errf.InvalidParameter, "begin_date is invalid")
	}

	endArr := strings.Split(opt.EndDate, "-")
	if len(endArr) != 3 {
		return errf.New(errf.InvalidParameter, "end_date is invalid")
	}

	return nil
}

// AwsBillListPage defines aws bill list page.
type AwsBillListPage struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

// Validate validate aws bill list page.
func (opt AwsBillListPage) Validate() error {
	if opt.Limit == 0 {
		return errf.New(errf.InvalidParameter, "limit is required")
	}

	if opt.Limit > core.AwsQueryLimit {
		return errf.New(errf.InvalidParameter, "aws.limit should <= 1000")
	}

	return nil
}

// -------------------------- List --------------------------

// TCloudBillListReq define tcloud bill list req.
type TCloudBillListReq struct {
	AccountID string `json:"account_id" validate:"required"`
	// 月份，格式为yyyy-mm，不能早于开通账单2.0的月份，最多可拉取24个月内的数据,不支持跨月查询
	Month string `json:"month" validate:"omitempty"`
	// 起始日期，周期开始时间，格式为Y-m-d H:i:s，Month和BeginDate&EndDate必传一个，如果有该字段则Month字段无效
	BeginDate string `json:"begin_date" validate:"omitempty"`
	// 截止日期，周期结束时间，格式为Y-m-d H:i:s，Month和BeginDate&EndDate必传一个，如果有该字段则Month字段无效
	EndDate string `json:"end_date" validate:"omitempty"`
	// Limit: 最大值为100
	Page *core.TCloudPage `json:"page" validate:"omitempty"`
}

// Validate tcloud bill list req.
func (opt TCloudBillListReq) Validate() error {
	if err := validator.Validate.Struct(opt); err != nil {
		return err
	}

	if opt.Month == "" && opt.BeginDate == "" && opt.EndDate == "" {
		return errf.New(errf.InvalidParameter, "month and begin_date and end_date can not be empty")
	}

	if opt.Page != nil {
		if err := opt.Page.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// HuaWeiBillListReq defines huawei bill list req.
type HuaWeiBillListReq struct {
	AccountID string `json:"account_id" validate:"required"`
	// 查询的资源详单所在账期,东八区时间,格式为YYYY-MM。 示例:2019-01 说明: 不支持2019年1月份之前的资源详单。
	Month string `json:"month" validate:"required"`
	// Limit: 最大值为1000
	Page *typesBill.HuaWeiBillPage `json:"page" validate:"omitempty"`
}

// Validate huawei bill list req.
func (opt HuaWeiBillListReq) Validate() error {
	if err := validator.Validate.Struct(opt); err != nil {
		return err
	}

	if opt.Page != nil {
		if err := opt.Page.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// AzureBillListReq define azure bill list req.
type AzureBillListReq struct {
	AccountID string `json:"account_id" validate:"required"`
	// 起始日期，格式为yyyy-mm-dd，不支持跨月查询
	BeginDate string `json:"begin_date" validate:"required"`
	// 截止日期，格式为yyyy-mm-dd，不支持跨月查询
	EndDate string                   `json:"end_date" validate:"required"`
	Page    *typesBill.AzureBillPage `json:"page" validate:"omitempty"`
}

// Validate azure bill list req.
func (opt AzureBillListReq) Validate() error {
	if err := validator.Validate.Struct(opt); err != nil {
		return err
	}

	if opt.Page != nil {
		if err := opt.Page.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GcpBillListReq defines gcp bill list req.
type GcpBillListReq struct {
	AccountID string `json:"account_id" validate:"required"`
	// 包含费用专列项的账单的年份和月份，格式为YYYYMM 示例:201901，可以使用此字段获取账单上的总费用
	Month string `json:"month" validate:"omitempty"`
	// 起始时间戳，时间戳值表示绝对时间点，与任何时区或惯例（如夏令时）无关，可精确到微秒，
	// 格式：0001-01-01 00:00:00 至 9999-12-31 23:59:59.999999（世界协调时间 (UTC)）
	// 也可以使用UTC格式：2014-09-27T12:30:00.45Z
	BeginDate string `json:"begin_date" validate:"omitempty"`
	// 截止时间戳，时间戳值表示绝对时间点，与任何时区或惯例（如夏令时）无关，可精确到微秒
	EndDate string                 `json:"end_date" validate:"omitempty"`
	Page    *typesBill.GcpBillPage `json:"page" validate:"omitempty"`
}

// Validate gcp bill list req.
func (opt GcpBillListReq) Validate() error {
	if err := validator.Validate.Struct(opt); err != nil {
		return err
	}

	if opt.Month == "" && opt.BeginDate == "" && opt.EndDate == "" {
		return errf.New(errf.InvalidParameter, "month and begin_date and end_date can not be empty")
	}

	if opt.Page != nil {
		if err := opt.Page.Validate(); err != nil {
			return err
		}
	}

	return nil
}
