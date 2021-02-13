## install Express js


#### 1. init npm
```
npm init
```

#### 2. init express js

```
npm install express --save
```


#### 3. create index.js
```
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  res.send('Hello World!')
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
```

#### 4. add script

```
"scripts": {
    ...
    "start": "node index.js",
    ...
}
```

#### 5. run express js

```
npm run start
```

#### 6. create Dockerfile
```
FROM node:12.18-alpine

RUN mkdir -p /usr/app
WORKDIR /usr/app

COPY ./package*.json ./

RUN npm install --quiet

COPY . .
```

#### 7. create docker-compose.yml (project folder)
```
version: '3'
services:
  api:
    build:
      context: ./api
    command: ['/bin/sh', '-c', 'npm install --quiet && npm run start']
    volumes:
      - ./api:/usr/app/
      - /api/usr/app/node_modules
    ports:
      - "3000:3000"
    networks:
      - custom-network
networks:
  custom-network:
    external: true
```

#### 8. run docker compose
```
docker-compose up
```

## install Nuxt js


#### 1. init nuxt js

```
npm init nuxt-app frontend
```

#### 2. add env to nuxt.config.js
```
env: {
  PORT: process.env.PORT,
  HOST: process.env.HOST,
},
```

#### 3. create .env
```
PORT=3001
HOST=0.0.0.0
```

#### 4. run nuxt js
```
npm run dev
```

#### 5. create Dockerfile
<!-- ```
FROM node:12.18-alpine

RUN mkdir -p /usr/app
WORKDIR /usr/app

COPY ./package*.json ./

RUN yarn --quiet

COPY . .
``` -->

#### 6. edit docker-compose.yml
<!-- ```
version: '3'
services:
  api:
    build:
      context: ./api
    command: ['/bin/sh', '-c', 'npm install --quiet && npm run start']
    volumes:
      - ./api:/usr/app/
      - /api/usr/app/node_modules
    ports:
      - "3000:3000"
    networks:
      - custom-network
  frontend:
    build:
      context: ./frontend
    volumes:
      - ./frontend:/usr/app/
      - /frontend/usr/app/node_modules
    command: ['/bin/sh', '-c', 'yarn --quiet && yarn dev']
    ports:
      - "3001:3001"
    env_file:
      - ./frontend/.env
    networks:
      - custom-network
networks:
  custom-network:
    external: true
``` -->

#### 7. run docker compose
```
docker-compose up
```










