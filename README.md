### watcher - Barely a packet sniffer
A small(really) go program that takes a protocol and starts listening and logging all the packets of that protocal.

Currently, it just logs the source of the packet and its size.

Note: It must be run with `sudo`.

### Options
1. `-proto` - The protocol whose packets are to be listened. Defaults to tcp. You can pass `icmp`, `igmp`, `tcp`, `udp`.
