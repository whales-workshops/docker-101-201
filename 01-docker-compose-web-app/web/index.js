const fastify = require('fastify')({logger: true})

const path = require('node:path')


fastify.register(require('@fastify/static'), {
    root: path.join(__dirname, 'public'),
})

fastify.register(require("point-of-view"), {
    engine: {
      ejs: require("ejs"),
    },
})

const { ADDRESS = '0.0.0.0', PORT = '8080' } = process.env;


fastify.get('/', (req, res) => {
    return res.view("/templates/index.ejs", {apiURL: process.env.API_URL})
})


fastify.listen({ host: ADDRESS, port: parseInt(PORT, 10) }, (err, address) => {
    if (err) {
        console.error(err)
        process.exit(1)
    }
    console.log(`Server listening at ${address}`)
})
