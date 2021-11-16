# ptz-status
Go app to query a VISCAP-IP enabled PTZ camera and return Pan, Tilt, Zoom and Focus values.


## Please note

This is not a proper go repo (yet) just a dump of an example that compiles and works to query a camera.

### Why?

Because the std "cgi/http" interface on PTZ cameras, while reasonable for control, do not offer 'query' ability - i.e. you cannot probe "where is the camera pointing, what is it's focus" etc.

### To Build
1. Install golang. I used 1.17.3 on Ubuntu 20.04 but this is so low-level, any release will do.
2. Run `make`

### To use
1. `./ptz-status <camera_ip_address>`
2. The output is JSON and is consumed by other tooling I have in this space.

Use how you wish - copy, hack, I don't really mind.

It should be really easy to expand as needed.

My reference documentation included these files:

* https://cdn.shopify.com/s/files/1/0456/2701/5326/files/BA20X-x_User_Manual.pdf?v=1625557692
* https://ptzoptics.com/wp-content/uploads/2020/11/PTZOptics-VISCA-over-IP-Rev-1_2-8-20.pdf


### TODO
1. I'll add _writing_ PTZF at some point - I've no need for that right now as I use the cgi/http interface for that.
