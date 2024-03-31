cp ./bin/navi ~/_/toolz/_
mkdir -p ~/.vim/pack/my/start/navi/plugin
ln -f ./navi.vim ~/.vim/pack/my/start/navi/plugin/navi.vim
echo "Don't froget to re-source navi.vim"
