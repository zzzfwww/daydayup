# nginx 

## nginx处理跨域配置

* 前段页面通过axios请求后端，由于前段页面的地址和请求后台的地址不统一，所以出现跨域问题

* 跨域处理本质上需要header增加`Access-Control-Allow-Origin`,并且设置允许不同源访问

* 设置之后，前段会试探性的进行OPTHIONS请求后台，这时候也需要处理OPTHONS请求，直接返回200，要不然Nginx访问日志会报405错误，method not allow

* 最终Nginx配置文件如下：

```sh
[root@VM-20-16-centos nginx]# grep -Ev "#|^$" nginx.conf
```

```conf
user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;
include /usr/share/nginx/modules/*.conf;
events {
    worker_connections 1024;
}
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';
    access_log  /var/log/nginx/access.log  main;
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Headers *;
    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    keepalive_timeout   65;
    types_hash_max_size 4096;
    include             /etc/nginx/mime.types;
    default_type        application/octet-stream;
    include /etc/nginx/conf.d/*.conf;
    server {
        listen       80;
        listen       [::]:80;
        server_name  _;
        root         /usr/share/nginx/html;
        include /etc/nginx/default.d/*.conf;
        error_page 404 /404.html;
        location = /404.html {
        }
        error_page 500 502 503 504 /50x.html;
        location = /50x.html {
        }
	location ~ /chat {
          if ($request_method = 'OPTIONS') {
              return 200;
          }
           proxy_set_header Host $http_host;
           proxy_set_header X-Real-IP $remote_addr;
           proxy_set_header REMOTE-HOST $remote_addr;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
           proxy_pass http://127.0.0.1:8888;
      }
    }
}
```
