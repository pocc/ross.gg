---
title: "Building Your Tshark Capture Command"
date: 2019-03-12T12:44:45Z
author: Ross Jacobs
desc: "Like building a regex but more fun!"
tags:
  - networking
  - tshark
image: https://allabouttesting.org/wp-content/uploads/2018/06/tshark-count.jpg

draft: true
---

1. Determine your interface 

If you run `ping 8.8.8.8 >/dev/null & tshark`, you should start seeing
numbered packets. If you don't, you should find out what interfaces you have
available, as the one you are currently using is not working. `tshark -D`
will showe you a list of available interfaces. If you are unsure of which
interface is the default one, you can use `ifconfig` on \*nix and
`ipconfig /all` on Windows. These will print the exact name:

```sh
# Using powershell on Windows
Get-NetAdapter | where {$_.Status -eq "Up"} | Select -ExpandProperty Name
# BSD & Macos
route get default | awk '/interface:/{print $NF}'
# Linux
route | awk '/^default|^0.0.0.0/{print $NF}'
```

finished p1 > resume from there