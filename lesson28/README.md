```
# 注册服务
curl --request PUT --data @config/health.json localhost:8500/v1/agent/service/register
# 注销服务
curl --request PUT localhost:8500/v1/agent/service/deregister/user
```