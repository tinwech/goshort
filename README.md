# URL shortner API server

Golang URL shortener API server

## Getting started

### Run the server

```shell
$ RDS_ADDR=<redis server address> \
RDS_USER=<username> \
RDS_PWD=<password> \
go run server.go
```

### Usage

- `GET /api/shorten`

    |Request Field   |Type   |Description   |
    |---|---|---|
    |longUrl   |string   |URL to be shorten   |

    |Response Field   |Type   |Description   |
    |---|---|---|
    |shortUrl   |string   |The shorten URL   |

    example:

    ```
    curl --location --request POST "localhost:8080/shorten?longUrl=$longUrl"
    ```

- `GET /api/expand`

    |Request Field   |Type   |Description   |
    |---|---|---|
    |shortUrl   |string   |The short URL   |

    |Response Field   |Type   |Description   |
    |---|---|---|
    |longUrl   |string   |Original URL   |

    
    example:

    ```
    curl --request GET "localhost:8080/api/expand?shortUrl=$shortURL"
    ```
## TODOs

There are some features plan to be done in the future:

- Load balancing
- handle duplicated key
- handle duplicated URL