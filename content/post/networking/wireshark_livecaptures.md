---
title: "Live Capturing on Unusual Interfaces"
date: 2019-03-12T12:44:45Z
author: Ross Jacobs
desc: "Wireshark Bonus Topics"
tags:
  - networking
  - wireshark
  - commandfu
  - draft1
image: https://allabouttesting.org/wp-content/uploads/2018/06/tshark-count.jpg

draft: true
---

_WANTED: Suspicious individual trafficking in packets. Reward paid upon capture_

There are many possible non-traditional interfaces that Wireshark can capture
live on. Wireshark's extcaps are a means to do the same through a plugin system.

If you are using Windows, you will want to use [Windows Subsystem for
Linux](https://docs.microsoft.com/en-us/windows/wsl/install-win10) as Windows
has problems with pipes and FIFOs. Note that you may want to live read with
tshark as the Wireshark GUI may have problems on WSL.

## Browser Download

Some services provide live packet captures through a browser. This may offer
convenience, but you need to wait for the file to completely download to use it.
Alternatively, if you open the partially downloaded file in wireshark, you
interrupt the download.

To dynamically load a downloading file as a live capture, find the filename and
then run: 

	tail -f -n +1 <download partial> | wireshark -k -i -

If you would like wireshark to automatically start reading the downloading
partial capture, I created a [bash
script]() that will do
just that. If you want this script to autostart, add the script locally and then add
`/path/to/script &` to your `~/.bashrc`.

## Capturing remotely over an SSH connection

Getting a live capture over an ssh connection is a solved problem on all
platforms. `ssh` works for this purpose on Linux, Macos, and WSL on Windows
while
[`Plink`](https://kaischroed.wordpress.com/2013/01/28/howto-use-wireshark-over-ssh/)
works for Windows PuTTY users. Briefly, I'll go over what
that looks like for `ssh`.

### SSH remote capture options

_You can check that your ssh-key is loaded with `ssh-add -L`._

Initially, let's set up variables for cleaner code. Replace each variable in <>
with a value that works for you.

```bash
ssh_opts="<user>@<server> -p <port>"
remote_cmd="sudo /usr/sbin/tcpdump -s0 -n -w - not port <port>"
read_cmd="< 'wireshark -k' -OR- 'tshark' > -i"
```

We then have the option of piping directly:

```bash
ssh $ssh_opts $remote_cmd | $read_cmd - 
```

__Or__ using a named pipe:

```bash
mkfifo /tmp/capfifo
ssh $ssh_options $ssh_command > /tmp/capinfo &
$read_cmd /tmp/capfifo
```

## Scapy

Scapy is a versatile Python library for Packet Crafting. Let's say that you
wanted to write packets directly to Wireshark.

```bash
$ touch scapy.pcap
$ wireshark -k -i scapy.pcap
$ scapy

                     aSPY//YASa       
             apyyyyCY//////////YCa       |
            sY//////YSpcs  scpCY//Pp     | Welcome to Scapy
 ayp ayyyyyyySCP//Pp           syY//C    | Version 2.4.2
 AYAsAYYYYYYYY///Ps              cY//S   |
         pCCCCY//p          cSSps y//Y   | https://github.com/secdev/scapy
         SPPPP///a          pP///AC//Y   |
              A//A            cyP////C   | Have fun!
              p///Ac            sC///a   |
              P////YCpc           A//A   | Wanna support scapy? Rate it on
       scccccp///pSP///p          p//Y   | sectools!
      sY/////////y  caa           S//P   | http://sectools.org/tool/scapy/
       cayCyayP//Ya              pY/Ya   |             -- Satoshi Nakamoto
        sY/PsY////YCc          aC//Yp    |
         sc  sccaCY//PCypaapyCP//YSs  
                  spCPY//////YPSps    
                       ccaacs   
					   
>>> ping = IP(dst="8.8.8.8")/ICMP()
>>> wrpcap('scapy.pcap', 4*[ping])
```




Currently, you can write to a pcap

It looks like this is a limitation of `wrpcap` though.

## extcaps

### extcap context

The typical way to see packets live in Wireshark is to use some form of piping:

```bash
packet-source | wireshark -k -i - 

# -OR-

mkfifo myfifo
packet-source > myfifo & 
wireshark -k -i myfifo 
```

Note though that you have to recreate this command every time. The [extcap
interface](https://www.wireshark.org/docs/man-pages/extcap.html) gives you the
ability to present your source of packets as an interface that you can capture
on in *shark. Run `tshark -D` and note that you can capture to any interface
listed (including ones you create).

If you are successful, *shark will add your interface to its list:

```bash
bash $ tshark -D
1) eth0 (Your default interface)
.
.
9) MyCap0 (Antiquated Device Remote Capture) 

### builtin extcaps

In wireshark, these should be presented as
interfaces with a picture of a gear. In tshark, you can list which ones are
available with `tshark -D`. Note that `dumpcap -D` will show you all interfaces
EXCEPT extcap ones.

There are currently the utilities that are bundled with Wireshark.
You may have access to more or fewer depending on your system:

- androiddump
- ciscodump
- randpktdump
- sshdump
- udpdump

Sometimes using extcap utilities from the CLI can be unintuitive. 
Taking randpktdump as an example, let's figure out how to use it. 
`randpktdump --help` provides usage information. 

```bash
randpktdump --extcap-interfaces
extcap {version=0.1.0}{help=file:///usr/local/share/wireshark/randpktdump.html}
interface {value=randpkt}{display=Random packet generator}
```

### randpkt vs randpktdump

If `randpkt` is an option when you use `tshark -D`, then you can use it as an
extcap interface like so: 

    tshark -i randpkt -w extcap_example.pcap

__Or__, click on the "Random packet generator: randpkt" option when you first open
Wireshark. In both cases, you will get a 1000-packet pcap for a random protocol
with a 5000-byte limit (randpkt defaults).

* [ ] <using randpkt gif>

For point of comparison, `randpktdump` has the same functionality as the
Wireshark-builtin `randpkt` command, except with the advantage of leveraging an
extcap interface.

```bash
# These commands both produce 10 dns packets up to 1000 bytes long.
randpkt -b 1000 -c 10 -t dns /tmp/randpkt.pcap
randpktdump --extcap-interface=randpkt --maxbytes=1000 --count=10 --type=dns \
  --fifo=/tmp/randpkt2.pcap --capture
```

### Building your own extcap resources

The major impetus for extcap is to make YOUR non-traditional packet source
easire to work with. Note that many Wireshark core developers work for
Device Manufacturers or ISPs where automating captures from nontraditional
interfaces is important. 

Documentation on
[extcap](https://www.wireshark.org/docs/wsdg_html_chunked/ChCaptureExtcap.html),
utilities is a good resource for interface creation.  If you want to build your
own, these options are required for capture to function:

- --capture
- --extcap-capture-filter 
- --fifo

If you are trying to figure out where to start with your extcap project,
[`extcap_example.py`](https://github.com/wireshark/wireshark/blob/master/doc/extcap_example.py)
written by Robert Knall is worth looking at. Below is an example of how to use
this example utility.

<script id="asciicast-nt1WaIPrYEyrO1uxmnlnBbpvX" src="https://asciinema.org/a/nt1WaIPrYEyrO1uxmnlnBbpvX.js" async></script>

