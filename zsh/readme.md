# oh-my-zsh install

## curl 安装
* GitHub:
```shell
sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
```
* Gitee ( 国内镜像 )

```shell
sh -c "$(curl -fsSL https://gitee.com/mirrors/oh-my-zsh/raw/master/tools/install.sh)"
```
## wget 安装
* GitHub:
```shell
sh -c "$(wget https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"
```

* Gitee ( 国内镜像 )
```shell
sh -c "$(wget -O- https://gitee.com/pocmon/mirrors/raw/master/tools/install.sh)
```

## use plugin list
* `plugins=(git extract z sublime zsh-autosuggestions)` 

## plugin zsh-autosuggestions
* `oh-my-zsh plugin ‘zsh-autosuggestions’ not found`
```shell
git clone https://github.com/zsh-users/zsh-autosuggestions ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions
source ~/.zshrc
```
* slove it

