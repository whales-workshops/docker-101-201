const express = require('express')
const app = express()

const client = require('prom-client')

let collectDefaultMetrics = client.collectDefaultMetrics
const register = new client.Registry()

// Create custom metrics
const homePageCounter = new client.Counter({
  name: "home_page_counter",
  help: "home page counter",
})

// Add your custom metric to the registry
register.registerMetric(homePageCounter)

client.collectDefaultMetrics({
  app: 'home',
  prefix: 'node_',
  timeout: 10000,
  gcDurationBuckets: [0.001, 0.01, 0.1, 1, 2, 5],
  register
})

// Create a route to expose metrics
app.get('/metrics', async (req, res) => {
  res.setHeader('Content-Type', register.contentType)
  res.send(await register.metrics())
})

app.get('/', (req, res) => {

  homePageCounter.inc()
  console.log(register.metrics())
  console.log("🤖", homePageCounter)

  res.send(`
  <html>
    <head>
      <meta charset="utf-8">
      <title>🌍 Home</title>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <style>
          .container { min-height: 100vh; display: flex; justify-content: center; align-items: center; text-align: center; }
          .title { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; display: block; font-weight: 300; font-size: 100px; color: #35495e; letter-spacing: 1px; }
          .subtitle { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; font-weight: 300; font-size: 42px; color: #526488; word-spacing: 5px; padding-bottom: 15px; }
          .smallertitle { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; font-weight: 300; font-size: 32px; color: #526488; word-spacing: 5px; padding-bottom: 15px; }
          .links { padding-top: 15px; }
          .smallesttitle { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; font-weight: 300; font-size: 18px; color: #526488; word-spacing: 5px; padding-bottom: 15px; }
        </style>
    </head>

    <body>
      <section class="container">
        <div>
          <h1 class="title">🏠 Home Page 📝👋</h1>
          <h2 class="subtitle">Served with 🧡 with Node.js</h2>
          <h2 class="subtitle">and 🐳 Docker</h2>
          <h3 class="smallertitle"><a href="http://localhost:8888/hello">Hello</a> | <a href="http://localhost:8888/hey">Hey</a></h3>
          <h4 class="smallesttitle">
            <a href="http://localhost:9090/targets">Prometheus</a> | 
            <a href="http://localhost:4000/d/O4gR7H5Iz/k33g_dashboard?orgId=1&refresh=5s">Grafana</a>
          </h4>

        </div>
      </section>
    </body>

  </html>
  `)

})

const http_port = 6000
app.listen(http_port, () => {
  console.log(`🌍 listening on ${http_port}`)
})
