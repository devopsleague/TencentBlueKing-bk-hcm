### 描述

- 该接口提供版本：v1.5.0+。
- 该接口所需权限：目标组删除。
- 该接口功能描述：业务下删除目标组。

### URL

DELETE /api/v1/cloud/bizs/{bk_biz_id}/target_groups/batch

### 输入参数

| 参数名称   | 参数类型       | 必选 | 描述        |
|-----------|--------------|------|------------|
| bk_biz_id | int          | 是   | 业务ID      |
| ids       | string array | 是   | 目标组ID数组 |

### 调用示例

```json
{
  "ids": [
    "00000001",
    "00000002"
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

| 参数名称 | 参数类型 | 描述    |
|---------|--------|---------|
| code    | int    | 状态码   |
| message | string | 请求信息 |