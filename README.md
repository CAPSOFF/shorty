# AMARTHA-SHORTY

<!-- toc -->

- [Overview](#overview)
- [Requirement](#requirement)
- [Docker](#docker)
- [API](#API)

<!-- tocstop -->

## Overview

this repository is made to answer AMARTHA pre-test question for Senior Backend Software Engineer ([Shorty](https://gist.github.com/williamn/cfad86ab218101e0c5d7be89226c5c85)), using Golang, Docker

## Requirement

- Go
- Docker
- GinGonic

## Docker
- Clone repo
```bash
$ git clone https://github.com/CAPSOFF/shorty.git
```

- Build Docker Images
```bash
$ docker build -t amartha-shorty:latest . 
```

- Run Docker Container
```bash
$ docker run --net host -it amartha-shorty:latest bash
$ cd /usr/local/bin/
$ ./main
```

<br>
tested on Ubuntu 18.04.2 LTS x86_64

## API

- Shorten
```
method      : POST
endpoint    : http://localhost:7777/api/v1/amartha/shorty/shorten
request     : {
                "url": "https://blog.trello.com/navigate-communication-styles-difficult-times",
                "shortCode": "abcdef"
              }
response    : {
                "shortCode": "abcdef"
              }
```

- ShortCode
```
method      : GET
endpoint    : http://localhost:7777/api/v1/amartha/shorty/:shortcode
request     : :shortcode = <your-saved-short-code>
response    : <redirect-to-specific-url>
```

- ShortCodeStats
```
method      : GET
endpoint    : http://localhost:7777/api/v1/amartha/shorty/:shortcode/stats
request     : :shortcode = <your-saved-short-code>
response    : {
                "lastSeenDate": "2021-06-15 00:11:14.477166127 +0700 WIB m=+13.574710341",
                "redirectCount": 3,
                "startDate": "2021-06-15 00:11:03.8630576 +0700 WIB m=+2.960601814"
              }
```