

## Swagger生成错误 

> ParseComment error in file xxx :cannot find type definition: DeletedAt gorm.DeletedAt

因为没有使用这个外部依赖，默认不解析

解决
```bash
swag init --parseDependency --parseInternal
```