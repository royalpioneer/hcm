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

package routetable

// TCloudRoute defines tencent cloud route struct.
type TCloudRoute struct {
	// 预留为string类型，以便后续腾讯云如果要切换为RouteItemId的话也可以兼容
	CloudID                  string  `json:"cloud_id"`
	CloudRouteTableID        string  `json:"cloud_route_table_id"`
	DestinationCidrBlock     string  `json:"destination_cidr_block"`
	DestinationIpv6CidrBlock *string `json:"destination_ipv6_cidr_block,omitempty"`
	GatewayType              string  `json:"gateway_type"`
	CloudGatewayID           string  `json:"cloud_gateway_id"`
	Enabled                  bool    `json:"enabled"`
	RouteType                string  `json:"route_type"`
	PublishedToVbc           bool    `json:"published_to_vbc"`
	Memo                     *string `json:"memo,omitempty"`
}
