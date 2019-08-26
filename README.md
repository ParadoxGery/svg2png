# svg2png
commandline tool that converts svg to png

# install
## Download
[releases](https://github.com/ParadoxGery/svg2png/releases)

## Build
### Build Flags

`export CGO_LDFLAGS_ALLOW='-I/usr/include/librsvg-2.0|-I/usr/include/glib-2.0|-I/usr/lib/glib-2.0/include|-I/usr/lib/libffi-3.2.1/include|-I/usr/include/libmount|-I/usr/include/blkid|-I/usr/include/gdk-pixbuf-2.0|-I/usr/include/cairo|-I/usr/include/pixman-1|-I/usr/include/freetype2|-I/usr/include/libpng16|-I/usr/include/harfbuzz|-I/usr/include/uuid'`  
`export CGO_CFLAGS_ALLOW='-lrsvg-2|-lm|-lgio-2.0|-lgdk_pixbuf-2.0|-lgobject-2.0|-lglib-2.0|-lz|-lcairo'`
