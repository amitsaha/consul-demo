dig @127.0.0.1 -p 8600 v2.webapp.service.consul SRV

; <<>> DiG 9.10.3-P4-Ubuntu <<>> @127.0.0.1 -p 8600 v2.webapp.service.consul SRV
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NXDOMAIN, id: 41491
;; flags: qr aa rd; QUERY: 1, ANSWER: 0, AUTHORITY: 1, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;v2.webapp.service.consul.      IN      SRV

;; AUTHORITY SECTION:
consul.                 0       IN      SOA     ns.consul. hostmaster.consul. 1510814793 3600 600 86400 0

;; Query time: 3 msec
;; SERVER: 127.0.0.1#8600(127.0.0.1)
;; WHEN: Thu Nov 16 17:46:33 AEDT 2017
;; MSG SIZE  rcvd: 103


