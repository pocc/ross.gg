---
title: "Using Urllib"
date: 2019-03-14T11:35:30Z
author: Ross Jacobs
desc: "Using Urllib"
tags: 
  - python
  - urllib
image: https://i.stack.imgur.com/wMYZG.png

draft: true
---

_Urllib is a part of the Python Standard Library_

[Urllib Docs](https://docs.python.org/3/library/urllib.html)

[Requests](http://docs.python-requests.org/en/master/) makes Python web scraping
easy. It also has an inspiring API, which is recommended reading for Python
neophytes. However, requests is not in the Python
standard library, and this will also not change any time soon [0] [1] [2]. I do not
think this is a bad thing though: Requests is an API around existing Python
libraries. If you want to use this easy to use API, that is your choice; however
you will be adding a dependency to your project.

In my opinion reducing external dependencies makes for cleaner code.
This article is aimed at people who are willing to use a clunkier http interface
(urllib) to reduce dependencies. This article provides examples of urllib and how to
migrate from requests to urllib. If you feel yourself wanting to simplify urllib
or write a wrapper around functionality, just use requests.

[0]: [Asking the community for input](https://github.com/kennethreitz/requests/issues/2424)
<br>[1]: [Kenneth Reitz on stdlib inclusion](http://docs.python-requests.org/en/master/dev/philosophy/#standard-library)
<br>[2]: [LWN discussion based on [1]](https://lwn.net/Articles/640838/)

## Asserts
- You are a python programmer who interacts with websites 

## What reasons are there for sticking with Requsets? 
(urllib might be behind the times when it comes to ssl)


http://docs.python-requests.org/en/master/user/quickstart/

```python
import requests
import urllib.request
import urllib.error

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
