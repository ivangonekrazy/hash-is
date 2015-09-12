HashIs
============

Ivan Tam (ivan@hipnik.net)

SHA1 hexdigest as a service via WHOIS.

1. Start HashIs. It will bind to port 43, so sudo may be required.
2. In another terminal: `whois -h <hostname> <string...>`
3. HashIs will return a SHA1 hexdigest for each string.
