# URL shortner API server

Golang URL shortener API server

## Getting started

### Prerequisites

- [Install docker compose](https://docs.docker.com/compose/install/)

### Configuration

Add `.env` to the repository, and set the required environment variables:

```
RDS_ADDR=<redis-server address>
RDS_USER=<redis-username>
RDS_PWD=<redis-password>
```

### Run with docker-compose

```shell
$ docker-compose up
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
    curl --request GET "localhost:8080/shorten?longUrl=$longUrl"
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

- Handle duplicated key/url