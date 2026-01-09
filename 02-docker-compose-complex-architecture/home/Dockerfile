FROM node:20.11.1-slim
# docker build -t home .
# docker run -p 6500:6000 -ti --rm --init home
# http://localhost:6500

WORKDIR /app
COPY package.json index.js ./
RUN npm install

EXPOSE 6000
CMD ["index.js"]