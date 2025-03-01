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

package disk

import (
	"net/http"

	"hcm/cmd/hc-service/service/capability"
	"hcm/pkg/rest"
)

// InitDiskService initial the disk service
func InitDiskService(cap *capability.Capability) {
	d := &diskAdaptor{
		adaptor: cap.CloudAdaptor,
		dataCli: cap.ClientSet.DataService(),
	}

	h := rest.NewHandler()

	// 硬盘创建
	h.Add("CreateDisks", http.MethodPost, "/vendors/{vendor}/disks/create", d.CreateDisks)
	// 删除云盘
	h.Add("DeleteDisk", http.MethodDelete, "/vendors/{vendor}/disks", d.DeleteDisk)
	// 挂载云盘
	h.Add("AttachDisk", http.MethodPost, "/vendors/{vendor}/disks/attach", d.AttachDisk)
	// 卸载云盘
	h.Add("DetachDisk", http.MethodPost, "/vendors/{vendor}/disks/detach", d.DetachDisk)

	h.Load(cap.WebService)
}
