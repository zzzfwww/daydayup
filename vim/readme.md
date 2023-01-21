# vim personal config

## pure nerdtree 
* use `vim.tar.gz`
* use `cd $HOME && tar -zxvf vim.tar.gz`
* `.vimrc` config detail
```vim
set rtp+=~/.vim/bundle/vundle/
call vundle#rc()
Bundle 'gmarik/vundle'
Bundle 'scrooloose/nerdcommenter'
Bundle 'scrooloose/nerdtree'
Bundle 'majutsushi/tagbar'
set autoread
colorscheme Tomorrow-Night-Eighties
nmap ;ll :NERDTreeToggle /data<CR>
nmap ;lw :NERDTree       /data<CR>
nmap ;ls :NERDTree       /data<CR>
nmap ;lm :NERDTree       /data<CR>
nmap ;li :NERDTree       /data<CR>
nmap ;ln :NERDTreeMirror<CR>
nmap ;vim :e ~/.vimrc<CR>
nmap ;tn :tabn<CR>
nmap ;tp :tabp<CR>
nmap ;tl :tabl<CR>
nmap ;tl :tabl<CR>
nmap ;te :tabe<CR>
nmap ;tc :tabc<CR>
nmap ;tf :tabfir<CR>
nmap ;bs :buffers<CR>
nmap ;bn :bn<CR>
nmap ;bp :bp<CR>
nmap ;bd :bd<CR>
nmap ;no :nohl<CR>
nmap ;sv :source ~/.vimrc<CR>
inoremap <C-k>             <C-X><C-k>
let NERDTreeDirArrows=0
```

## more pulgin

* use `vim_run.tar.gz`
