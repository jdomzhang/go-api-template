{{description}}

# Prerequisites

## `yarn`

> npm install -g yarn

or, if you have `scoop` installed on windows

> scoop install yarn

## `fresh` - live reload

> go get -u github.com/jdomzhang/fresh

## `hygen` - code generation

> yarn global add hygen

## 'goconvey'

> go get -v github.com/smartystreets/goconvey

# Database

## `PostgreSQL`

Please check `scripts/create-database.sql` for more details

# Commands

## Dev

> yarn dev

## Deploy & release at first time

> yarn install:all

### Manually setup nginx

> vim /etc/nginx/sites-available/default

- See the sample configure at the sample file `nginx.conf` under `scripts` folder of this project

### Manually config DNS

- Go to domain provider website to map your domain accordingly

## Furthur release

> yarn release
