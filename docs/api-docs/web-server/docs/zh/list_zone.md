### 描述

- 该接口提供版本：v1.0.0+。
- 该接口所需权限：资源查看。
- 该接口功能描述：查询可用区列表。

#### 路径参数说明
| 参数名称          | 参数类型                           | 必选 | 描述                                                      |
| ----------------- | ------------------------------- | ---- | -------------------------------------------------------- |
| vendor | string | 是 | 云厂商 |
| region | string | 是 | 云地域 |                                             

### 输入参数

| 参数名称   | 参数类型         | 必选  | 描述                                            |
|--------|--------------|-----|-----------------------------------------------|
| filter | object       | 否   | 查询过滤条件                                        |
| page   | object       | 是   | 分页设置                                          |

#### filter

| 参数名称  | 参数类型        | 必选  | 描述                                                              |
|-------|-------------|-----|-----------------------------------------------------------------|
| op    | enum string | 是   | 操作符（枚举值：and、or）。如果是and，则表示多个rule之间是且的关系；如果是or，则表示多个rule之间是或的关系。 |
| rules | array       | 是   | 过滤规则，最多设置5个rules。如果rules为空数组，op（操作符）将没有作用，代表查询全部数据。             |

#### filter.rules[n] （详情请看 rules 表达式说明）

| 参数名称  | 参数类型        | 必选  | 描述                                          |
|-------|-------------|-----|---------------------------------------------|
| field | string      | 是   | 查询条件Field名称，具体可使用的用于查询的字段及其说明请看下面 - 查询参数介绍  |
| op    | enum string | 是   | 操作符（枚举值：eq、neq、gt、gte、le、lte、in、nin、cs、cis） |
| value | 可变类型        | 是   | 查询条件Value值                                  |

##### rules 表达式说明：

##### 1. 操作符

| 操作符 | 描述                                        | 操作符的value支持的数据类型                             |
|-----|-------------------------------------------|----------------------------------------------|
| eq  | 等于。不能为空字符串                                | boolean, numeric, string                     |
| neq | 不等。不能为空字符串                                | boolean, numeric, string                     |
| gt  | 大于                                        | numeric，时间类型为字符串（标准格式："2006-01-02T15:04:05Z"） |
| gte | 大于等于                                      | numeric，时间类型为字符串（标准格式："2006-01-02T15:04:05Z"） |
| lt  | 小于                                        | numeric，时间类型为字符串（标准格式："2006-01-02T15:04:05Z"） |
| lte | 小于等于                                      | numeric，时间类型为字符串（标准格式："2006-01-02T15:04:05Z"） |
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

#### page

| 参数名称   | 参数类型    | 必选  | 描述                                                                                                                                               |
|--------|---------|-----|--------------------------------------------------------------------------------------------------------------------------------------------------|
| count	 | bool	   | 是	  | 是否返回总记录条数。 如果为true，查询结果返回总记录条数 count，但不返回查询结果详情数据 detail，此时 start 和 limit 参数将无效，且必需设置为0。如果为false，则根据 start 和 limit 参数，返回查询结果详情数据，但不返回总记录条数 count |
| start	 | uint32	 | 否	  | 记录开始位置，start 起始值为0                                                                                                                               |
| limit	 | uint32	 | 否	  | 每页限制条数，最大500，不能为0                                                                                                                                |
| sort	  | string	 | 否	  | 排序字段，返回数据将按该字段进行排序                                                                                                                               |
| order	 | string	 | 否	  | 排序顺序（枚举值：ASC、DESC）                                                                                                                               |

### 调用和响应 示例

### tcloud req 
```json 
{
  "filter": {
    "op": "and",
    "rules": [
      {
        "field": "state",
        "op": "eq",
        "value": "AVAILABLE"
      }
    ]
  },
  "page": {
    "count": false,
    "start": 0,
    "limit": 500
  }
}
```

### tcloud resp
```json 
{
    "code": 0,
    "message": "",
    "data": {
        "details": [
            {
                "id": "000001rn",
                "vendor": "tcloud",
                "cloud_id": "100006",
                "name": "ap-guangzhou-6",
                "name_cn": "广州六区",
                "region": "ap-guangzhou",
                "state": "AVAILABLE",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-23T20:18:38Z",
                "updated_at": "2023-03-23T20:18:38Z"
            },
            {
                "id": "000001ro",
                "vendor": "tcloud",
                "cloud_id": "100007",
                "name": "ap-guangzhou-7",
                "name_cn": "广州七区",
                "region": "ap-guangzhou",
                "state": "AVAILABLE",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-23T20:18:38Z",
                "updated_at": "2023-03-23T20:18:38Z"
            },
            {
                "id": "000001rr",
                "vendor": "tcloud",
                "cloud_id": "100003",
                "name": "ap-guangzhou-3",
                "name_cn": "广州三区",
                "region": "ap-guangzhou",
                "state": "AVAILABLE",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-23T20:18:38Z",
                "updated_at": "2023-03-23T20:18:38Z"
            },
            {
                "id": "000001rs",
                "vendor": "tcloud",
                "cloud_id": "100004",
                "name": "ap-guangzhou-4",
                "name_cn": "广州四区",
                "region": "ap-guangzhou",
                "state": "AVAILABLE",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-23T20:18:38Z",
                "updated_at": "2023-03-23T20:18:38Z"
            }
        ]
    }
}
```

### aws req (huawei、gcp类似) filter可以不填
```json 
{
  "page": {
    "count": false,
    "start": 0,
    "limit": 500
  }
}
```

### aws resp (huawei、gcp类似) filter可以不填
```json 
{
    "code": 0,
    "message": "",
    "data": {
        "details": [
            {
                "id": "000001rl",
                "vendor": "aws",
                "cloud_id": "usw2-az1",
                "name": "us-west-2b",
                "name_cn": "",
                "region": "us-west-2",
                "state": "available",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-23T19:51:00Z",
                "updated_at": "2023-03-23T19:51:00Z"
            }
        ]
    }
}
```

# azure 不可使用filter（所有资源可用区都是固定的三个 "1"、 "2" 、"3"）
