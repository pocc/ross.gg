---
title: "Using urllib or requests"
date: 2019-03-14T11:35:30Z
author: Ross Jacobs
desc: "How to use Python's http libraries"
tags: 
  - python
  - urllib
image: https://i.stack.imgur.com/wMYZG.png

draft: true
---

_What Python http library should you use?_

If you use Python 2, please use [this article](https://docs.python.org/3/howto/pyporting.html) first.

[`urllib`](https://docs.python.org/3/library/urllib.html) and
[`requests`](http://docs.python-requests.org/en/master/) are two of many http
libraries. `urllib` is a part of the Python standard library while requests is
not. `requests` is an API around other Python libraries and is easier to use.
In fact, it is so well written, that it is [recommended reading](citation) for
Python neophytes. However, it is not in the Python standard library, and this
will also not change any time soon [^githubissue] [^docs] [^lwn]. In this article, I will go over
how to use each to do the same things and why you should choose one or othe
other for your project. 

If in doubt, go with requests. 

[^githubissue]: [Asking the community for input in Github Issue](https://github.com/kennethreitz/requests/issues/2424)
[^docs]: [Kenneth Reitz on stdlib inclusion](http://docs.python-requests.org/en/master/dev/philosophy/#standard-library)
[^lwn]: [LWN discussion based on (2)](https://lwn.net/Articles/640838/)

## urllib
`urllib` has these characteristics: 

  A) Part of stdlib 
  B) Inconsistent API that is a pain to use

Reasons to use:

* Project collaborators do not have pip installed
* System does not have internet access/pip
* Desire to minimize project size
  * You may want to choose a different language
    if optimization is important
* Desire to have no project dependencies
* This is designed for small scripts for yourself

## requests
`requests` has these characteristics:

  A) Is not a part of and will not be added to stdlib
  B) Beuatiful API that is easy to use 

Requests to use:

* It is more readable, so makes code more maintainable
* If you are a part of a team 
* If there are more than >3 requests made in project
* If you feel like you are writing wrappers for urllib's interface
* If you are new to Python / are prototyping 

## Alike and not

## Setup

<table>
  <thead>
    <tr>
      <th><h6 style="color:gray;">urllib</h6></th>
      <th><h6 style="color:gray;">requests</h6></th>
    </tr>
  </thead>

  <tbody>
    <tr>
      <td><h6>Import</h6></td>
	  <td></td>
    </tr>
    <tr>
      <td>
        {{< highlight python3 >}}
from urllib import request, parse, error
import gzip{{< /highlight >}}
      </td>
      <td>
        {{< highlight python3 >}}import requests{{< /highlight >}} 
      </td>
    </tr>
    <tr>
      <td><h6>Init Vars</h6></td>
	  <td></td>
    </tr>
    <tr>
      <td>
        {{< highlight python3 >}}
site = "https://httpbin.org/"
data = {'key': 'value'}
headers = {'Accept': '*/*', 
    'Accept-Encoding': 'gzip, deflate'}{{< /highlight >}} 
	  </td>
	  <td>
        {{< highlight python3 >}}
site = "https://httpbin.org/"
data = {'key': 'value'}
"""default headers' Accept-Encoding is  gzip
default receive behavior is to decode gzip"""{{< /highlight >}} 
	  </td>
    </tr>
  </tbody>
</table>

## Basic Options

<table>
  <thead>
    <tr>
      <th><h6 style="color:gray;">urllib</h6></th>
      <th><h6 style="color:gray;">requests</h6></th>
    </tr>
  </thead>
  
  <tbody>
    <tr>
      <td><h6>GET</h6></td>
	  <td></td>
    </tr>
    <tr>
      <td>{{< highlight python3 >}}
ul_obj = request.Request(site + "get", 
                         headers=headers)
resp = request.urlopen(ul_obj){{< /highlight >}} 
	  </td>
	  <td>{{< highlight python3 >}}
resp = requests.get(site + "get"){{< /highlight >}} 
      </td>
    </tr>
    <tr>
      <td><h6>POST</h6></td>
	  <td></td>
    </tr>
    <tr>
      <td>{{< highlight python3 >}}
encoded = parse.urlencode(data).encode('utf-8') 
resp = request.urlopen(site + "post", encoded){{< /highlight >}} 
	  </td>
	  <td>{{< highlight python3 >}}
resp = requests.post(site + "post", data=data){{< /highlight >}} 
      </td>
    </tr>
    <tr>
      <td><h6>PUT</h6></td>
	  <td></td>
    </tr>
    <tr>
      <td>{{< highlight python3 >}}
encoded = parse.urlencode(data).encode('utf-8') 
ul_obj = request.Request(site + "put", 
					     data=encoded, 
						 method="PUT")
resp = request.urlopen(ul_obj){{< /highlight >}} 
	  </td>
	  <td>{{< highlight python3 >}}
resp = requests.put(site + "post", data=data){{< /highlight >}} 
      </td>
    </tr>
    <tr>
      <td><h6>DELETE</h6></td>
	  <td></td>
    </tr>
    <tr>
      <td>{{< highlight python3 >}}
ul_obj = request.Request(site + "delete", 
                         method="DELETE")
resp = request.urlopen(ul_obj){{< /highlight >}} 
	  </td>
	  <td>{{< highlight python3 >}}
resp = requests.delete(site + "delete"){{< /highlight >}} 
      </td>
    </tr>
    <tr>
      <td><h6>PATCH</h6></td>
	  <td>Modify a resource</td>
    </tr>
    <tr>
      <td>{{< highlight python3 >}}
{{< /highlight >}} 
	  </td>
	  <td>{{< highlight python3 >}}
	  {{< /highlight >}} 
      </td>
    </tr>
     <tr>
      <td><h6>OPTIONS</h6></td>
	  <td><i>Get the available operations from the server</i></td>
    </tr>
    <tr>
      <td>{{< highlight python3 >}}
ul_obj = request.Request(site + "get", 
                         method="OPTIONS")
resp = request.urlopen(ul_obj)
options = resp.headers{{< /highlight >}} 
	  </td>
	  <td>{{< highlight python3 >}}
resp = requests.options(site + "get")
options = resp.headers{{< /highlight >}} 
      </td>
    </tr>
  </tbody>
</table>

For this table, assume that resp = response class object

| Property                       | requests                                              | urllib                       |
|--------------------------------|-------------------------------------------------------|------------------------------|
| _response class_               | _requests_                                            | [_http.lib.HTTPResponse_][1] |
|--------------------------------|-------------------------------------------------------|------------------------------|
| Return resp class (type(resp)) | requests.Response                                     | http.lib.HTTPConnection      |
| Close the connection           | resp.close()                                          |                              |
|                                | resp.iter_content(chunk_size=1, decode_unicode=False) |                              |
| Get underlying socket fileno   | -                                                     | resp.fileno()                |
|                                |                                                       |                              |
| server's response headers      |                                                       | resp.msg                     |
| status code                    | resp.status_code                                      | resp.status                  |
| status reason                  | resp.reason                                           | resp.reason                  |
| server HTTP version            | -                                                     | resp.version                 |
| Is stream closed?              | -                                                     | resp.closed                  |

[1]: https://docs.python.org/3/library/http.client.html#httpresponse-objects



## Further Reading

* [Python docs urllib HOWTO](https://docs.python.org/3.7/howto/urllib2.html)

**Questions and Exercises**

- Question/Exercise
- Question/Exercise

**Relevant Articles**

- [Article 1]()
- [Article 2]()

**Sources** [0]() [1]()

## Drafting

### Prewriting

**Audience**

Who is your audience?

**Deliverable**

What is the ONE thing your audince gain from reading this?

**Niche**

What makes this unique compared to existing articles?

### Checklist

** Basic**

- [ ] Intro: How WILL they get the deliverable?
- [ ] 300-600 words
- [ ] Images: Cover image, Reengage image/table
- [ ] Conclusion: How DID they get the deliverable?
- [ ] Questions/Exercises/Call To Action

**Extended**

- [ ] Keywords: Front Matter, Title, Desc, Post: (top, end), Images: (alt,
      title)
- [ ] 3-4 external links
- [ ] 1-2 sources
- [ ] 2-4 internal links
- [ ] Lint!

### Prepublish

- Engaging: Why will the reader read until the end?
- Organized: Identify specific things that the reader might be looking for in
  subsections. How easy are they to find?
- Optimized: Can the Deliverable be provided to the reader in fewer words?
