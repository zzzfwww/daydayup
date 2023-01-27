# Prometheus

## Prometheus server
* docker-compose.yml
```yaml
version: '3.7'
services:
   prometheus:
      image: prom/prometheus:latest
      container_name: prometheus
      restart: on-failure:1
      ports:
         - "9090:9090"
      volumes:
         - ./prometheus/:/etc/prometheus/
      deploy:
         resources:
            limits:
               memory: 1G
```
* docker-compose up -d 

## node exporter
```shell
docker run -d --name node-exporter -p 9010:9100 --restart=always \
-h "工具服务器" \
-v "/proc:/host/proc:ro" \
-v "/sys:/host/sys:ro" \
-v "/:/rootfs:ro" \
prom/node-exporter
```

## alter manager
* docker-compose.yml
```yaml
version: '3.7'
services:
   grafana:
      image: prom/alertmanager
      container_name: alertmanager
      restart: always
      ports:
         - "9093:9093"
      volumes:
         - ./alertmanager/:/etc/alertmanager/
      deploy:
         resources:
            limits:
               cpus: '0.50'
               memory: 2G
```