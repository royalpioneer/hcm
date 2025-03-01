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

// GcpDiskExtensionCreateReq ...
type GcpDiskExtensionCreateReq struct {
	SelfLink    string `json:"self_link" validate:"required"`
	SourceImage string `json:"source_image,omitempty"`
	Description string `json:"description,omitempty"`
	Encrypted   *bool  `json:"encrypted,omitempty"`
}

// GcpDiskExtensionResult ...
type GcpDiskExtensionResult struct {
	SelfLink    string `json:"self_link" validate:"required"`
	SourceImage string `json:"source_image,omitempty"`
	Description string `json:"description,omitempty"`
	Encrypted   *bool  `json:"encrypted,omitempty"`
}

// GcpDiskExtensionUpdateReq ...
// 根据情况增加 omitempty tag, 因为会调用 json.UpdateMerge 完成字段合并
type GcpDiskExtensionUpdateReq struct {
	SelfLink    string `json:"self_link" validate:"required"`
	SourceImage string `json:"source_image,omitempty"`
	Description string `json:"description,omitempty"`
	Encrypted   *bool  `json:"encrypted,omitempty"`
}
