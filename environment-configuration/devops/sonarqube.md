# sonarqube

## 安装sonarqube
1. pull images 拉取镜像
```bash
docker pull postgres
docker pull sonarqube:8.9.6-community

[root@jenkins docker]# docker images
REPOSITORY                 TAG               IMAGE ID       CREATED          SIZE
mytest                     v2.0.1            0acb5da9428b   39 minutes ago   833MB
mytest                     v2.0.0            3e1b984404e3   42 minutes ago   833MB
mytest                     latest            ff48859718bb   43 minutes ago   833MB
postgres                   latest            4c6b3cc10e6b   3 days ago       379MB
sonarqube                  8.9.6-community   3f623568fa64   11 months ago    497MB
jenkins/jenkins            2.319.1-lts       2a4bbe50c40b   12 months ago    441MB
daocloud.io/library/java   8u40-jdk          4aefdb29fd43   7 years ago      816MB
```
2. 生成docker-compose yml文件

* vi /usr/local/docker/sonarqube_docker/docker-compose.yml
```yaml
version: '3.1'
services:
  db:
    image: postgres
    container_name: db
    ports:
      - 5432:5432
    networks:
      - sonarnet
    environment:
      POSTGRES_USER: sonar
      POSTGRES_PASSWORD: sonar
  sonarqube:
    image: sonarqube:8.9.6-community
    container_name: sonarqube
    depends_on:
      - db
    ports:
      - 9000:9000
    networks:
      - sonarnet
    environment:
      SONAR_JDBC_URL: jdbc:postgresql://db:5432/sonar
      SONAR_JDBC_USERNAME: sonar
      SONAR_JDBC_PASSWORD: sonar
networks:
  sonarnet:
    driver: bridge
```
* docker-compose up -d 
```bash
[root@jenkins sonarqube_docker]# docker-compose up -d
Creating network "sonarqube_docker_sonarnet" with driver "bridge"
Creating db ... done
Creating sonarqube ... done
```
3. 查看sonarqube启动日志
```bash
[root@jenkins sonarqube_docker]# docker logs -f sonarqube
2022.12.10 08:16:01 INFO  es[][o.e.t.TransportService] publish_address {127.0.0.1:40075}, bound_addresses {127.0.0.1:40075}
2022.12.10 08:16:01 INFO  es[][o.e.b.BootstrapChecks] explicitly enforcing bootstrap checks

ERROR: [1] bootstrap checks failed. You must address the points described in the following [1] lines before starting Elasticsearch.
bootstrap check failure [1] of [1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
ERROR: Elasticsearch did not exit normally - check the logs at /opt/sonarqube/logs/sonarqube.log
2022.12.10 08:16:01 INFO  es[][o.e.n.Node] stopping ...
2022.12.10 08:16:01 INFO  es[][o.e.n.Node] stopped
2022.12.10 08:16:01 INFO  es[][o.e.n.Node] closing ...
2022.12.10 08:16:01 INFO  es[][o.e.n.Node] closed
2022.12.10 08:16:01 WARN  app[][o.s.a.p.AbstractManagedProcess] Process exited with exit value [es]: 78
2022.12.10 08:16:01 INFO  app[][o.s.a.SchedulerImpl] Process[es] is stopped
2022.12.10 08:16:01 INFO  app[][o.s.a.SchedulerImpl] SonarQube is stopped
```

4. 修改虚拟内存大小
```bash
[root@jenkins sonarqube_docker]# vim /etc/sysctl.conf
[root@jenkins sonarqube_docker]# sysctl -p
vm.max_map_count = 262144
[root@jenkins sonarqube_docker]# docker-compose up -d
db is up-to-date
Starting sonarqube ... done

2022.12.10 08:28:31 INFO  ce[][o.s.c.c.CePluginRepository] Load plugins
2022.12.10 08:28:31 INFO  ce[][o.s.c.p.PluginInfo] Plugin [l10nzh] defines 'l10nen' as base plugin. This metadata can be removed from manifest of l10n plugins since version 5.2.
2022.12.10 08:28:33 INFO  ce[][o.s.c.c.ComputeEngineContainerImpl] Running Community edition
2022.12.10 08:28:33 INFO  ce[][o.s.ce.app.CeServer] Compute Engine is operational
2022.12.10 08:28:33 INFO  app[][o.s.a.SchedulerImpl] Process[ce] is up
2022.12.10 08:28:33 INFO  app[][o.s.a.SchedulerImpl] SonarQube is up
```

5. 登录sonarqube
```bash
初始用户名和密码都是admin
登录进去需要修改密码
http://192.168.3.102:9000/projects
```
6. maven settings.xml 配置sonarqube
```xml
<profile>
       <id>sonar</id>
       <activation>
        <activeByDefault>true</activeByDefault>
       </activation>
       <properties>
        <sonar.login>admin</sonar.login>
        <sonar.password>111111</sonar.password>
        <sonar.host.url>http://192.168.3.102:9000</sonar.host.url>
       </properties>
     </profile>
```
7. 本地springboot项目编译直接sonarqube扫码代码
```bash
mvn sonar:sonar
```
![sonar](./sonarqube.png)