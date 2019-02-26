# 缴费管理系统服务端

外包某校财务处缴费管理系统的服务端。

需求简述：学生逐个收费，在收费同时打印收费单，并且将记录录入系统。当然也要包括记录查询等功能啦……

## 如何跑起来

1. 抓取项目

  ````
  go get github.com/mzz2017/VisualizationPlatform_service_service
  ````

2. 解决依赖

   ```
   cd $GOPATH/src/github.com/mzz2017/VisualizationPlatform_service_service
   dep ensure
   ```

3. 将`conf/app_reference.json`复制到`conf/app.json`，修改`app.json`中数据源配置

4. build并运行

   ```
   go build
   ./payment-management_service
   ```

   或者使用docker

   ```
   docker build -t payment-management .
   docker run -p 8127:8127 payment-management
   ```

   ​