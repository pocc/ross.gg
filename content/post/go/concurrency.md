---
title: "Go Concurrency"
date: 2019-03-04T16:20:52Z
author: Ross Jacobs
desc: "Is it overused?"
draft: true
---

(Look at and INTEGRATE the sieve of eratosthenes files in same folder)

## Sharing is Caring
One of Go's philosophies is to [Share Memory by Communicating](https://blog.golang.org/share-memory-by-communicating). Goroutines and channels are a unique feature to Go, and I was wondering about performance hits to carelessly using them where they're not necessary.

## Benchmarking channel IO

I created an [arbitrary test](https://gist.github.com/pocc/075b38f097e68ae90a07efde6de3daa5) with summing. These are my unscientific findings:

Benchmark | # iterations | Rate
-------------------------|----------|----------
BenchmarkSerial          |2000000000|0.32 ns/op
BenchmarkChanBuffered    |  30000000|     50.9 ns/op
BenchmarkChanUnbuffered  |   3000000|     422 ns/op

Other articles find a similar 300-400ns channel penalty.

## Similar tests/articles

* Article: [testing channel throughput](https://syslog.ravelin.com/so-just-how-fast-are-channels-anyway-4c156a407e45) article 
* Gist: [testing sending different types to a channel](https://gist.github.com/atotto/9342938)
* Article: [Channels!!!](https://www.jtolio.com/2016/03/go-channels-are-bad-and-you-should-feel-bad/) and related [Hacker News discussion](https://news.ycombinator.com/item?id=11210578)

## Closing thoughts

I wrote this up because I was overusing channels. To quote the top comment in the HN discussion, "APIs should be designed synchronously, and the callers should orchestrate concurrency if they choose". This is a good heuristic...
