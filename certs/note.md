#### Recommendation: For production use https://letsencrypt.org/getting-started/# is free or https://www.ssls.com/ has economical price for certificates, for domain name https://www.domainnameapi.com/pricing?po=register or https://www.namecheap.com/


## Create RootCA

argument [ -aes256 ] is optional for encryption with password

- Create Root CA private Key
```
openssl genpkey -algorithm RSA -out rootCA.key 
```
- Create CA Certificate
```
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1825 -out rootCA.crt
```

## Server Public Certificate
- Create server private Key
```
openssl genpkey -algorithm RSA -out server.key
```
- Create Certificate Signing Request (CSR)
```
openssl req -new -key server.key -out server.csr
```
- Sing with rootCA
```
openssl x509 -req -in server.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out server.crt -days 999 -sha256
```

*For this example, the same pair of keys is used on the server as in the app, but in production this should never be done this way.*