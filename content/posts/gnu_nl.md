---
title: "Using nl"
date: 2019-03-03T21:53:58Z
author: "Ross Jacobs"
draft: true
image: "https://dl.dropboxusercontent.com/s/tkzxe15a057rhag/gnu_nl_lightgray.webp"
---

There are many GNU coreutils, some of which are more useful than others.  

## Basic Usage

## Using regex
This is an example of what nl can be:
```
$ nl arith.jl | sed -n '/# .*/p' 
     1	# Simple arithmetic function
     5	#    return 0
``` 
