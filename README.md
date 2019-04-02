# Listen
Simple MP3-only Music Player; built with Gotk3.

## Building

### Dependencies

- GTK 3.6-3.22
- GLib 2.36-2.40
- Cairo 1.10
- libasound2

requires at least go 1.8

```git clone https://github.com/oct2pus/listen; cd listen; go get -u; go build .```

## Feature Checklist for 1.0

- [x] displays embedded track art
- [x] pausing and resuming
- [x] displays unembedded album art
- [x] volume control
- [x] free move
- [x] restart song after completed
- [x] launch song from terminal
