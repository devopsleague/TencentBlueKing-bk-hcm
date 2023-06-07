### 描述

- 该接口提供版本：v1.0.0+。
- 该接口所需权限：业务-IaaS资源操作。
- 该接口功能描述：更新VPC。

### URL

PATCH /api/v1/cloud/bizs/{bk_biz_id}/vpcs/{id}

### 输入参数

| 参数名称      | 参数类型     | 必选    | 描述     |
|-----------|----------|-------|--------|
| bk_biz_id | string   | 是     | 业务ID   |
| id        | string   | 是     | VPC的ID |
| memo      | string   | 否     | 备注     |

接口调用者可以根据以上参数自行根据更新场景设置更新的字段，除了ID之外的更新字段至少需要填写一个。

### 调用示例

```json
{
  "memo": "default vpc"
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