# MedicalCare

[TOC]

## Simple Intro

一个便民医护软件，主要目的是为了方便查询药价，并可以通过聊天室联系医师获取一些医疗建议。

Swagger API document address: [Swagger UI](http://139.9.221.94:8000/swagger/index.html)
Postman API document address: [Postman Public Document](https://documenter.getpostman.com/view/16949749/Uyxoh3ko)

## Project Structure

```
.
├── api
│   ├── common.go # response struct
│   └── v1 # the version 1.0 api
├── cache # cache functions and cron task
├── conf # contains the conf file
├── Dockerfile
├── docs # swagger docu folder
├── service
├── router
├── middleware
├── model
├── pkg 
│   ├── e # error message
│   ├── logging # log function
│   └── util # util function
├── go.mod # go module file
├── go.sum
└── main.go # the enterance of project
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

- Reminder Object: behaviour or medical eat.

### Medical Part

- Medical Info
- Local Price

### Map Part

- Distance minimum hospital



## Development Environment

- Golang 1.18
- PostgreSQL 11.13
- Redis 3.2.100

### Main Dependence

- github.com/gin-gonic/gin -> A MVC web framework
- gorm.io/gorm -> ORM for golang
- github.com/go-redis/redis/v8 -> Connect to Redis
- github.com/swaggo/gin-swagger -> Auto api docs generation
- github.com/gorilla/websocket -> Web Socket library
- golang.org/x/crypto/bcrypt -> Password encrypt
- github.com/dgrijalva/jwt-go -> JWT authority
- github.com/go-ini/ini -> ini Config file library

### Development Tools

- Goland 2021.3.3
- Docker 20.10.14
- Postman 9.19.3
