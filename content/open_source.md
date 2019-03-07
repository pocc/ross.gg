---
title: "Open Source"
date: 2019-03-01T17:17:03Z
author: Ross Jacobs
draft: false
---

## [Wireshark](https://www.wireshark.org/)

### Description

Wireshark is the world’s most widely-used network protocol analyzer. As part of
my role as a Network Engineer at Meraki, I use it on a daily basis. I wanted to
give back, so I donated time to Wireshark.

### Commits

`git clone https://code.wireshark.org/review/wireshark`

- [29bfecc](https://code.wireshark.org/review/#/c/31356/) | Fixed CRC6 lookup
  table and functions.  
  _CRC6 is used for MBMS SYNC, which is a multicast LTE SYNC protocol._ <br>
  <img
  src="https://dl.dropboxusercontent.com/s/kdiz1lnamjstrgd/sync_crc6_checker.webp"
  alt="SYNC CRC6 checker functioning properly."
  style="width:100%;height=100%;text-align:left"/>

## [Tutorial Edge](https://tutorialedge.net)

### Description

Tutorial Edge is a website containing hundreds of polished tutorials. A personal
project of [Elliot Forbes](https://twitter.com/Elliot_F), it covers topics
varying from VueJS to Rust. It uses the hugo framework, and inspired this
website.

### Commits

`git clone https://github.com/elliotforbes/tutorialedge-v2`

- [2f4bbd6](https://github.com/elliotforbes/tutorialedge-v2/commit/2f4bbd6f37b0070ecb9559ff36ff5ad93f1112eb)
  | Wrap all 327 MD files to 80 lines for readability
- [cc14367](https://github.com/elliotforbes/tutorialedge-v2/commit/cc143671357b8ed3f7cb4ba36dd28e9159f284cc)
  | README: Added usage information and expanded authors list
