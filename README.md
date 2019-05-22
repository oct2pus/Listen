# Listen
Simple MP3-only Music Player; built with Gotk3.

## Building

- GTK 3.6-3.22
- GLib 2.36-2.40
- Cairo 1.10
- libasound2
- librsvg2 (*build*)
- fish (*build*)
- meson (*build*)
- fish (*build*)
- debuild (*build*)

this project uses go modules, as such it requires go 1.11 minimum.

### .Deb

run ```debuild --no-sign``` in the project root directory. the deb file will be
in the directory above the project root directory.

### Meson
run ```meson build; cd build; ninja build; sudo ninja install; cd ..; 
sudo mv moe.jade.oct2pus.listen /usr/bin```

## Licence

This project is licenced under the GPL 3.0 or later except as noted,
the icon comes from [fxemoji](https://github.com/mozilla/fxemoji) and
is licenced under the MPLv2.

## Credits

- [Gotk3's](https://github.com/gotk3/gotk3) developers for their excellent (if incomplete) gtk bindings
for golang.
- [Heisantosh](https://github.com/heisantosh) for his
[useful packaging guide](https://github.com/heisantosh/howto-golang-gtk).
