# brew系统自带安装网络非常慢
## 替换为阿里源
```bash
# 查看 brew.git 当前源
cd "$(brew --repo)" && git remote -v
origin    https://github.com/Homebrew/brew.git (fetch)
origin    https://github.com/Homebrew/brew.git (push)

# 查看 homebrew-core.git 当前源
cd "$(brew --repo homebrew/core)" && git remote -v
origin    https://github.com/Homebrew/homebrew-core.git (fetch)
origin    https://github.com/Homebrew/homebrew-core.git (push)

# 修改 brew.git 为阿里源
git -C "$(brew --repo)" remote set-url origin https://mirrors.aliyun.com/homebrew/brew.git

# 修改 homebrew-core.git 为阿里源
git -C "$(brew --repo homebrew/core)" remote set-url origin https://mirrors.aliyun.com/homebrew/homebrew-core.git

# zsh 替换 brew bintray 镜像
echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.aliyun.com/homebrew/homebrew-bottles' >> ~/.zshrc
source ~/.zshrc

# bash 替换 brew bintray 镜像
echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.aliyun.com/homebrew/homebrew-bottles' >> ~/.bash_profile
source ~/.bash_profile

# 刷新源
brew update
```

## 验证
```bash
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/libpng-1.6.39.monterey.bottle.tar.gz
######################################################################## 100.0%
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/fontconfig-2.14.1.monterey.bottle.tar.gz
######################################################################## 100.0%
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/glib-2.74.3.monterey.bottle.tar.gz
######################################################################## 100.0%
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/libx11-1.8.2.monterey.bottle.tar.gz
######################################################################## 100.0%
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/pixman-0.42.2.monterey.bottle.tar.gz
######################################################################## 100.0%
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/harfbuzz-5.3.1.monterey.bottle.1.tar.gz
######################################################################## 100.0%
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/jpeg-turbo-2.1.4.monterey.bottle.tar.gz
######################################################################## 100.0%
==> Downloading https://mirrors.aliyun.com/homebrew/homebrew-bottles/lz4-1.9.4.monterey.bottle.tar.gz
```