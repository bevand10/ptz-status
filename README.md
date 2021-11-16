# ptz-status
Go app to query a PTZ camera and return Pan, Tilt, Zoom and Focus values.


## Please note

This is not a proper go repo (yet) just a dump of an example that compiles and works to query a camera.

### To Build
1. Install golang. I used 1.17.3 on Ubuntu 20.04 but this is so low-level, any release will do.
2. Run `make`

### To use
1. `./ptz-status <camera_ip_address>`
2. The output is JSON, just for show.

Use how you wish - copy, hack, I don't really mind.
