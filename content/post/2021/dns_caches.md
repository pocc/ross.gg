---
title: "DNS Caches"
date: 2021-10-11T19:24:42Z
author: Ross Jacobs
desc: "Regular post"
tags: 
  - dns
  - caching
  - networking
draft: false
---

*This article shows how to get browser and OS DNS caches.*

According to [Cloudflare docs](https://www.cloudflare.com/learning/dns/what-is-dns/), the sequence of a DNS query from Chrome to the recursive resolver looks something like this:

```diagram
Check browser DNS cache --miss--> Check OS DNS cache --miss--> Recursive resolver
          |                               |
          V                               V
         hit                             hit
```

## Firefox DNS Cache

Firefox's DNS cache `about:networking#dns` can be flushed with a button and lists entries.

![firefox dns cache](/img/firefox_dns_cache.png)

## Chrome DNS Cache

Chrome's DNS cache `chrome://net-internals/#dns` can be flushed with a button.
It used to be possible to see the list entries in Chrome easily, but it's since been obfuscated.
To list entries in Chrome:

1. Go to `chrome://flags/#enable-network-logging-to-file` in Chrome
2. Set it to enabled
3. Relaunch Chrome using the blue button provided that pops up on enabling/disabling of settings
4. Go to `chrome://net-export/` to start your capture and see the file location
5. Import your  `chrome-net-export-log.json` file with [https://netlog-viewer.appspot.com/#import](https://netlog-viewer.appspot.com/#import)
6. Go to the [DNS section](https://netlog-viewer.appspot.com/#dns) section

![chrome dns cache](/img/chrome_dns_cache.png)

## Windows

On windows, you can use the `ipconfig /displaydns` command like so:

```powershell
PS C:\> ipconfig /displaydns

Windows IP Configuration


    chrome.cloudflare-dns.com
    ----------------------------------------
    Record Name . . . . . : chrome.cloudflare-dns.com
    Record Type . . . . . : 1
    Time To Live  . . . . : 54
    Data Length . . . . . : 4
    Section . . . . . . . : Answer
    A (Host) Record . . . : 104.18.27.211


    vortex.data.microsoft.com
    ----------------------------------------
    Record Name . . . . . : vortex.data.microsoft.com
    Record Type . . . . . : 5
    Time To Live  . . . . : 6
    Data Length . . . . . : 8
    Section . . . . . . . : Answer
    CNAME Record  . . . . : asimov.vortex.data.trafficmanager.net

```

## Linux

In linux, if you are using systemd, you can see them like so:

```bash
time=$(date "+%F %T"); systemd-resolve cloudflare.com >/dev/null; systemctl kill -s USR1 systemd-resolved; journalctl -b -0 --since "$time" -u systemd-resolved | grep " IN "
Oct 10 22:21:23 vps systemd-resolved[3255524]:         vortex.data.microsoft.com IN CNAME asimov.vortex.data.trafficmanager.net
Oct 10 22:21:23 vps systemd-resolved[3255524]:         cloudflare.com IN AAAA 2606:4700::6810:85e5
Oct 10 22:21:23 vps systemd-resolved[3255524]:         cloudflare.com IN AAAA 2606:4700::6810:84e5
...
```

## Macos

It looks like there are different ways to do this per OS minor version.

On 10.14, which I can test, you can use mDNSResponder to get data. You need to make sure that records are not private so that you can see them.

```bash
sudo log config --mode "private_data:on"
dns_caches=$(log stream --predicate 'process == "mDNSResponder"' --info)
echo "$dns_caches" | rg -U '^[\s\S]*?\- Cache \-+\n.*rdlen\n([\s\S]*?)\n.*?Cache size[\s\S]*?$' -r '$1'
```

More information can be found with this [stackoverflow question](https://stackoverflow.com/questions/38867905/how-to-view-dns-cache-in-osx).
