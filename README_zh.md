# gt

**中文|[English](https://github.com/bingfenglai/gt/blob/main/README.md)** 

## 一句话的简介

这是一个基于Golang实现的长短链接转换应用。
## 功能规划
1. 长短链接转换
2. 访问鉴权
3. 链接分组管理
4. 访问统计
## 开发框架
gin@v1.7.7

go-oauth2@v4.4.3

go-redis@v8.11.4

gorm@v1.22.5

jordan-wright/email@v4.0.1

mileusna/useragent@v1.0.2

patrickmn/go-cache@v2.1.0

swag@v0.21.1

## 关于swagger的使用

```go
// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags 用户信息   //swagger API分类标签, 同一个tag为一组
// @accept json  //浏览器可处理数据类型，浏览器默认发 Accept: */*
// @Produce  json  //设置返回数据的类型和编码
// @Param Authorization header string true  "访问令牌"
// @Param req body listModel true  "相关信息"
// @Param req formData listModel true  "相关信息"
// @Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
// @Param name query string false "name"
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @success 200 {object} jsonresult.JSONResult{data=proto.Order} "desc"
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /test/{id} [get]    //路由信息，一定要写上
```

