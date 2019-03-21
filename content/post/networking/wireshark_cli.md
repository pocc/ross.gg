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

* [ ] [Getting Started]() 
  <!-- [[wireshark_setup]] -->
	* [ ] Installing 3.0.0
* [ ] [Capture](/post/wireshark-capturing#capture) 
  <!-- [[wireshark_capturing]] --> 
	* [ ] [dumpcap](/post/wireshark-capturing#dumpcap) 
	* [ ] [tshark](/post/wireshark-capturing#tshark)
* [X] [Generate](/post/wireshark-generation#generate) 
  <!-- [[wireshark_generation]] -->
	* [X] [randpkt](/post/wireshark-generation#randpkt) 
* [ ] [Edit](/post/wireshark-editing#edit) 
  <!-- [[wireshark_editing]] -->
	* [ ] [editcap](#editcap)
	* [ ] [mergecap](#mergecap)
	* [ ] [reordercap](#reordercap)
	* [ ] [text2pcap](#text2pcap)
* [X] [Info](/post/wireshark-info#info) 
  <!-- [[wireshark_info]] -->
	* [X] [capinfos](/post/wireshark-info#capinfos)  
	* [X] [rawshark](/post/wireshark-info#rawshark)
* [.] [Additional Topics](/post/wireshark-bonus-topics#additional-topics)  
  <!--[[wireshark_bonus]] -->
	* [X] [Editing Hex](/post/wireshark-bonus-topics#editing-hex)
	* [ ] [Piping](/post/wireshark-bonus-topics#piping) 
	* [ ] [Capturing over ssh](/post/wireshark-bonus-topics#ssh-capture)

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

