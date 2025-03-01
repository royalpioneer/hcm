### 描述

- 该接口提供版本：v1.0.0+
- 该接口所需权限：资源查看
- 该接口功能描述：查询 eip 列表(只包含通用字段)

### 请求参数
| 参数名称          | 参数类型                           | 必选 | 描述                                                         |
| ----------------- | ---------------------------------- | ---- | ------------------------------------------------------------ |
|page|Page|是|分页配置|
|filter|FilterExp|否|查询条件。不传时表示查询所有 eip |

#### Page
| 参数名称          | 参数类型                           | 必选 | 描述                                                         |
| ----------------- | ---------------------------------- | ---- | ------------------------------------------------------------ |
|count | bool | 是 | 是否返回总记录条数。 如果为true，查询结果返回总记录条数 count，但不返回查询结果详情数据 detail，此时 start 和 limit 参数将无效，且必需设置为0。如果为false，则根据 start 和 limit 参数，返回查询结果详情数据，但不返回总记录条数 count|
| limit | uint | 是 | 每页限制条数，最大500，不能为0 |
| start | uint | 否 | 记录开始位置，start 起始值为0 |
| sort	  | string	 | 否	  | 排序字段，返回数据将按该字段进行排序                                                                                                           |
| order	 | string	 | 否	  | 排序顺序（枚举值：ASC、DESC）                                                                                                                               |

#### FilterExp
| 参数名称          | 参数类型                           | 必选 | 描述                                                         |
| ----------------- | ---------------------------------- | ---- | ------------------------------------------------------------ |
|op | string | 是 |操作符（枚举值：and、or）。如果是and，则表示多个rule之间是且的关系；如果是or，则表示多个rule之间是或的关系|
|rules|Rule Array | 是 | 过滤规则，最多设置5个。如果 rules 为空数组，op（操作符）将没有作用，代表查询全部数据|

#### Rule[n]
| 参数名称          | 参数类型                           | 必选 | 描述                                                         |
| ----------------- | ---------------------------------- | ---- | ------------------------------------------------------------ |
| field | string | 是 | 查询条件 Field 名称，具体可使用的用于查询的字段及其说明请看下面 - 查询参数介绍 |
| op | string | 是 | 操作符（枚举值：eq、neq、gt、gte、le、lte、in、nin）|
| value | any | 是 | 查询条件 Value 值|

##### rule 表达式说明：

##### 1. 操作符

| 操作符 | 描述                                        | 操作符的value支持的数据类型                             |
|-----|-------------------------------------------|----------------------------------------------|
| eq  | 等于。不能为空字符串                                | boolean, numeric, string                     |
| neq | 不等。不能为空字符串                                | boolean, numeric, string                     |
| gt  | 大于                                        | numeric，时间类型为字符串（标准格式："2006-01-02 15:04:05"） |
| gte | 大于等于                                      | numeric，时间类型为字符串（标准格式："2006-01-02 15:04:05"） |
| lt  | 小于                                        | numeric，时间类型为字符串（标准格式："2006-01-02 15:04:05"） |
| lte | 小于等于                                      | numeric，时间类型为字符串（标准格式："2006-01-02 15:04:05"） |
| in  | 在给定的数组范围中。value数组中的元素最多设置100个，数组中至少有一个元素  | boolean, numeric, string                     |
| nin | 不在给定的数组范围中。value数组中的元素最多设置100个，数组中至少有一个元素 | boolean, numeric, string                     |
| cs  | 模糊查询，区分大小写                                | string                                       |
| cis | 模糊查询，不区分大小写                               | string                                       |

##### 2. 协议示例

查询 name 是 "Jim" 且 age 大于18小于30 且 servers 类型是 "api" 或者是 "web" 的数据。

```json
{
  "op": "and",
  "rules": [
    {
      "field": "name",
      "op": "eq",
      "value": "Jim"
    },
    {
      "field": "age",
      "op": "gt",
      "value": 18
    },
    {
      "field": "age",
      "op": "lt",
      "value": 30
    },
    {
      "field": "servers",
      "op": "in",
      "value": [
        "api",
        "web"
      ]
    }
  ]
}
```
#### 查询参数介绍：

| 参数名称         | 参数类型   | 描述                            |
|--------------|--------|-------------------------------|
| id           | string | 云盘 ID                         |
| vendor       | string | 云厂商                           |
| account_id   | string | 账号 ID                         |
| region | string | 区域 |
| public_ip | string | 公网 IP |
| bk_biz_id    | int64  | 业务ID，-1表示没有分配到业务      |
| creator      | string | 创建者                           |
| reviser      | string | 更新者                           |
| create_at    | string | 创建时间，标准格式：2006-01-02 15:04:05 |
| updated_at   | string | 更新时间，标准格式：2006-01-02 15:04:05 |

接口调用者可以根据以上参数自行根据查询场景设置查询规则。

### 调用示例
#### 请求参数示例
如查询云厂商是 tcloud 的 eip 列表
```json
{
    "page": {
        "limit": 10,
        "start": 0
    },
     "filter": {
        "op": "and",
        "rules": [
            {
                "field": "vendor",
                "op": "eq",
                "value": "tcloud"
            }
        ]
    }
}
```
#### 返回参数示例
```json
{
    "code": 0,
    "message": "",
    "data": {
        "details": [
            {
                "id": "0000000g",
                "vendor": "tcloud",
                "account_id": "abc",
                "name": "eip-test",
                "cloud_id": "eip-123123123",
                "bk_biz_id": -1,
                "region": "ap-guangzhou",
                "instance_id": "cvm-1123",
                "instance_type": "cvm",
                "public_ip": "9.128.0.34",
                "creator": "abc",
                "reviser": "abc",
                "created_at": "2023-02-14T11:42:24Z",
                "updated_at": "2023-02-14T11:42:24Z"
            },
            {
                "id": "0000000i",
                "vendor": "tcloud",
                "account_id": "abc",
                "cloud_id": "eip-003123123",
                "bk_biz_id": -1,
                "region": "ap-guangzhou",
                "public_ip": "9.128.0.35",
                "creator": "abc",
                "reviser": "abc",
                "created_at": "2023-02-14T11:42:38Z",
                "updated_at": "2023-02-14T11:42:38Z"
            }
        ]
    }
}
```
### 响应参数说明

| 参数名称    | 参数类型   | 描述   |
|---------|--------|------|
| code    | int  | 状态码  |
| message | string | 请求信息 |
| data    | Data | 响应数据 |
#### Data
| 参数名称   | 参数类型   | 描述                                       |
|--------|--------|------------------------------------------|
| count  | int | 当前规则能匹配到的总记录条数，当 limit > 0 时，才会返回，用于分页 |
| detail | Eip Array  | 查询返回的数据|

#### Eip[n]
| 参数名称         | 参数类型         | 描述                            |
|--------------|--------------|-------------------------------|
| id | string | 云盘 ID |
| vendor | string | 云厂商 |
| account_id | string | 云账号 ID |
| name | string | eip 名称. 该字段可能不返回，表示名称为 null |
| bk_biz_id | int | 分配给的cc 业务 ID， -1 表示未分配 |
| cloud_id | string | eip 在云厂商上的 ID |
| region | string | 地域 |
| instance_id | string | 绑定的实例 ID. 该字段如果不返回，表示未查询到绑定实例信息 |
| instance_type | string | 绑定的实例类型. 该字段如果不返回，表示未查询到绑定的实例资源类型 |
| public_ip | string | 公网 IP |
| creator | string | 创建者 |
| reviser | string | 更新者 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |