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

package handler

import (
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/iam/meta"
	"hcm/pkg/rest"
	"hcm/pkg/runtime/filter"
)

// RecycleValidWithAuth validate and authorize cloud resource for recycle bin manager handler
func RecycleValidWithAuth(cts *rest.Contexts, opt *ValidWithAuthOption) error {
	// authorize one resource
	if opt.BasicInfo != nil {
		if opt.BasicInfo.RecycleStatus != enumor.RecycleStatus {
			return errf.New(errf.InvalidParameter, "resource is not in recycle bin")
		}

		authRes := meta.ResourceAttribute{Basic: &meta.Basic{Type: meta.RecycleBin, Action: opt.Action,
			ResourceID: opt.BasicInfo.AccountID}, BizID: opt.BasicInfo.BkBizID}
		return opt.Authorizer.AuthorizeWithPerm(cts.Kit, authRes)
	}

	// batch authorize resource
	authRes := make([]meta.ResourceAttribute, 0, len(opt.BasicInfos))
	for _, info := range opt.BasicInfos {
		if info.RecycleStatus != enumor.RecycleStatus {
			return errf.New(errf.InvalidParameter, "resource is not in recycle bin")
		}

		authRes = append(authRes, meta.ResourceAttribute{Basic: &meta.Basic{Type: meta.RecycleBin, Action: opt.Action,
			ResourceID: info.AccountID}, BizID: info.BkBizID})
	}

	return opt.Authorizer.AuthorizeWithPerm(cts.Kit, authRes...)
}

// BizRecycleValidWithAuth validate and authorize cloud resource for biz recycle bin manager handler
func BizRecycleValidWithAuth(cts *rest.Contexts, opt *ValidWithAuthOption) error {
	bizID, err := cts.PathParameter("bk_biz_id").Int64()
	if err != nil {
		return err
	}

	if bizID <= 0 {
		return errf.New(errf.InvalidParameter, "resource is not in biz")
	}

	// authorize one resource
	if opt.BasicInfo != nil {
		if opt.BasicInfo.RecycleStatus != enumor.RecycleStatus {
			return errf.New(errf.InvalidParameter, "resource is not in recycle bin")
		}

		authRes := meta.ResourceAttribute{Basic: &meta.Basic{Type: meta.RecycleBin, Action: opt.Action,
			ResourceID: opt.BasicInfo.AccountID}, BizID: bizID}
		return opt.Authorizer.AuthorizeWithPerm(cts.Kit, authRes)
	}

	// batch authorize resource
	authRes := make([]meta.ResourceAttribute, 0, len(opt.BasicInfos))
	for _, info := range opt.BasicInfos {
		if info.RecycleStatus != enumor.RecycleStatus {
			return errf.New(errf.InvalidParameter, "resource is not in recycle bin")
		}

		authRes = append(authRes, meta.ResourceAttribute{Basic: &meta.Basic{Type: meta.RecycleBin, Action: opt.Action,
			ResourceID: info.AccountID}, BizID: bizID})
	}

	return opt.Authorizer.AuthorizeWithPerm(cts.Kit, authRes...)
}

// ListResourceRecycleAuthRes list authorized recycled resource for resource manager.
func ListResourceRecycleAuthRes(cts *rest.Contexts, opt *ListAuthResOption) (*filter.Expression, bool, error) {
	authRes := meta.ResourceAttribute{Basic: &meta.Basic{Type: meta.RecycleBin, Action: opt.Action}}
	_, authorized, err := opt.Authorizer.Authorize(cts.Kit, authRes)
	if err != nil {
		return nil, false, err
	}

	if !authorized {
		return nil, true, nil
	}

	return opt.Filter, false, err
}

// ListBizRecycleAuthRes list authorized recycled biz resource for resource manager.
func ListBizRecycleAuthRes(cts *rest.Contexts, opt *ListAuthResOption) (*filter.Expression, bool, error) {
	opt.ResType = meta.RecycleBin
	return ListBizAuthRes(cts, opt)
}
