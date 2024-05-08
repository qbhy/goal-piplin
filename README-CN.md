[English](README.md) | 中文
# Goal Piplin
[![codecov](https://codecov.io/gh/qbhy/goal-piplin/branch/master/graph/badge.svg)](https://codecov.io/gh/qbhy/goal-piplin)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbhy/goal-piplin)](https://goreportcard.com/report/github.com/qbhy/goal-piplin)
[![GoDoc](https://pkg.go.dev/badge/github.com/qbhy/goal-piplin?status.svg)](https://pkg.go.dev/github.com/qbhy/goal-piplin?tab=doc)
[![Join the chat at https://gitter.im/qbhy/goal-piplin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/qbhy/goal-piplin?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Sourcegraph](https://sourcegraph.com/github.com/qbhy/goal-piplin/-/badge.svg)](https://sourcegraph.com/github.com/qbhy/goal-piplin?badge)
[![Open Source Helpers](https://www.codetriage.com/qbhy/goal-piplin/badges/users.svg)](https://www.codetriage.com/qbhy/goal-piplin)
[![Release](https://img.shields.io/github/release/qbhy/goal-piplin.svg?style=flat-square)](https://github.com/qbhy/goal-piplin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/qbhy/goal-piplin)](https://www.tickgit.com/browse?repo=github.com/qbhy/goal-piplin)

## 关于 Goal-Piplin

Goal piplin 是一个非常好用的部署工具.

- 容易使用
- 支持多用户
- 支持分组
- 支持协作
- 一键复制项目
- 通过 CURL 部署

## 安装
克隆代码

```bash
git clone https://github.com/qbhy/goal-piplin.git
```

使用 docker compose 启动（包括 MySQL、Redis、server、views）
```bash
docker compose up -d
```

 执行迁移
```bash
docker compose exec server /var/www/piplin migrate
```

初始化用户和密钥
```bash
docker compose exec server /var/www/piplin init
```
> The default account is piplin and the password is password

nginx 配置
```bash
cp nginx.conf /etc/nginx/site-enabled/piplin.conf
```
> 这一步可以根据自己的需要修改 path 和 piplin.conf

## 更新
如果Goal-Piplin有更新，会自动打包最新的镜像推送到docker仓库。这时，你只需要在你的Goal-Piplin目录下执行以下命令即可。
```bash
make update
```

## 前端
goal piplin是一个前后端分离的项目。前端基于antd pro开发，使用umijs作为脚手架。

 - [goal-piplin-views](https://github.com/qbhy/goal-piplin-views)
 - [ant-design-pro](https://github.com/ant-design/ant-design-pro)
 - [umijs](https://github.com/umijs/umi)

<img width="1466" alt="image" src="https://github.com/qbhy/goal-piplin/assets/24204533/d0e0c034-02d7-4eca-ad91-2f7090dd5c1c">

## 贡献

感谢您考虑为 Goal Piplin 做出贡献！
您可以向此存储库提交 PR 或问题来参与该项目。
> 您也可以直接添加我的微信 qbhy0715 给我建议或者意见。

## 开源协议

Goal Poplin 是根据以下协议授权的开源软件 [MIT license](https://opensource.org/licenses/MIT).
