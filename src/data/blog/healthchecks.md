---
title: "Healthchecks"
draft: false
pubDatetime: 2026-04-21T00:00:00Z
description: "Test Global Cloudflare TCP/TLS/HTTP connectivity to your origin server"
tags:
  - tcp
  - networking
  - tool
---

I built [healthchecks.ross.gg](https://healthchecks.ross.gg) as an vibecoding experiment to test the TCP capabilities of Cloudflare workers. It sends a TCP SYN ping to every colo that I have access to in my personal account. https://ping.pe/ is useful for troubleshooting with ICMP globally, and this tool complements ping.pe with TCP pings from "all" Cloudflare colos. It's possible that this could help you identify regional outages with more granularity. Source is on GitHub as [global-healthchecks](https://github.com/pocc/global-healthchecks).

## Table of contents

# Healthchecks

![](@/assets/images/healthchecks/healthchecks.webp)

## Why You Might Use This

- **"Is my origin up from $REGION?"**: TLS cert renewals, firewall changes, and BGP weirdness often only break from one region. Hitting the host from 143 colos at once surfaces it in seconds.
- **"Is the slowness the network or the app?"**: Timings of different phases (TCP / TLS / HTTP) tell you whether your latency budget is being eaten by the handshake, the cipher negotiation, or the application itself.
- **"Does my cipher policy actually work globally?"**: Pin a min/max TLS version or a specific cipher list and confirm the server negotiates what you expect globally. This obviates the need for spinning up a vps in the right region just to run `openssl s_client`.

## How It Works

The tool performs [TCP pings](https://www.cloudflare.com/learning/network-layer/what-is-a-computer-port/) from **143 Cloudflare Worker endpoints** deployed across the globe. Each endpoint opens a raw [TCP socket](https://developers.cloudflare.com/workers/runtime-apis/tcp-sockets/) to your target host:port and measures the round-trip latency from that location. Unlike ICMP ping, a TCP ping completes the three-way handshake (SYN → SYN-ACK → ACK) to verify the port is actually accepting connections. Results show which data center handled the request and how long the connection took.

## Two Placement Strategies

The 143 endpoints use two different Cloudflare mechanisms to control *where* the Worker executes:

| Strategy | Count | Configured On | Behavior | Region Codes |
| --- | --- | --- | --- | --- |
| [Regional Services](https://developers.cloudflare.com/workers/configuration/regions/) | 10 | DNS record | Worker is **guaranteed** to run inside the target region. Ingress and egress are the same colo. | `us`, `eu`, `jp`, etc. |
| [Region Placement](https://developers.cloudflare.com/workers/configuration/smart-placement/) | 133 | Worker config | Request hits your nearest edge, then is **forwarded** to a colo near the cloud provider region. | `aws:us-east-1`, `gcp:europe-west1`, etc. |

**Ingress Colo** is the data center that first received your request. **Egress Colo** is where the Worker actually executed and ran the TCP test. This is derived from the `cf-placement` response header. These colos will be the same with Regional Services and may differ for Region Placement.

> **Note:** Connections to targets on Cloudflare's network (AS13335) are blocked for security reasons. The test button will be disabled for any target on AS13335.

## Test Modes & OSI Layers

The tool supports two testing modes. **TCP Only** opens a raw socket at the Transport Layer (L4) and measures the three-way handshake. **Full Stack** builds on top of that: after TCP, it establishes a TLS session with configurable version and cipher constraints (L5/L6), then optionally sends an HTTP request and measures time to first byte (L7). Each phase is timed independently.

| Action | OSI Layer | Analogy | Measured |
| --- | --- | --- | --- |
| TCP three-way handshake | Layer 4 — Transport | *Dialing the phone* | TCP ms |
| TLS handshake & session establishment | Layer 5 — Session | *Starting the meeting, agreeing on terms* | TLS ms |
| Cipher selection & encryption | Layer 6 — Presentation | *Choosing the translator* | TLS ms |
| HTTP request & time to first byte | Layer 7 — Application | *Having the conversation* | TTFB |

In a typical browser or `curl` request, TLS session setup happens automatically inside the networking stack, and the caller never touches Layer 5 or 6 directly. This tool is different: the Worker uses [node:tls](https://developers.cloudflare.com/workers/runtime-apis/nodejs/) to act as a *Session Manager*, explicitly controlling the TLS handshake parameters (min/max version, cipher suites, SNI) that normally live below the application's reach. Because you're choosing *which* ciphers to offer and measuring handshake latency and protocol compatibility, **you** operate at **Layer 5-6**, managing the dialogue between client and server, not just consuming it.

## Color Thresholds

| Color | Latency |
| --- | --- |
| 🟢 Green | < 100ms |
| 🟡 Yellow | 100–250ms |
| 🔴 Red | > 250ms |

---

Powered by Cloudflare Workers [Sockets API](https://developers.cloudflare.com/workers/runtime-apis/tcp-sockets/). Try it at [healthchecks.ross.gg](https://healthchecks.ross.gg) or read the source on Github as [global-healthchecks](https://github.com/pocc/global-healthchecks).
