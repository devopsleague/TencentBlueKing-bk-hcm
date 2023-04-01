### 描述

- 该接口提供版本：v1.0.0+。
- 该接口所需权限：资源查看。
- 该接口功能描述：查询硬盘与主机的关联关系，且带disk信息。

### 输入参数

| 参数名称              | 参数类型   | 必选  | 描述            |
|-------------------|--------|-----|---------------|
| filter            | object | 是   | 查询过滤条件        |
| page              | object | 是   | 分页设置          |

#### filter

| 参数名称  | 参数类型        | 必选  | 描述                                                              |
|-------|-------------|-----|-----------------------------------------------------------------|
| op    | enum string | 是   | 操作符（枚举值：and、or）。如果是and，则表示多个rule之间是且的关系；如果是or，则表示多个rule之间是或的关系。 |
| rules | array       | 是   | 过滤规则，最多设置5个rules。如果rules为空数组，op（操作符）将没有作用，代表查询全部数据。             |

#### rules[n] （详情请看 rules 表达式说明）

| 参数名称  | 参数类型        | 必选  | 描述                                         |
|-------|-------------|-----|--------------------------------------------|
| field | string      | 是   | 查询条件Field名称，具体可使用的用于查询的字段及其说明请看下面 - 查询参数介绍 |
| op    | enum string | 是   | 操作符（枚举值：eq、neq、gt、gte、le、lte、in、nin、cs、cis）       |
| value | 可变类型        | 是   | 查询条件Value值                                 |

##### rules 表达式说明：

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

#### page

| 参数名称  | 参数类型   | 必选  | 描述                                                                                                                                                  |
|-------|--------|-----|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| count | bool   | 是   | 是否返回总记录条数。 如果为true，查询结果返回总记录条数 count，但查询结果详情数据 details 为空数组，此时 start 和 limit 参数将无效，且必需设置为0。如果为false，则根据 start 和 limit 参数，返回查询结果详情数据，但总记录条数 count 为0 |
| start | uint32 | 否   | 记录开始位置，start 起始值为0                                                                                                                                  |
| limit | uint32 | 否   | 每页限制条数，最大500，不能为0                                                                                                                                   |
| sort  | string | 否   | 排序字段，返回数据将按该字段进行排序                                                                                                                                  |
| order | string | 否   | 排序顺序（枚举值：ASC、DESC）                                                                                                                                  |

#### 查询参数介绍：

| 参数名称                | 参数类型   | 描述                                   |
|---------------------|--------|--------------------------------------|
| cvm_id              | string | 云主机ID                                |
| cloud_id            | string | 云资源ID                                |
| name                | string | 名称                                   |
| vendor              | string | 供应商（枚举值：tcloud、aws、azure、gcp、huawei） |
| bk_biz_id           | int64  | 业务ID                                 |
| bk_cloud_id         | int64  | 云区域ID                                |
| account_id          | string | 账号ID                                 |
| region              | string | 地域                                   |
| zone                | string | 可用区                                  |
| cloud_image_id      | string | 云镜像ID                                |
| os_name             | string | 操作系统名称                               |
| memo                | string | 备注                                   |
| status              | string | 状态                                   |
| machine_type        | string | 设备类型                                 |
| cloud_created_time  | string | Cvm在云上创建时间                           |
| cloud_launched_time | string | Cvm启动时间                              |
| cloud_expired_time  | string | Cvm过期时间                              |
| creator             | string | 创建者                                  |
| reviser             | string | 修改者                                  |
| created_at          | string | 创建时间                                 |
| updated_at          | string | 修改时间                                 |

接口调用者可以根据以上参数自行根据查询场景设置查询规则。

### 调用示例

#### 获取详细信息请求参数示例

查询创建者是Jim，且与硬盘 '00000001' 没有绑定关系的关联关系列表且带有Cvm信息。

```json
{
  "filter": {
    "op": "and",
    "rules": [
      {
        "field": "created_at",
        "op": "eq",
        "value": "Jim"
      }
    ]
  },
  "page": {
    "count": false,
    "start": 0,
    "limit": 500
  },
  "not_equal_disk_id": "00000001"
}
```

#### 获取数量请求参数示例

查询创建者是Jim的Cvm数量。

```json
{
  "filter": {
    "op": "and",
    "rules": [
      {
        "field": "created_at",
        "op": "eq",
        "value": "Jim"
      }
    ]
  },
  "page": {
    "count": true
  }
}
```

### 响应示例

#### 获取详细信息返回结果示例

```json
{
    "code": 0,
    "message": "",
    "data": {
        "count": 0,
        "details": [
            {
                "id": "00000447",
                "vendor": "gcp",
                "account_id": "0000002d",
                "cloud_id": "8981356937866689091",
                "bk_biz_id": -1,
                "name": "guohu-login-login4-0001-1",
                "region": "us-central1",
                "zone": "us-central1-a",
                "disk_size": 50,
                "disk_type": "https://www.googleapis.com/compute/v1/projects/tencentgcpieg6/zones/us-central1-a/diskTypes/pd-balanced",
                "status": "READY",
                "recycle_status": "",
                "is_system_disk": null,
                "memo": "",
                "extension": "",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-30T16:45:35Z",
                "updated_at": "2023-03-30T16:45:35Z",
                "rel_creator": "",
                "rel_created_at": null
            },
            {
                "id": "00000449",
                "vendor": "gcp",
                "account_id": "0000002d",
                "cloud_id": "3383952457308410923",
                "bk_biz_id": -1,
                "name": "disk-3",
                "region": "us-central1",
                "zone": "us-central1-a",
                "disk_size": 100,
                "disk_type": "https://www.googleapis.com/compute/v1/projects/tencentgcpieg6/zones/us-central1-a/diskTypes/pd-balanced",
                "status": "READY",
                "recycle_status": "",
                "is_system_disk": null,
                "memo": "",
                "extension": "",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-30T16:45:35Z",
                "updated_at": "2023-03-30T16:45:35Z",
                "rel_creator": "",
                "rel_created_at": null
            },
            {
                "id": "0000044e",
                "vendor": "gcp",
                "account_id": "0000002d",
                "cloud_id": "4423279928137400016",
                "bk_biz_id": -1,
                "name": "guohu-batch-create-0036",
                "region": "us-central1",
                "zone": "us-central1-a",
                "disk_size": 50,
                "disk_type": "https://www.googleapis.com/compute/v1/projects/tencentgcpieg6/zones/us-central1-a/diskTypes/pd-balanced",
                "status": "READY",
                "recycle_status": "",
                "is_system_disk": null,
                "memo": "",
                "extension": "",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-30T16:45:35Z",
                "updated_at": "2023-03-30T16:45:35Z",
                "rel_creator": "",
                "rel_created_at": null
            },
            {
                "id": "0000044g",
                "vendor": "gcp",
                "account_id": "0000002d",
                "cloud_id": "4809607271472925776",
                "bk_biz_id": -1,
                "name": "disk-2",
                "region": "us-central1",
                "zone": "us-central1-a",
                "disk_size": 100,
                "disk_type": "https://www.googleapis.com/compute/v1/projects/tencentgcpieg6/zones/us-central1-a/diskTypes/pd-balanced",
                "status": "READY",
                "recycle_status": "",
                "is_system_disk": null,
                "memo": "",
                "extension": "",
                "creator": "guohuliu",
                "reviser": "guohuliu",
                "created_at": "2023-03-30T16:45:35Z",
                "updated_at": "2023-03-30T16:45:35Z",
                "rel_creator": "",
                "rel_created_at": null
            }
        ]
    }
}
```

#### 获取数量返回结果示例

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "count": 1
  }
}
```