<script lang="ts" setup>
// import InfoList from '../../../common/info-list/info-list';
import DetailInfo from '@/views/resource/resource-manage/common/info/detail-info';
import { ImageTypeEnum } from '@/typings';

import {
  PropType,
} from 'vue';
import { TypeEnum, useRouteLinkBtn } from '@/hooks/useRouteLinkBtn';
import { CLOUD_HOST_STATUS } from '@/common/constant';

const props = defineProps({
  data: {
    type: Object as PropType<any>,
  },
});

const cvmInfo = [
  {
    name: '实例名称',
    prop: 'name',
  },
  {
    name: '实例ID',
    prop: 'cloud_id',
  },
  {
    name: '账号',
    prop: 'account_id',
    render: () =>  useRouteLinkBtn(props.data, { 
      id: 'account_id',
      name: 'account_id',
      type: TypeEnum.ACCOUNT
    })
  },
  {
    name: '云厂商',
    prop: 'vendorName',
  },
  {
    name: '地域',
    prop: 'region',
  },
  {
    name: '可用区域',
    prop: 'zone',
  },
  {
    name: '业务',
    prop: 'bk_biz_id',
    render() {
      return `${props.data.bk_biz_id_name} (${props.data.bk_biz_id})`
    }
  },
  {
    name: '启动时间',
    prop: 'cloud_launched_time',
  },
  {
    name: '创建时间',
    prop: 'cloud_created_time',
  },
  {
    name: '过期时间',
    prop: 'cloud_expired_time',
  },
  {
    name: '当前状态',
    prop: 'status',
    render() {
      return CLOUD_HOST_STATUS[props.data.status];
    },
  },
  {
    name: '实例销毁保护',
    prop: 'disable_api_termination',
    render() {
      return props.data.disable_api_termination ? '是' : '否';
    },
  },
  {
    name: '备注',
    prop: 'memo',
  },
];

const netInfo = [
  {
    name: '所属网络',
    prop: 'cloud_vpc_ids',
    render: () =>  useRouteLinkBtn(props.data, { 
      id: 'vpc_ids',
      name: 'cloud_vpc_ids',
      type: TypeEnum.VPC
    })
  },
  {
    name: '所属子网',
    prop: 'cloud_subnet_ids',
    render: () =>  useRouteLinkBtn(props.data, { 
      id: 'subnet_ids',
      name: 'cloud_subnet_ids',
      type: TypeEnum.SUBNET
    })
  },
  {
    name: '私有 IPv4 地址',
    prop: 'private_ipv4_addresses',
  },
  {
    name: '公有 IPv4 地址',
    prop: 'public_ipv4_addresses',
  },
  {
    name: '私有 IPv6 地址',
    prop: 'private_ipv6_addresses',
  },
  {
    name: '公有IPv6 地址',
    prop: 'public_ipv6_addresses',
  },
];

const expandNetInfo = [
  {
    name: '所属网络',
    prop: 'cloud_vpc_ids',
    render: () =>  useRouteLinkBtn(props.data, { 
      id: 'vpc_ids',
      name: 'cloud_vpc_ids',
      type: TypeEnum.VPC
    })
  },
  {
    name: '所属子网',
    prop: 'cloud_subnet_ids',
    render: () =>  useRouteLinkBtn(props.data, { 
      id: 'subnet_ids',
      name: 'cloud_subnet_ids',
      type: TypeEnum.SUBNET,
      isExpand: true
    })
  },
  {
    name: '拓展IPv4内网IP',
    prop: 'private_ipv4_addresses',
  },
  {
    name: '拓展IPv4公网IP',
    prop: 'public_ipv4_addresses',
  },
];

const settingInfo = [
  {
    name: '实例规格',
    prop: 'machine_type',
  },
  {
    name: 'CPU',
    render() {
      return `${props.data?.flavor?.vcpus}核`;
    },
  },
  {
    name: '内存',
    render() {
      return `${props.data?.flavor?.ram}G`;
    },
  },
  {
    name: '操作系统',
    prop: 'os_name',
  },
  {
    name: '镜像id',
    prop: 'cloud_image_id',
    render: () => useRouteLinkBtn(props.data, {
      id: 'image_id',
      type: TypeEnum.IMAGE,
      name: 'cloud_image_id'
    })
  },
];

</script>

<template>
  <h3 class="info-title">实例信息</h3>
  <div class="warp-info">
    <detail-info :fields="cvmInfo" :detail="props.data"></detail-info>
  </div>
  <h3 class="info-title">网络信息-主网卡</h3>
  <div class="warp-info">
    <detail-info class="mt20" :fields="netInfo" :detail="props.data"></detail-info>
  </div>
  <h3 class="info-title">网络信息-扩展网卡</h3>
  <div class="warp-info">
    <detail-info class="mt20" :fields="expandNetInfo" :detail="props.data"></detail-info>
  </div>
  <h3 class="info-title">配置信息</h3>
  <div class="warp-info">
    <detail-info class="mt20" :fields="settingInfo" :detail="props.data"></detail-info>
  </div>
</template>

<style lang="scss" scoped>
  .info-title {
    font-size: 14px;
    margin: 20px 0 5px;
  }
  :deep(.host-info) .detail-info-main {
    height: auto !important;
  }
</style>
