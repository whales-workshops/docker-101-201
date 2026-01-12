#!/bin/bash

redis-cli -h redis-server -p 6379 \
HSET restaurant:1 \
name "Le Ciel" \
tags "🥦 🥘 + 🍷" \
website "http://restaurantlecielparis11.com/" \
address "17 rue Alexandre Dumas 75011 Paris" \
phone "+33 (0)6 64 66 14 16"

redis-cli -h redis-server -p 6379 \
SADD restaurants 1

redis-cli -h redis-server -p 6379 \
HSET restaurant:2 \
name "A la Renaissance"  \
tags "🥘 + 🍷" \
website "https://www.tripadvisor.com/Restaurant_Review-g187147-d786075-Reviews-A_la_Renaissance-Paris_Ile_de_France.html" \
address "87 rue de la Roquette, 75011 Paris, France" \
phone "+33 (0)1-43-79-83-09"

redis-cli -h redis-server -p 6379 \
SADD restaurants 2

redis-cli -h redis-server -p 6379 \
HSET restaurant:3 \
name "La Cave de l'Insolite"  \
tags "🥘 + 🍷" \
website "https://www.lacavedelinsolite.fr/" \
address "30 rue de la Folie Méricourt 75011 Paris" \
phone "+33 (0)1 53 36 08 33"

redis-cli -h redis-server -p 6379 \
SADD restaurants 3

redis-cli -h redis-server -p 6379 \
HSET restaurant:4 \
name "Au Nouveau Nez" \
tags "🥦 🥘 + 🍷" \
website "https://www.facebook.com/aunouveaunez/" \
address "104 Rue Saint-Maur, 75011 Paris" \
phone "+33 (0)1 43 55 02 30"

redis-cli -h redis-server -p 6379 \
SADD restaurants 4

redis-cli -h redis-server -p 6379 \
HSET restaurant:5 \
name "Le Cotte Roti" \
tags "🥦 🥘 + 🍷" \
website "https://www.tripadvisor.com/Restaurant_Review-g187147-d1337007-Reviews-Le_Cotte_Roti-Paris_Ile_de_France.html" \
address "1 rue de Cotte, 75012 Paris France" \
phone "+33 1 43 45 06 37"

redis-cli -h redis-server -p 6379 \
SADD restaurants 5

redis-cli -h redis-server -p 6379 \
HSET restaurant:6 \
name "Aime" \
tags "🥦 🥘 + 🍷" \
website "https://www.tripadvisor.com/Restaurant_Review-g187147-d12631949-Reviews-Aime-Paris_Ile_de_France.html" \
address "116 boulevard de Menilmontant, 75020 Paris France" \
phone "+33 9 50 33 27 21"

redis-cli -h redis-server -p 6379 \
SADD restaurants 6

redis-cli -h redis-server -p 6379 \
HSET restaurant:7 \
name "Biondi" \
tags "🥩 + 🍷" \
website "http://biondiparis.fr/en" \
address "118 Rue Amelot, 75011 Paris" \
phone "+33147009018"

redis-cli -h redis-server -p 6379 \
SADD restaurants 7