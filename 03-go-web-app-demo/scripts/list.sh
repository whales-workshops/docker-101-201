#!/bin/bash

restaurant_ids=$(redis-cli -h redis-server -p 6379 SMEMBERS restaurants)
#echo "${restaurant_ids}"

for id in $restaurant_ids; do
    restaurant_name=$(redis-cli -h redis-server -p 6379 HGET "restaurant:$id" name)
    restaurant_tags=$(redis-cli -h redis-server -p 6379 HGET "restaurant:$id" tags)

    echo "${id}. ${restaurant_name} ${restaurant_tags}"
done