Run this docker container.
Its multithreaded and can listen to both TCP & UDP on port 53

* Sample usage:

```
docker build -t dnsserv .
docker run -d -name=dnssrv --net=host --name=dnsserv -t dnssrv
```

* Note:
--net=host is important


* Verification:
dig @192.168.101.39 A www.indiatimes.com

Credits:
========
- https://github.com/miekg/dns for wonderful library. saved lot of time with rfcs
- https://jameshfisher.com/2017/08/04/golang-dns-server.html

