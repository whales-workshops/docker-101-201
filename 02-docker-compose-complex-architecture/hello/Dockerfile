FROM node:20.11.1-slim
# docker build -t hello .
# docker run -p 6100:6000 -ti --rm --init hello
# http://localhost:6100

WORKDIR /app
COPY package.json index.js ./
RUN npm install

EXPOSE 6000
CMD ["index.js"]