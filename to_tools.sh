make gomod
make
cp -fluv --preserve=all ./bin/navi $TOOLZ_DIR/_
cp -fluv --preserve=all ./.navi_env $TOOLZ_DIR/
cp -fluv --preserve=all ./navi.vim ~/.vim
echo "Don't froget to \`:source ~/.vim/navi.vim\`."
