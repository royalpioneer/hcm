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
	ressync "hcm/cmd/hc-service/logics/res-sync"
	"hcm/cmd/hc-service/logics/res-sync/gcp"
	"hcm/cmd/hc-service/service/sync/handler"
	"hcm/pkg/adaptor/types"
	typecore "hcm/pkg/adaptor/types/core"
	"hcm/pkg/api/hc-service/sync"
	"hcm/pkg/criteria/constant"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
	"hcm/pkg/rest"
)

// SyncVpc ....
func (svc *service) SyncVpc(cts *rest.Contexts) (interface{}, error) {
	return nil, handler.ResourceSync(cts, &vpcHandler{cli: svc.syncCli})
}

// vpcHandler vpc sync handler.
type vpcHandler struct {
	cli ressync.Interface

	// Prepare 构建参数
	request   *sync.GcpGlobalRegionResSyncReq
	syncCli   gcp.Interface
	pageToken string
}

var _ handler.Handler = new(vpcHandler)

// Prepare ...
func (hd *vpcHandler) Prepare(cts *rest.Contexts) error {
	req := new(sync.GcpGlobalRegionResSyncReq)
	if err := cts.DecodeInto(req); err != nil {
		return errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	if err := req.Validate(); err != nil {
		return errf.NewFromErr(errf.InvalidParameter, err)
	}

	syncCli, err := hd.cli.Gcp(cts.Kit, req.AccountID)
	if err != nil {
		return err
	}

	hd.request = req
	hd.syncCli = syncCli

	return nil
}

// Next ...
func (hd *vpcHandler) Next(kt *kit.Kit) ([]string, error) {
	listOpt := &types.GcpListOption{
		Page: &typecore.GcpPage{
			PageSize:  constant.CloudResourceSyncMaxLimit,
			PageToken: hd.pageToken,
		},
	}

	vpcResult, err := hd.syncCli.CloudCli().ListVpc(kt, listOpt)
	if err != nil {
		logs.Errorf("request adaptor list gcp vpc failed, err: %v, opt: %v, rid: %s", err, listOpt, kt.Rid)
		return nil, err
	}

	if len(vpcResult.Details) == 0 {
		return nil, nil
	}

	cloudIDs := make([]string, 0, len(vpcResult.Details))
	for _, one := range vpcResult.Details {
		cloudIDs = append(cloudIDs, one.CloudID)
	}

	hd.pageToken = vpcResult.NextPageToken
	return cloudIDs, nil
}

// Sync ...
func (hd *vpcHandler) Sync(kt *kit.Kit, cloudIDs []string) error {
	params := &gcp.SyncBaseParams{
		AccountID: hd.request.AccountID,
		CloudIDs:  cloudIDs,
	}
	if _, err := hd.syncCli.Vpc(kt, params, new(gcp.SyncVpcOption)); err != nil {
		logs.Errorf("sync gcp vpc failed, err: %v, opt: %v, rid: %s", err, params, kt.Rid)
		return err
	}

	return nil
}

// RemoveDeleteFromCloud ...
func (hd *vpcHandler) RemoveDeleteFromCloud(kt *kit.Kit) error {
	if err := hd.syncCli.RemoveVpcDeleteFromCloud(kt, hd.request.AccountID); err != nil {
		logs.Errorf("remove vpc delete from cloud failed, err: %v, accountID: %s, rid: %s", err,
			hd.request.AccountID, kt.Rid)
		return err
	}

	return nil
}

// Name ...
func (hd *vpcHandler) Name() enumor.CloudResourceType {
	return enumor.VpcCloudResType
}
