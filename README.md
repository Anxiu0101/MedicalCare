## Simple Intro

一个便民医护软件，主要目的是为了方便查询药价，并获取一些医疗建议。

## Main Dependence

```shell
# web structure
$ go get -u github.com/gin-gonic/gin

# database
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/mysql

# util
$ go get -u github.com/go-ini/ini # config file reader
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
