# MedicalCare

[TOC]

## Simple Intro

一个便民医护软件，主要目的是为了方便查询药价，并获取一些医疗建议。

## Main Dependence

```shell
# web structure
$ go get -u github.com/gin-gonic/gin

# database: postgreSQL, redis
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/postgres
$ ggithub.com/go-redis/redis/v8

# doc
$ go get -u github.com/swaggo/swag/cmd/swag
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files

# websocket
$ go get -u github.com/gorilla/websocket

# util
$ go get -u github.com/go-ini/ini # config file reader
$ go get -u golang.org/x/crypto/bcrypt # transform Ciphertext
```

## Function

### User Part

- User Authority，including password，the qualification in medical
- User Info, including emails, telephone, gender, age and so on
- Medical Qualifications is in design.

### Blog Part

- Blog header, including title, intro, authority, avatars
- Blog body, the content
- Blog footer, including tag, Recommended medication.

### Reminder Part

- Reminder Object: behavior or medical eat.

### Medical Part

- Medical Info
- Local Price

### Map Part

- Distance minimum hospital



## Development Environment

- Goland 2021.3.3
- Golang 1.18
- MySQL 8.0.27
- 
