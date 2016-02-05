HAVOC SERVER
============
[![forthebadge](http://forthebadge.com/images/badges/fuck-it-ship-it.svg)](http://forthebadge.com)

[![Build Status](https://drone.dearing.tech/api/badges/dearing/havoc_server/status.svg)](https://drone.dearing.tech/dearing/havoc_server)

*This is a web server wrapping the [havoc] library.*

ABOUT
-----

I find it tedious to always hunt down something basic enough to demonstrate a working deployment and also be useful enough to tinker with when it is up.  So this is born, a replacement for the ubiquitous [hello_world].  Out of the box it offers [http/pprof] on :8081 and listens for API requests on :8080. Which ought to be configurable in the future. :coffee: 

The only real DATA is a byte array set to the size and acted on by the api itself:

url | affect
:---- | :------
/                 | see the generated name and base64 of the current DATA array
/data/fill        | fill DATA with 1's
/data/fill/crypto | fill DATA with cryto's random (kinda intensive)
/data/reset       | reallocate DATA as 0 and tell the vm to release what memory it feels like
/data/set/*n*     | increase DATA to *n* bytes
/kill             | flat out exit
/procs/*n*        | spin up *n* goroutines running forever (simulate cpu work per core)


TODO & Help Wanted
------------
 - see [issues]

Contributing
------------
1. Fork the repository on Github
2. Create a named feature branch (like `add_component_x`)
3. Write your change
4. Write tests for your change (if applicable)
5. Run the tests, ensuring they all pass
6. Submit a Pull Request using Github

License and Authors
-------------------
Author: Jacob Dearing // jacob.dearing@gmail.com

```
The MIT License (MIT)

Copyright (c) 2016 Jacob Dearing

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
[havoc]: https://github.com/dearing/havoc
[issues]: https://github.com/dearing/havoc_server/issues
[http/pprof]: https://golang.org/pkg/net/http/pprof
[hello_world]: https://github.com/search?q=hello_world&type=Repositories&utf8=%E2%9C%93
