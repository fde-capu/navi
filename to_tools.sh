make gomod
make
PLUGINDIR=~/.vim/after/plugin
cp -fluv --preserve=all ./bin/navi $TOOLZ_DIR/_
cp -fluv --preserve=all ./.navi_env $TOOLZ_DIR/
cp -fluv --preserve=all ./navi.vim $PLUGINDIR
echo "Don't froget to \`:source $PLUGINDIR/navi.vim\`."
