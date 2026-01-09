const express = require('express')
const app = express()

app.get('/', (req, res) => res.send(`
<html>
  <head>
    <meta charset="utf-8">
    <title>🌍 Hello</title>
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
        <h1 class="title">👋 Hello World 🌍😁</h1>
        <h2 class="subtitle">Served with 💜 with Node.js</h2>
        <h2 class="subtitle">and 🐳 Docker</h2>
        <h3 class="smallertitle"><a href="http://localhost:8888">🏠 Home</a></h3>
      </div>
    </section>
  </body>

</html>
`))

const http_port = 6000
app.listen(http_port, () => {
  console.log(`🌍 listening on ${http_port}`)
})