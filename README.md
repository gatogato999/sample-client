# sample-client
a simple go client for the simple web-server.

## installation

- install go 
- clone the repo
- install dependencies

## usage 

```bash
./app query -email <an email> -password <least 8charcters>
```

## to implement

- [x] ping the base route
- [x] query for a user 

- [x] all routes are protected
- [x] fix jwt failure path and make apply to all os file systems
- [x] handling errors (returned errors)
- [x] better arugment parsing (use flagset)
- [x] make /query to use GET (it send no body)
- [x] add proper HTTP headers

## future modifications

- [ ] do auth using refresh token
- [ ] beatify output
- [ ] make a UI (web page)
