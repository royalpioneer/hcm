### 描述

- 该接口提供版本：v1.0.0+。
- 该接口所需权限：账号删除。
- 该接口功能描述：删除指定账号（注：如果账号下在HCM还有资源存在，则无法删除）。

### 输入参数

| 参数名称       | 参数类型   | 必选  | 描述   |
|------------|--------|-----|------|
| account_id | string | 是   | 账号ID |

### 调用示例

```json

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
