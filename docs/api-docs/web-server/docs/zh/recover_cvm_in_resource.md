### 描述

- 该接口提供版本：v1.0.0+。
- 该接口所需权限：回收站管理。
- 该接口功能描述：从回收站恢复虚拟机。

### 输入参数

| 参数名称 | 参数类型         | 必选  | 描述         |
|------|--------------|-----|------------|
| record_ids | string array | 是   | 回收记录ID |

### 调用示例

```json
{
  "record_ids": [
    "000000001"
  ]
}
```

### 响应示例

```json
{
  "code": 0,
  "message": "ok"
}
```

### 响应参数说明

| 参数名称    | 参数类型   | 描述   |
|---------|--------|------|
| code    | int32  | 状态码  |
| message | string | 请求信息 |
