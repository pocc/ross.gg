---
title: "Wireshark CLI"
date: 2019-03-12T12:44:45Z
author: Ross Jacobs
desc: "Using the Wireshark CLI for Packet Analysis"
tags:
  - networking
  - wireshark
image: https://allabouttesting.org/wp-content/uploads/2018/06/tshark-count.jpg

draft: true
---

_Packet Analysis, Scripted_

In line with the Unix philosophy of "Do one thing well", Wireshark has many
small CLI utilities. If you are reading this article because you want to know
how to use to do X with the CLI, you've come to the right place.

<!-- Kludgy TOC until I can figure out how to include {{ hugo tags }} in the content -->
## Table of Contents

* [Info](#info)
	* [capinfos](#capinfos)
	* [rawshark](#capinfos)
* [Edit](#edit) 
	* [mergecap](#mergecap)
	* [reordercap](#reordercap)
* [Generate](#generate)
	* [randpkt](#randpkt) 
* [Capture](#capture)
	* [dumpcap](#dumpcap) 
	* [tshark](#tshark)
* [Additional Topics](#additional-topics)
	* [Editing Hex](#editing-hex)
	* [Piping](#piping) 
	* [Capturing over ssh](#ssh-capture)

## <a name="info"></a>Info

Read a packet capture and print data about it. While it is possible to specify
the  
these utilities, I recommend specifying all options and then munging data in
your $language.

### <a name="info"></a>capinfos

capinfos gets metadata about a packet capture. You can be very granular about
what pieces of data you want displayed and the output format. For example, 

It's fairly straightforward to parse this into a hashtable in your $language.
For an example, check out get_capinfos() in my [wsutils gist](https://gist.github.com/pocc/2c89dd92d6a64abca3db2a29a11f1404).

### <a name=rawshark></a>rawshark

rawshark is a utility that takes an input stream and parses it. It is low-level
and provides options you would expect to see if you were working
with the source code. 

#### Reasons not to use rawshark

![Not recommended](https://media2.giphy.com/media/d31vYmpaCrKs9Z6w/giphy.gif)

- You MUST specify the [tcpdump link-layer header
  type](https://www.tcpdump.org/linktypes.html) or protocol name before any
  others (and sometimes it isn't clear [which
  one](https://stackoverflow.com/questions/14092321/rawshark-output-format-for-802-11-and-radiotap-headers)
  you should use)
- You MUST send in an input stream because it cannot parse files
- You MUST send in raw packets without the header. rawshark only knows how to
  remove a pcap-type header before processing and errors out on any other
  capture file. 
- If piping to text-processing tools like awk, needless text cruft is added
  pertaining to the c-style struct of the packets. 
- `rawshark` has this generic error for any parsing problem:
   
    ```bash
    rawshark: The standard input appears to be damaged or corrupt.
    (Bad packet length: 693250156
    )
    ```

#### Replace with tshark

But the reason you should avoid using it because tshark can do everything it can
do, and better. To transition, rawshark's options `-nNrR` are the same as
tshark's, and all of the others can be discarded.

#### Attempts to use rawshark (skip to [next section](#edit))

In this example, I am using the
[`dhcp.pcap`](https://wiki.wireshark.org/SampleCaptures#General_.2F_Unsorted)
from wireshark's
[SampleCaptures](https://wiki.wireshark.org/SampleCaptures#General_.2F_Unsorted).

1. So rawshark will not take tshark raw output...

	```bash
    $ tshark -r dhcp.pcap -w - | rawshark -s -r - -d proto:udp -F udp.port
	
    0 FT_UINT16 BASE_PT_UDP - 
	rawshark: The standard input appears to be damaged or corrupt.
	(Bad packet length: 673213298
	)
	```
	
2. You would think that specifying `proto` of udp for DHCP would work, but it
  shows incorrect output. DHCP uses UDP ports 67 and 68:

    ```bash
	$ cat dhcp.pcap | rawshark -s -r - -d proto:udp -F udp.port
	
	0 FT_UINT16 BASE_PT_UDP - 1 FT_UINT16 BASE_PT_UDP - 
	1 1="65535" 0="65535" -
	2 1="11" 0="33281" -
	3 1="65535" 0="65535" -
	4 1="11" 0="33281" -
	```
3. Finally, by specifying encap type instead of proto, we get useful output.

	```bash
	$ cat dhcp.pcap | rawshark -s -r - -d encap:1 -F udp.port
	
	FT_UINT16 BASE_PT_UDP - 1 FT_UINT16 BASE_PT_UDP - 
	1 1="68" 0="67" -
	2 1="67" 0="68" -
	3 1="68" 0="67" -
	4 1="67" 0="68" -
	```

4. `tshark` is more useful with less work though, even if we pass in as a stream
	(the supposed purpose of `rawshark`:
	
	```bash
	$ cat dhcp.pcap | tshark -r -
	
	1   0.000000      0.0.0.0 → 255.255.255.255 DHCP 314 DHCP Discover - Transaction ID 0x3d1d
    2   0.000295  192.168.0.1 → 192.168.0.10 DHCP 342 DHCP Offer    - Transaction ID 0x3d1d
    3   0.070031      0.0.0.0 → 255.255.255.255 DHCP 314 DHCP Request  - Transaction ID 0x3d1e
    4   0.070345  192.168.0.1 → 192.168.0.10 DHCP 342 DHCP ACK      - Transaction ID 0x3d1e
	```
	
	tshark has the advantage of being able to read files too: `tshark -r dhcp.pcap`.
	
	

## <a name=edit></a>Edit

### <a name=reordercap></a>reordercap

Sometimes packets are out of order. Reordercap fixes that.

### <a name=mergecap></a>mergecap

Merge two or more packet captures together

### <a name=text2pcap></a>text2pcap
Convert a hexstring into a packet capture

### <a name="editcap"></a>editcap
Edit the attributes of a 

_Given an existing pcap, generate a pcap_

 
## <a name=generate></a>Generate

### <a name=randpkt></a>randpkt
Generate packets

## <a name=capture></a>Capture

Question: In what significant ways do dumpcap and tshark differ?

### <a name="dumpcap"></a>dumpcap

Utility that other Wireshark utilities use to capture packets 

### <a name="tshark"></a>tshark

This article does not cover tshark for reasons of brevity. Continue the journey
at the [tshark page]()

## <a name=additional-topics></a>Additional Topics

### <a name=editing-hex></a>Editing Hex

### <a name=piping></a>Piping 

Piping is important to using many of these utilities. For example, it is not
really possible to use rawshark without piping as it expects a FIFO or stream. 

### <a name=ssh-capture></a>Capturing over ssh

Let's say that you have a remote computer and want to monitor it remotely.

## <a name=closing-thoughts></a>Closing Thoughts

Personally, I think that wireshark's CLI needs a better API. For example, git
has a large amount of functionality, but.

## Further Reading

* [Official Docs](https://www.wireshark.org/docs/man-pages/)
