# Random
A toy that picks out the luckiest one.

## Web app
- Built with Vue.js & Vue Material.
- Randomly choose a name from a name list.
- Uses GitHub login as authentication.
- Store lists and sync them between devices.
- Share lists with friends.
- [Demo here.](https://app.liuyun.me/random)
### Build
Run `npm install && npm build`.

## Go server
- Built using gin-gonic with some middlewares.
- Store data and sessions with PostgreSQL
### Build and run
- Write configs in `.env` file.
- Run `docker-compose up`.
- The server listens on port 7000.
### Usage
Read `routes.go` and guess how to use it.

## Node server
Very very simple APIs bulit with Koa & redis.
### Build and run
- Run `docker-compose up`.
- The server listens on port 3000.
### Usage
#### Test the server
```
$ curl https://api.liuyun.me/rnd/
Welcome to rnd.
```
#### Create a new name list
```
$ curl -d "names=A,B,C" https://api.liuyun.me/rnd/list
A, B, C
```
#### Get an existing name list
```
$ curl https://api.liuyun.me/rnd/list/1
OK, ID=1
```
#### Get a random name
1 is the name list ID.
```
$ curl https://api.liuyun.me/rnd/1
B
```

## WeChat App (Deprecated)
Due to regulations in China, server domains must be whitelisted in order to make requests. So there will be no sync or share features. This version is no longer maintained.
### Usage
Join [B515](https://github.com/B515) organization and you will get access to the private beta.