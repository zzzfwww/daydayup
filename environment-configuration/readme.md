# 前端开发环境准备

1. 使用 vscode 和 vue 插件形式配置开发环境

2. vue 插件选择，搜索`vue volar`

3. nvm 安装

```bash
1. 安装nvm shell脚本
sudo curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
2. 安装完成之后默认会设置nvm环境变量
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
```

3. nvm 常用命令

```bash
nvm ls-remote  查看 Node 远程版本库
nvm install node  将安装最新版本的 Node
nvm install v12.7.0  将安装 12.7.0 版本的 Node
nvm uninstall v12.7.0  卸载 12.7.0 版本的 Node
nvm ls  查看已经安装的 Node 版本
nvm use v12.7.0 切换 12.7.0 为当前使用的版本
nvm alias default v12.7.0 将 12.7.0 设置为 Node 的默认版本
nvm which v12.7.0 查看 12.7.0 版本的 Node 的安装目录，比如：/Users/ccp/.nvm/versions/node/v12.7.0/bin/node
nvm --help  查看更多命令用法
```

4. 查看本地安装 node 版本

```bash
➜  ~ nvm ls
->      v17.9.1
        v19.1.0
         system
default -> stable (-> v19.1.0)
iojs -> N/A (default)
unstable -> N/A (default)
node -> stable (-> v19.1.0) (default)
stable -> 19.1 (-> v19.1.0) (default)
lts/* -> lts/hydrogen (-> N/A)
lts/argon -> v4.9.1 (-> N/A)
lts/boron -> v6.17.1 (-> N/A)
lts/carbon -> v8.17.0 (-> N/A)
lts/dubnium -> v10.24.1 (-> N/A)
lts/erbium -> v12.22.12 (-> N/A)
lts/fermium -> v14.21.1 (-> N/A)
lts/gallium -> v16.18.1 (-> N/A)
lts/hydrogen -> v18.12.1 (-> N/A)
```

# node 更新版本

```shell
1. 设置npm的镜像源为淘宝的镜像
npm config set registry https://registry.npm.taobao.org

2. 清空缓存以及安装版本
sudo npm cache clean -f
sudo npm install -g n
npm view node versions
3. 指定安装特定版本
➜  ~ sudo n 18.12.1
     copying : node/18.12.1
   installed : v18.12.1 (with npm 8.19.2)
```

# vue 前端常规操作

1. `npm install` 或者 `yarn`
2. 出错处理

- 错误

```error
getting "Error: EINVAL: invalid argument, read" for "npm install --save-dev eslint --verbose"
```

- 解决办法 `rm -rf node_modules package-lock.json` 之后再重新 `npm install`

3. 编译成功之后 `npm run serve` 运行前端程序

```shell
# 成功运行提示信息
  VITE v3.2.4  ready in 778 ms

  ➜  Local:   http://localhost:8080/
  ➜  Network: http://192.168.3.165:8080/
```

# vue项目报错踩坑
1. 报错` Expected indentation of 4 spaces but found 8`
```text
解决办法
 Expected indentation of 2 spaces but found 0
      indent: [
        "off", // off 掉就不会报错，但是有一个后果，不会格式化，先让clone下来的开源项目不报错再说
        2,
        2,
        {
          SwitchCase: 1
        }
      ],
```
2. 插件报错`eqeqeq `
```text
同样是插件修改配置
      eqeqeq: ['off','error', 'always', { null: 'ignore' }],
off 是解决!= 不报错的问题，要不软一直就报错，看着有点不爽
```

# MySQL 本地配置

- alias 命名 mysql1

```shell
alias mysql1="mysql --defaults-file=$HOME/.mysqlconf"
```

- mysqlconf 配置文件

```ini
➜  ~ cat .mysqlconf
[client]
port=3306
user=root
password="yourpassword"
```

# 代码统计工具使用

- 如果想要排除某个文件夹则使用如下命令

```shell
➜  web git:(1e9cdcd5) ✗ cloc --exclude-dir=node_modules ./
     317 text files.
     316 unique files.
      11 files ignored.

github.com/AlDanial/cloc v 1.90  T=0.26 s (1181.6 files/s, 182453.3 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
JSON                             3              0              0          30268
Vuejs Component                 55            598            185          10391
JavaScript                     186            189           1076           2294
Sass                             7            198             31           1498
SVG                              3              7              4            146
XML                              7              0              0            145
Markdown                         1              8              0             96
CSS                             42              8            123             80
HTML                             2              7              0             36
Dockerfile                       1              4              0             11
-------------------------------------------------------------------------------
SUM:                           307           1019           1419          44965
-------------------------------------------------------------------------------
```
