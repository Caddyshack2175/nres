# nres
Resolve Nessus URL from a Nessus Report. 
This resolver uses api.tenable.com not nessus.org this action seems to to take less time with a halt on the redirect getting the URI from the Location header.

`$ nres http://www.nessus.org/u?cc4a822a
http://www.nessus.org/u?cc4a822a => https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Cipher-block_chaining_.28CBC.29
`
