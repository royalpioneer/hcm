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

package huawei

import (
	"errors"

	logicsaccount "hcm/cmd/cloud-server/logics/account"
)

// CheckReq 检查申请单的数据是否正确
func (a *ApplicationOfCreateHuaWeiCvm) CheckReq() error {
	if err := a.req.Validate(); err != nil {
		return err
	}

	if err := logicsaccount.IsResourceAccount(a.Cts.Kit, a.Client.DataService(), a.req.AccountID); err != nil {
		return err
	}

	// TCloud 支持 DryRun，可预校验
	result, err := a.Client.HCService().HuaWei.Cvm.BatchCreateCvm(
		a.Cts.Kit.Ctx,
		a.Cts.Kit.Header(),
		a.toHcProtoHuaWeiBatchCreateReq(true),
	)
	if err != nil {
		return err
	}
	if result != nil && result.FailedMessage != "" {
		return errors.New(result.FailedMessage)
	}

	return nil
}
