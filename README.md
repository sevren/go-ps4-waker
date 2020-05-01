# PS4-Waker

This tiny go application allows you to control your ps4 via a set of  CLI commands.

This application is intended to function just like the existing ps4-waker project in Python and in NodeJS

* Credits and references:
* https://github.com/hthiery/python-ps4
* https://github.com/dhleong/ps4-waker


# PS4 Communication

This application implements a number of features. It must be used on the same network as that of the PS4. 

## Discovery of PS4s

In order to discover what PS4s exist on the network we must setup a UDP server which will listen on all interfaces  on a random port and broadcast a simple message using DDP (Device Discovery Protocol)
All PS4s in the network will respond with thier status. 

To initate a search, the application will broadcast a packet to 
255.255.255.255:987 with the following string

`SRCH * HTTP/1.1\ndevice-discovery-protocol-version:00020020\n`

A sucessful response is read from the same connection used to send the broadcasted SRCH message. The ip address of the clients which respond to the broadcasted message is naturally the ipv4 address' belonging to the discovered PS4s

```
HTTP/1.1 620 Server Standby
host-id:BC60A7E3BF04
host-type:PS4
host-name:PS4-020
host-request-port:997
device-discovery-protocol-version:00020020
system-version:07508011
```

## Issuing commands to the PS4s

In order to issue commands to the PS4s you must send a specific DDP message to the ipv4 address of the PS4. Typically these messages are TCP based as opposed to the SRCH messages which require a UDP server.


