# Goal Piplin
[![codecov](https://codecov.io/gh/qbhy/goal-piplin/branch/master/graph/badge.svg)](https://codecov.io/gh/qbhy/goal-piplin)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbhy/goal-piplin)](https://goreportcard.com/report/github.com/qbhy/goal-piplin)
[![GoDoc](https://pkg.go.dev/badge/github.com/qbhy/goal-piplin?status.svg)](https://pkg.go.dev/github.com/qbhy/goal-piplin?tab=doc)
[![Join the chat at https://gitter.im/qbhy/goal-piplin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/qbhy/goal-piplin?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Sourcegraph](https://sourcegraph.com/github.com/qbhy/goal-piplin/-/badge.svg)](https://sourcegraph.com/github.com/qbhy/goal-piplin?badge)
[![Open Source Helpers](https://www.codetriage.com/qbhy/goal-piplin/badges/users.svg)](https://www.codetriage.com/qbhy/goal-piplin)
[![Release](https://img.shields.io/github/release/qbhy/goal-piplin.svg?style=flat-square)](https://github.com/qbhy/goal-piplin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/qbhy/goal-piplin)](https://www.tickgit.com/browse?repo=github.com/qbhy/goal-piplin)

## About Goal Piplin

Goal piplin is a very easy-to-use deployment tool.

- Easy to use
- Multi user
- Support grouping

## Install
Clone code

```bash
git clone https://github.com/qbhy/goal-piplin.git
```

Use docker compose to start the service (including mysql, redis, views and server)
```bash
docker compose up -d
```

Execute migration
```bash
docker compose exec server /var/www/piplin migrate
```

Initialize users and keys
```bash
docker compose exec server /var/www/piplin init
```
> The default account is piplin and the password is password

nginx configuration
```bash
cp nginx.conf /etc/nginx/site-enabled/piplin.conf
```
> In this step, you can modify the path and piplin.conf according to your needs

## update
If Goal-Piplin is updated, the latest image will be automatically packaged and pushed to the docker warehouse. At this time, you only need to execute the following command in your Goal-Piplin directory.
```bash
make update
```

## Views
goal piplin is a project that separates the front and back ends. The front end is developed based on antd pro, which uses umijs as a scaffolding.

 - [goal-piplin-views](https://github.com/qbhy/goal-piplin-views)
 - [ant-design-pro](https://github.com/ant-design/ant-design-pro)
 - [umijs](https://github.com/umijs/umi)

<img width="1466" alt="image" src="https://github.com/qbhy/goal-piplin/assets/24204533/d0e0c034-02d7-4eca-ad91-2f7090dd5c1c">

## Contributing

Thank you for considering contributing to the Goal Piplin! 

## License

The Goal Piplin is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
