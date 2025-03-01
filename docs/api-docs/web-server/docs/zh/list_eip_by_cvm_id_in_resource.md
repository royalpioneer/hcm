### 描述

- 该接口提供版本：v1.0.0+
- 该接口所需权限：资源查看
- 该接口功能描述：根据主机查询eip列表(带eip详情的扩展字段)

#### 路径参数说明
| 参数名称          | 参数类型                           | 必选 | 描述                                                         |
| ----------------- | ---------------------------------- | ---- | ------------------------------------------------------------ |
| vendor | string | 是 | 云厂商 |
| cvm_id | string | 是 | 主机 ID|

### 调用示例
#### 请求参数示例
如查询云厂商是 tcloud 的，主机 ID 是 00000050 的 eip 列表
#### 返回参数示例
```json
{
    "code": 0,
    "message": "",
    "data": [{
        "id": "0000000g",
        "account_id": "abc",
        "vendor": "tcloud",
        "name": "eip-test",
        "cloud_id": "eip-123123123",
        "bk_biz_id": 368,
        "region": "ap-guangzhou",
        "instance_id": "cvm-1123",
        "instance_type": "cvm",
        "public_ip": "9.128.0.34",
        "extension": {
            "bandwidth": 65535,
            "internet_charge_type": "BANDWIDTH_PACKAGE"
        },
        "creator": "abc",
        "reviser": "abc",
        "created_at": "2023-02-14T11:42:24Z",
        "updated_at": "2023-02-14T14:47:27Z",
        "rel_creator": "bk",
        "rel_created_at": "2023-02-24T19:33:55Z"
    }]
}
```
### 响应参数说明

| 参数名称    | 参数类型   | 描述   |
|---------|--------|------|
| code    | int  | 状态码  |
| message | string | 请求信息 |
| data    | Data | 响应数据 |
#### Data
eip 列表
| 参数名称   | 参数类型   | 描述                                       |
|--------|--------|------------------------------------------|
| id | string | Eip ID |
| vendor | string | 云厂商 |
| account_id | string | 云账号 ID |
| name | string | Eip 名称. 如果未返回该字段，表示为 null |
| bk_biz_id | int | 分配给的cc 业务 ID， -1 表示未分配 |
| cloud_id | string | Eip 在云厂商上的 ID |
| region | string | 地域 |
| public_ip | string | 公网 IP |
| instance_id | string | 绑定实例的 ID. 如果未返回该字段，表示未查询到绑定实例 ID |
| instance_type | string | 绑定实例的类型. 如果未返回该字段，表示未查询到绑定实例类型 |
| creator | string | 创建者 |
| reviser | string | 更新者 |
| created_at | string | 创建时间，标准格式：2006-01-02T15:04:05Z |
| updated_at | string | 更新时间，标准格式：2006-01-02T15:04:05Z |
| extension | EipExtension[vendor] | 各云厂商的差异化字段|
| cvm_id | string | 主机 ID |
| rel_creator | string | 绑定者 |
| rel_created_at | string | 绑定时间，标准格式：2006-01-02T15:04:05Z |

#### EipExtension[tcloud]

| 参数名称                           | 参数类型 |描述                                                         |
|--------------------------------| -------- |  ------------------------------------------------------------ |
| bandwidth | uint | 带宽 |
| internet_charge_type | string | 计费模式 |

#### EipExtension[azure]

| 参数名称                | 参数类型   | 描述                                                                       |
|---------------------|--------|--------------------------------------------------------------------------|
| ip_configuration_id | string | Resource ID (The IP configuration associated with the public IP address) |
| sku                 | string | sku                                                                      |
| sku_tier            | string | 层                                                                        |

#### EipExtension[huawei]
| 参数名称                           | 参数类型 | 描述                                                         |
|--------------------------------| -------- | ------------------------------------------------------------ |
| bandwidth_id | string | 带宽 ID |
| bandwidth_name | string | 带宽名 |
| bandwidth_size | uint | 带宽值 |
| port_id | HuaWeiDiskChargePrepaid |预付费参数。|

#### EipExtension[gcp]
| 参数名称 | 参数类型 | 描述 |
| address_type | string | 地址类型 |
| ip_version | string | IP 类型 |

#### EipExtension[aws]
| 参数名称 | 参数类型 | 描述 |
| public_ipv4_pool | string | 地址池 |
| domain | string | 范围 |
