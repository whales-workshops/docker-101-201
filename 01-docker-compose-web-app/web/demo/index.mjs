import { createClient } from 'redis'

const client = createClient({
    url: 'redis://redis-server:6379'
})

client.on('error', err => console.log('Redis Client Error', err))

await client.connect()

const keys = await client.keys("*")

console.log(keys)


