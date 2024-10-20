# 介绍

由于没有好用的 [Headscale](https://headscale.net/) 控制器的后台管理中心，

随即使用 `VUE` `GO` 开发，并独立与 [Headscale](https://headscale.net/) 运行，不破坏其库表。

通过 ACL 可以实现多用户管理。

# 安装

1. 修改 `.env.production` 文件中 `VUE_APP_BASE_API` 的路由
2. `docker compose up -d` 启动
3. 创建 headscale 用户
```
docker exec -it headscale headscale users create admin
```
4. 创建密钥
```
docker exec -it headscale headscale apikeys create
```
5. 修改 `hub\config.yml` 中的 `headscale.api_key` 参数，然后重启 `headscale` 容器

# 截图

![user](https://github.com/user-attachments/assets/9f3feddd-b8b3-4a41-8c62-b04d8214ed63)
![node](https://github.com/user-attachments/assets/ee919cff-a8c5-43d6-952c-f85a39942ab7)
![key](https://github.com/user-attachments/assets/c4ddb38a-6c5c-4e4a-8ce3-0a556d387fb5)
![route](https://github.com/user-attachments/assets/e09612e2-071a-43fb-93b2-1231e818bf94)
![acls](https://github.com/user-attachments/assets/b26b478f-1d7f-4a60-8653-4c23d725dfe9)