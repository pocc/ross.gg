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

### import

| urllib                              | requests          |
|-------------------------------------|-------------------|
| `from urllib import request, error` | `import requests` |

### init vars

_These variables are required for calls below_

| urllib                                                                                                                                                                             | requests                                                                                                                        |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|
| {{< highlight python3 >}}site = "https://httpbin.org/"<br>data = {'key': 'value'}<br>headers = {'Accept': '*/*',<br>&nbsp;&nbsp;&nbsp;&nbsp;'Accept-Encoding': 'gzip, deflate'<br>}{{< /highlight >}} | {{< highlight python3 >}}site = "https://httpbin.org/"<br>data = {'key': 'value'}<br># headers are en/decoded as gzip by default{{< /highlight >}} |

<table>
<thead>
<tr>
<th>urllib</th>
<th>requests</th>
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
from urllib import request, error
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
# default headers ask to use gzip
# default receive behavior is to decode gzip{{< /highlight >}} 
	</td>
  </tr>
</tbody>
</table>

# 
site = "https://httpbin.org/"
headers = {'Accept': '*/*', 'Accept-Encoding': 'gzip, deflate'}
data = {'key': 'value'} 

# GET
requests_resp = requests.get(site + "get")
# --- 
urllib_obj = urllib.request.Request(site + "get")
urllib_resp = urllib.request.urlopen(urllib_obj)

# POST
requests_resp = requests.get(site + "post", data=data)
# --- 
encoded = urllib.parse.urlencode(data).encode('utf-8') 
urllib_resp = urllib.request.urlopen(site + "post", encoded)

# PUT
requests_resp = requests.get(site + "put", data=data)
# --- 
encoded = urllib.parse.urlencode(data).encode('utf-8') 
urllib_obj = urllib.request.Request(site + "put", data=encoded, method="PUT")
urllib_resp = urllib.request.urlopen(urllib_obj)

# DELETE
requests_resp = requests.delete(site + "delete")
# --- 
urllib_obj = urllib.request.Request(site + "delete", method="DELETE")
urllib_resp = urllib.request.urlopen(urllib_obj)

# OPTIONS
requests_resp = requests.options(site + "get")
print(requests_resp.headers)
# ---
urllib_obj = urllib.request.Request(site + "get", method="OPTIONS")
urllib_resp = urllib.request.urlopen(urllib_obj)
print(urllib_resp.headers)


# Returned object type:
type(r)   # requests
type(ur)  # http.lib.HTTPConnection
```

For this table, assume that obj = response class object

| Property                     | requests                                             | urllib                       |
|------------------------------|------------------------------------------------------|------------------------------|
| _response class_             | _requests_                                           | [_http.lib.HTTPResponse_][1] |
|                              |                                                      |                              |
|------------------------------|------------------------------------------------------|------------------------------|
| Does this work?                                                                                                  |
|------------------------------|------------------------------------------------------|------------------------------|
| Close the connection         | obj.close()                                          |                              |
|                              | obj.iter_content(chunk_size=1, decode_unicode=False) |                              |
| Get underlying socket fileno | -                                                    | obj.fileno()                 |
|                              |                                                      |                              |
| server's response headers    |                                                      | obj.msg                      |
| status code                  | obj.status_code                                      | obj.status                   |
| status reason                | obj.reason                                           | obj.reason                   |
| server HTTP version          | -                                                    | obj.version                  |
| Is stream closed?            | -                                                    | obj.closed                   |

[1]: https://docs.python.org/3/library/http.client.html#httpresponse-objects



## Further Reading

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
