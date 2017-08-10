'use strict';

const Koa = require("koa");
const bodyParser = require('koa-bodyparser');
const Router = require("koa-router");
const Boom = require('boom');
const redis = require("redis");
const wrapper = require('co-redis');

const app = new Koa();
const router = new Router();
const client = redis.createClient({
  'host': 'redis',
});
const clientCo = wrapper(client);

app.use(bodyParser());
app.use(async(ctx, next) => {
  const start = Date.now();
  await next();
  const ms = Date.now() - start;
  ctx.set('X-Response-Time', `${ms}ms`);
});

client.on('error', err => {
  console.log('Error ' + err)
});

clientCo.get('count')
  .then(res => {
    console.log(`Currently ${res} lists in redis.`);
    if (res === null) {
      clientCo.set('count', "0")
        .then(res => console.log(res));
    }
  });

router
  .get('/rnd/list/:id', async(ctx) => {
    const lst = await clientCo.smembers(`list:${ctx.params.id}`);
    ctx.assert(lst.length !== 0, 404, 'Error: Not found');
    ctx.body = lst.join(', ');
  })
  .post('/rnd/list', async(ctx) => {
    const count = await clientCo.get('count');
    const id = parseInt(count, 10) + 1;

    const names = ctx.request.body.names ? ctx.request.body.names.split(',') : [];
    ctx.assert(names.length !== 0, 400, 'Error: Empty');

    const res = await clientCo.multi()
      .sadd(`list:${id}`, names)
      .incr('count')
      .exec();
    ctx.body = `OK, ID=${id}`;
  })
  .get('/rnd/:id', async(ctx) => {
    const chosen = await clientCo.srandmember(`list:${ctx.params.id}`);
    ctx.assert(chosen !== null, 400, 'Error');
    ctx.body = chosen;
  })
  .get('/rnd', async(ctx) => {
    ctx.body = 'Welcome to rnd.';
  });

app
  .use(router.routes())
  .use(router.allowedMethods({
    throw: true,
    notImplemented: () => new Boom.notImplemented(),
    methodNotAllowed: () => new Boom.methodNotAllowed()
  }));

app.listen(3000, () => {
  console.log('Rnd listening on port 3000!')
});