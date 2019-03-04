---
title: "Using nl"
date: 2019-03-03T14:53:58Z
author: "Ross Jacobs"
draft: true
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
