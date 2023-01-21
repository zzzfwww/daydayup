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

## Theme Agnoster

* To use the agnoster theme you have to install some fonts. Otherwise you will see some question marks where symbols are supposed to go. Here are the steps to install:

* git clone https://github.com/powerline/fonts 
* download and unzip [zip](./fonts.zip)
* cd fonts
* ./install.sh from terminal / command line
* Open iTerm2->Preferences->Profiles->Change Font-> 12pt Meslo LG S DZ Regular for Powerline
* After doing that you should see this. It’s a pretty popular look out there:

```text
ZSH_THEME="agnoster"
```

--------

zsh [theme](https://travis.media/top-12-oh-my-zsh-themes-for-productive-developers/)
