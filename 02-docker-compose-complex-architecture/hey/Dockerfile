FROM node:20.11.1-slim
# docker build -t hey .
# docker run -p 6200:6000 -ti --rm --init hey
# http://localhost:6200
WORKDIR /app
COPY package.json index.js ./
RUN npm install

EXPOSE 6000
CMD ["index.js"]