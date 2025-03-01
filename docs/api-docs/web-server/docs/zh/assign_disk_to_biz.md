### 描述

- 该接口提供版本：v1.0.0+
- 该接口所需权限：资源分配
- 该接口功能描述：分配云盘到指定业务

#### 请求参数
| 参数名称          | 参数类型                           | 必选 | 描述                                                         |
| ----------------- | ---------------------------------- | ---- | ------------------------------------------------------------ |
|disk_ids | string array | 是 | 云盘 ID 数组|
| bk_biz_id | uint | 是 | cc 业务 ID |

### 调用示例
#### 请求参数示例
```json
{
    "disk_ids": ["00000002"],
    "bk_biz_id": 398
}
```

#### 返回参数示例
```json
{
    "code": 0,
    "message": "",
    "data": null
}
```
### 响应参数说明

| 参数名称    | 参数类型   | 描述   |
|---------|--------|------|
| code    | int  | 状态码  |
| message | string | 请求信息 |
| data    | null | 响应数据 |