---
title: "Projects"
description: "Things I am proud of"
date: 2019-03-01T17:17:03Z
author: Ross Jacobs
draft: false
---

|         | Merlink        | PcapGraph  | MAD CodeGen |
| ------- | -------------- | ---------- | ----------- |
| Status  | Alpha          | Alpha      | Alpha       |
| Install | PyPI, binaries | PyPI       | -           |
| License | Apache 2.0     | Apache 2.0 | Apache 2.0  |

## [Merlink](https://github.com/pocc/merlink)

_Meraki VPN Client_

A GUI client that will connect desktop devices to Meraki firewalls via an
L2TP/IPSEC connection. This program uses a Meraki dashboard admin's credentials
to pull the data required for a Client VPN connection, create the connection
with OS built-ins, and then connect. Using the PyQt5 framework, this is
available for Linux, Macos, and Windows.

## [PcapGraph](https://github.com/pocc/pcapgraph)

_Analyze multiple Packet Captures_

<img 
  src="https://github.com/pocc/pcapgraph/raw/master/examples/pcap_graph.png?raw=true"
  alt="3 Packet Captures taken consecutively" 
  style="height:60%;width:60%;margin:0px" />

Assists with flow-based troubleshooting where there are at least 2 pcaps. It can
create a horizontal bar graph to visualize when pcaps were taken to verify that
packet captures taken across multiple interfaces or devices are contemporaneuos.
Set operations can also be performed on packet captures where packets are
elements. This can be useful when trying to establish whether there is any
traffic across packet captures that is **EXACTLY** the same.

## [Merakygen](https://github.com/pocc/mad-codegen)

_API Clients generated from Meraki API Docs_

The Meraki Engineering team is continuously releasing new API endoints. Client
libraries are maintained by open source contributors, but often have troule
keeping up. The Meraki API docs contain enough information to automate the
creation of these client libraries. MAD CodeGen will generate an API Client in
your \$language based upon the most recent release of the API docs.
