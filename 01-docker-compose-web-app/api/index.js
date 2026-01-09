const fastify = require('fastify')({logger: true})

const path = require('node:path')

fastify.register(require('@fastify/redis'), { 
    host: 'redis-server', 
    port: 6379
})

fastify.register(require('@fastify/cors'), (instance) => {
    return (req, callback) => {
      const corsOptions = {
        origin: process.env.FRONT_URL || "http://localhost:7070"
      };
      // callback expects two parameters: error and options
      callback(null, corsOptions)
    }
})


const { ADDRESS = '0.0.0.0', PORT = '8080' } = process.env;

fastify.get('/counter/:animal', (request, reply) => {
    const { animal } = request.params;

    const { redis } = fastify

    redis.get(animal, (err, counter_value) => {
        var counter = err ? 0 : Number(counter_value)
        reply.send(counter)
    })   
})

fastify.post('/counter', (request, reply) => {
    const { redis } = fastify
    let animal = request.body.animal

    redis.get(animal, (err, counter_value) => {
        var counter = err ? 0 : Number(counter_value)

        redis.set(animal, counter+1, (err) => {
            reply.send(err || { status: 'ok' })
        })
    })   
})

fastify.listen({ host: ADDRESS, port: parseInt(PORT, 10) }, (err, address) => {
    if (err) {
        console.error(err)
        process.exit(1)
    }
    console.log(`Server listening at ${address}`)
})
