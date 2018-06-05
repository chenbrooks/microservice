# 小工具集
开发网站时，经常会用到一些第三方的库，为了有一个统一的地方来处理类似的需求，所以就打算把这些小工具全部都集成一个服务器，做成一个微信服务中心。

## 第三方集成todo list
- [x] 在线文件转存到AliOSS
- [ ] 微信文件转存到AliOSS

## 功能todo list
- [ ] 日志文件
- [ ] 权限设置

## 关于config.yaml
根目录上的config.yaml缺失,因为包含服务器用户名与密码的配置，所以没有放在Git里面，需要自己去创建，格式如下

```Bash
alioss:
  bucket_domain: oss-cn-shenzhen.aliyuncs.com
  bucket_name: your-own-bucket-name
  secret:
    key: your-own-aliyun-account-secret-key
    value: you-own-aliyun-account-secret-value
```
