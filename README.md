# URL shortener

Fast URL shortener service written in Go

# API

Get shortened URL

```bash
curl -sX POST -H 'Content-Type: application/json' 'localhost/shorten' -d '{"url":"http://a.very.long.url"}'

> HTTP 200
> '{"short":"http://localhost/abcdef"}'
```

Get original URL

```bash
curl -sX GET -H 'Content-Type: application/json' 'localhost/original' -d '{"short":"http://localhost/abcdef"}'

> HTTP 200
> {"original":"http://a.very.long.url"}
```

## Features

* no revoke/update
* MRU

## Dependencies

* httprouter
* groupcache
* mgo

# License

MIT

