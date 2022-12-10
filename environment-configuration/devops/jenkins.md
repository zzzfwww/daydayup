
# Jenkins安装

## docker-compose启动Jenkins

编辑docker-compose yaml文件
```yaml
version: "3.1"
services:
  jenkins:
    image: jenkins/jenkins:2.319.1-lts
    container_name: jenkins
    ports:
      - 8080:8080
      - 50000:50000
    volumes:
       - ./data/:/var/jenkins_home/
```

## 报错信息查看
```bash
[root@jenkins jenkins_docker]# docker logs -f jenkins
touch: cannot touch '/var/jenkins_home/copy_reference_file.log': Permission denied
Can not write to /var/jenkins_home/copy_reference_file.log. Wrong volume permissions?

## 解决办法
chmod 777 -R data
```
再次重启docker-compose
```bash
[root@jenkins jenkins_docker]# docker-compose restart
```
正常启动的信息
```text
*************************************************************
*************************************************************
*************************************************************

Jenkins initial setup is required. An admin user has been created and a password generated.
Please use the following password to proceed to installation:

0608dd6f8e59411a927a4fe26a0c66d7

This may also be found at: /var/jenkins_home/secrets/initialAdminPassword

*************************************************************
*************************************************************
*************************************************************
```
##  如果Jenkins下载插件耗时太长

使用如下网址去找对应的插件，并且下载下来放到对应Jenkins插件目录
```url
https://plugins.jenkins.io/
```

## Jenkins配置简单项目
1. 配置git
![git](./jenkins/git.png)
2. 配置构建操作
![build](./jenkins/build.png)
3. 配置构建完成之后的操作
![build-after](./jenkins/build-after.png)
4. 参数化构建
![parameter](./jenkins/git-parameter.png)
![build-parameter](./jenkins/parameter-build.png)