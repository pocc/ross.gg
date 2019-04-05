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
how to use to do X with the CLI, you've come to the right place. As a
contributor to \*shark and daily user, I am writing this as an unofficial
tshark guide.

This guide will help you to capture traffic, edit it, clean it, and send it. The
scenario being that you are reporting on a network problem and want to use
wireshark to provide a packet capture you can then send on to
colleagues/customers.

<!-- Kludgy TOC until I can figure out how to include {{ hugo toc }} in the content -->
## Table of Contents

* [X] [Getting Started](/post/wireshark-setup) 
  <!-- [[wireshark_setup]] -->
	* [X] Installing 3.0.0](/post/wireshark-setup/#install)
* [ ] [Capture](/post/wireshark-capturing#capture) 
  <!-- [[wireshark_capturing]] --> 
	* [ ] [dumpcap](/post/wireshark-capturing#dumpcap) 
	* [ ] [tshark](/post/wireshark-capturing#tshark)
* [o] [Analyze](/post/wireshark-info#info) 
  <!-- [[wireshark_info]] -->
	* [ ] [Syntax]
		* [ ] What is BPF
		* [ ] How does Wireshark syntaxt work?
		* [ ] Testing your filter: dftest
	* [X] [capinfos](/post/wireshark-info#capinfos)  
	* [ ] [captype]
	* [ ] tshark -G
	* [X] [rawshark](/post/wireshark-info#rawshark)
* [X] [Generate](/post/wireshark-generation#generate) 
  <!-- [[wireshark_generation]] -->
	* [X] [randpkt](/post/wireshark-generation#randpkt) 
* [ ] [Edit](/post/wireshark-editing#edit) 
  <!-- [[wireshark_editing]] -->
	* [ ] [editcap](#editcap)
	* [ ] [mergecap](#mergecap)
	* [ ] [reordercap](#reordercap)
	* [ ] [text2pcap](#text2pcap)
* [o] [Additional Topics](/post/wireshark-bonus-topics#additional-topics)  
  <!--[[wireshark_bonus]] -->
	* [ ] [Export Object](/post/wireshark-export-object)
	* [X] [Editing Hex](/post/wireshark-bonus-topics#editing-hex)
	* [X] [Piping](/post/wireshark-bonus-topics#piping) 
* [ ] [Unusual Interfaces and Where to Find Them]  
  <!--[[wireshark_livecaptures]] -->
	* [ ] Add Gif of chrome download live capture
	* [ ] Add Scapy gif of live capture
* [ ] [extcap: Make your own interface]  
  <!--[[wireshark_extcap]] -->
	* [ ] randpkt 1
	* [ ] randpkt 2
	* [ ] randpkt 3
	* [ ] randpkt 4
	* [ ] <using randpkt gif (upload from desktop as gif/webp)>
	* [ ] Wireshark extcap_example.py with GUI and screen recording

## <a name=closing-thoughts></a>Closing Thoughts

Personally, I think that wireshark's CLI needs a better API. For example, git
has a large amount of functionality, but.

## Further Reading

_The end of one adventure is the beginning of another._

### Network Scripting with Python

* [Python for Network Engineers](https://www.youtube.com/watch?v=s6SIVc7C5U0):
  David Bombal is a CCIE who has good lectures on using Python (costs $$$)
* [Sentdex Tutorials](https://www.youtube.com/user/sentdex): A Pythonista who
  will inspire you
* [Python Guide](https://docs.python-guide.org/): For when you want to turn your
  script into a project.

### Wireshark

* [Official Docs](https://www.wireshark.org/docs/man-pages/)
* [Get the Sourcecode](https://www.wireshark.org/develop.html)
* [File a Bug Report](https://wiki.wireshark.org/ReportingBugs)
* [Contribute!](https://www.wireshark.org/docs/wsdg_html_chunked/)

