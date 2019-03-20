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

<!-- Kludgy TOC until I can figure out how to include {{ hugo toc }} in the content -->
## Table of Contents

* [Info](#info)
	* [capinfos](#capinfos)
	* [rawshark](#rawshark)
* [Edit](#edit) 
	* [editcap](#editpcap)
	* [mergecap](#mergecap)
	* [reordercap](#reordercap)
	* [text2pcap](#text2pcap)
* [Generate](#generate)
	* [randpkt](#randpkt) 
* [Capture](#capture)
	* [dumpcap](#dumpcap) 
	* [tshark](#tshark)
* [Additional Topics](#additional-topics)
	* [Editing Hex](#editing-hex)
	* [Piping](#piping) 
	* [Capturing over ssh](#ssh-capture)

# <a name="info"></a>Info

_Read a packet capture and print data about it._

## <a name="capinfos"></a>capinfos

capinfos gets metadata about a packet capture. You can be very granular about
what pieces of data you want displayed and the output format. 

### List Data

To list the data (the default), use `capinfos <file>`:

```bash
$ capinfos dhcp.pcap

File name:           dhcp.pcap
File type:           Wireshark/tcpdump/... - pcap
File encapsulation:  Ethernet
File timestamp precision:  microseconds (6)
Packet size limit:   file hdr: 65535 bytes
Number of packets:   4
File size:           1,400 bytes
Data size:           1,312 bytes
Capture duration:    0.070345 seconds
First packet time:   2004-12-05 19:16:24.317453
Last packet time:    2004-12-05 19:16:24.387798
Data byte rate:      18 kBps
Data bit rate:       149 kbps
Average packet size: 328.00 bytes
Average packet rate: 56 packets/s
SHA256:              2471b5420bdac826eecf8f61a2bbb4a3eb20dbfab7c02ff2be502f349f368214
RIPEMD160:           43f96835c4501ccdbf53dea46f711a3b8c7f4cff
SHA1:                07583b66a5b12a6b557cef4ed38d7a0c77968f68
Strict time order:   True
Number of interfaces in file: 1
Interface #0 info:
                     Encapsulation = Ethernet (1 - ether)
                     Capture length = 65535
                     Time precision = microseconds (6)
                     Time ticks per second = 1000000
                     Number of stat entries = 0
                     Number of packets = 4
```

This format is useful for getting _all_ the metadata from a pcap.

### Tabular Data

`capinfos -T dhcp.pcap` provides the same info in a tab-delimited table. 

```bash
$ capinfos -T dhcp.pcap

File name	File type	File encapsulation	File time precision	Packet size limit	Packet size limit min (inferred)	Packet size limit max (inferred)	Number of packets	File size (bytes)	Data size (bytes)	Capture duration (seconds)	Start time	End time	Data byte rate (bytes/sec)	Data bit rate (bits/sec)	Average packet size (bytes)	Average packet rate (packets/sec)	SHA256	RIPEMD160	SHA1	Strict time order	Capture hardware	Capture oper-sys	Capture application	Capture comment
dhcp.pcap	Wireshark/tcpdump/... - pcap	Ethernet	microseconds	65535	n/a	n/a	4	1400	1312	0.070345	2004-12-05 19:16:24.317453	2004-12-05 19:16:24.387798	18650.89	149207.13	328.00	56.86	2471b5420bdac826eecf8f61a2bbb4a3eb20dbfab7c02ff2be502f349f368214	43f96835c4501ccdbf53dea46f711a3b8c7f4cff	07583b66a5b12a6b557cef4ed38d7a0c77968f68	True
```

What better way to present tabular options than a table‽

| option | description                    |
|--------|--------------------------------|
| `-r`   | No header                      |
| `-m`   | comma-delimited                |
| `-b`   | space-delimited                |
| `-q`   | quote infos with single quotes |
| `-Q`   | quote infos with double quotes |

_Note that interface information is stripped in this format._

### Recommendations

`capinfos` offers 22 options `-acdDeEFHiIkKlnosStuxyz` to print specific
elements. My perspective is that it is better to use a scripting language to
convert all of the infos into a reusable format.  It's fairly straightforward to
parse `capinfos <file>` into a hashtable in your $language. For an example in
Python, check out get_capinfos() in my [wsutils
gist](https://gist.github.com/pocc/2c89dd92d6a64abca3db2a29a11f1404).

## <a name=rawshark></a>rawshark

rawshark is a utility that takes an input stream and parses it. It is low-level
and provides options you would expect to see if you were working
with the source code. 

<div>
<img src="https://media2.giphy.com/media/d31vYmpaCrKs9Z6w/giphy.gif" alt="Not Recommended"><i>&nbsp;&nbsp;What using rawshark feels like</i></img>
<p></p></div>

### Reasons not to use rawshark

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

### You should use tshark instead

But the reason you should avoid using it because tshark can do everything it can
do, and better. To transition, rawshark's options `-nNrR` are the same as
tshark's, and all of the others can be discarded.

### If you must... (skip to [next section](#edit))

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
	
# <a name=edit></a>Edit

_Edit packet captures for fun and profit._

## <a name="editcap"></a>editcap

Edit the attributes of a 

## <a name=mergecap></a>mergecap

Merge two or more packet captures together

## <a name=reordercap></a>reordercap

Sometimes packets are out of order. Reordercap fixes that.

## <a name=text2pcap></a>text2pcap
Convert a hexstring into a packet capture


_Given an existing pcap, generate a pcap_

 
# <a name=generate></a>Generate

_Make traffic that didn't exist before._ 

## <a name=randpkt></a>randpkt
Generate packets

## Comparisons with other tools
* Scapy
* Ostinato

# <a name=capture></a>Capture

> _Everything comes to us that belongs to us if we create the capacity to receive it._ 

_-Rabindranath Tagore_

Question: In what significant ways do dumpcap and tshark differ?

## <a name="dumpcap"></a>dumpcap

Utility that other Wireshark utilities use to capture packets 

## <a name="tshark"></a>tshark

This article does not cover tshark for reasons of brevity. Continue the journey
at the [tshark page]()

# <a name=additional-topics></a>Additional Topics

## <a name=editing-hex></a>Editing Hex

There are a couple ways to edit the hex of a packet capture.  For this scenario,
let's say we want to change all instances of broadcast address 255.255.255.255
in our dhcp.pcap to something else. Let's choose 255.0.255.0 because it's a
funny-looking broadcast address. In hex, this is `0xffffffff` => `0xff00ff00`.

### sed

`sed` gives you the ability to munge filehex. 

`sed -Ei 's/([^\xff])\xff{4}([^\xff])/\1\xff\x00\xff\x00\2/g' dhcp.pcap`

#### Explanation

- `sed -i` : Change in place.
- `sed -E` : Use extended regular expressions
- `\x??` : Hex byte. E.g. `echo -e '\x41'` => `A`, just like an [ASCII
  table](http://www.asciitable.com/) would suggest. Note that a hex byte is 8
  bits and that in `\xff`, each f is 4 bits.  
- `1st [^\xff]` : We know that the 32 bits before this regex will be the
  client's IP address, 0.0.0.0 (0x00000000), and the last byte, 0x00, will match. 
- `2nd [^\xff]` : We know that the 32 bits after this regex are the UDP ports 
  for DHCP, 67 and 68. `[^\xff]` will math the source udp port 68 (00 in 0x0068).
- `\xff{4}`: Given that this packet capture is DHCP, the client
  sends traffic to a MAC address of ffffffffffff. Thus, a
  [regex](https://regexone.com/) of `\xff{4}` will match the dest MAC as well.
  Putting it all together, we get `[^\xff]\xff{4}[^\xff]`. 
- `([^\xff])` Add parentheses (capturing group) to both preceding and trailing
  byte, so they are included in the result
- `\1`, `\2` : We cannot use lookaheand/lookbehind with sed, so use capture
  groups (corresponding to previous) for preceding and trailing bytes

### perl

Exactly like `sed`, except we can use negative lookaheads and lookbehinds:

`perl -pi -e 's/(?<!\xff)\xff{4}(?!\xff)/\xff\x00\xff\x00/g' dhcp.pcap`

### vim & xxd 

If you are using a *nix system (or WSL), [vim](https://www.openvim.com/) and
[xxd](https://linux.die.net/man/1/xxd) are built in and can be used in
conjunction to visually change file bytes. You will need to convert the file
bytes to something readable using `xxd`. `xxd` without options will provide offsets
and spaces between bytes while `xxd -p` will show you just the bytes, both in 16
byte lines. `xxd -r` converts ASCII hex back to the hex literals of your file.
<script id="asciicast-234965" src="https://asciinema.org/a/234965.js" async></script>

### emacs

The joke goes that
"[emacs](https://www.gnu.org/software/emacs/manual/html_node/emacs/Editing-Binary-Files.html)
is a great OS, if only it had a good text editor". Where vim integrates better
with unixy tools like xxd, emacs tries to be your everything.
Case in point: hexl is a builtin that allows for hex literal editing. Open
with `M-x hexl-find-file` and use `C-M-x` to insert hex:
<script id="asciicast-234962" src="https://asciinema.org/a/234962.js" async></script>

### Honorable Mentions

* [hexcurse](https://github.com/arm0th/hexcurse ): curses-based hex editing utility.
* [wxhexeditor](http://www.wxhexeditor.org/): The only cross-platform GUI hex editor with binaries.

## <a name=piping></a>Piping 

Piping is important to using many of these utilities. For example, it is not
really possible to use rawshark without piping as it expects a FIFO or stream. 

| Utility        | stdin formats                 | input formats         | stdout formats  | output formats |
|----------------|-------------------------------|-----------------------|-----------------|----------------|
| **capinfos**   | -                             | all pcaps<sup>1</sup> | text            |                |
| **dumpcap**    |                               |                       | all pcaps       |                |
| **editcap**    |                               |                       | pcap (all?)     |                |
| **mergecap**   |                               |                       | pcap (all?)     |                |
| **randpkt**    | -                             |                       | pcap (all?)     |                |
| **rawshark**   | raw pcap                      |                       | text            |                |
| **reordercap** |                               |                       | pcap (all?)     |                |
| **text2pcap**  | formatted hexdump<sup>2</sup> |                       | raw pcap        | pcap, pcapng   |
| **tshark**     |                               |                       | all pcaps, text |                |

<sup>1</sup> "All pcaps" denotes all pcap types available on the system. You can see them with `tshark -F`.
<sup>2</sup> Formatted hexdump can be canonically generated by `od -Ax -tx1 -v`. As of Wireshark v3.0.0, `tshark -r <my.pcap> -x` will [usually](https://bugs.wireshark.org/bugzilla/show_bug.cgi?id=14639) generate this as well.

## <a name=ssh-capture></a>Capturing over ssh

Let's say that you have a remote computer and want to monitor it remotely.

# <a name=closing-thoughts></a>Closing Thoughts

Personally, I think that wireshark's CLI needs a better API. For example, git
has a large amount of functionality, but.

# Further Reading

* [Official Docs](https://www.wireshark.org/docs/man-pages/)
