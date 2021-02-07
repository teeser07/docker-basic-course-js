---
theme : "black"
transition: "slide"
highlightTheme: "monokai"
slideNumber: false
logoImg: "images/logo-igg.png"
title: "Docker basic course (JS Edition) - Part 2"
---

# Checkpoint

- [x] Container
- [x] Image
- [x] Dockerfile
- [x] Registry
- [x] Auto build with Github (Bonus!)

---

# Checkpoint

- [ ] Basic docker volume
- [ ] Docker network
- [ ] Docker multistage build
- [ ] Real-world example project
- [ ] Deploy container to Heroku (Bonus!)

---

## Lab 9. Create upload script image

```
docker build -t app-upload:0.1 .
```

```
docker run -d -p 3000:3000 app-upload:0.1
```

```
docker run -P app-upload:0.1
```

---

### Volume mount

![Image](images/02.06.png)

---

### 1. tmpfs mount

![Image](https://docs.docker.com/storage/images/types-of-mounts-tmpfs.png)

---

### 2. Bind mount

![Image](https://docs.docker.com/storage/images/types-of-mounts-bind.png)

```
docker run -d -p 3000:3000 -v $(pwd)/uploads:/home/node/uploads app-upload:0.1
```

- -v is a volume mounting `HOST DIRECTORY`:`CONTAINER DIRECTORY`.

---

### 3. Volume

![Image](https://docs.docker.com/storage/images/types-of-mounts-volume.png)

```
docker volume create upload-app
```

```
docker volume ls
```

```
docker volume inspect {{volumeId}}
```

```
docker run -d -p 3000:3000 -v upload-app:/home/node/uploads app-upload:0.1
```

```
docker volume rm upload-app
```

---

# Checkpoint

- [x] Basic docker volume
- [ ] Docker network
- [ ] Docker multistage build
- [ ] Real-world example project
- [ ] Deploy container to Heroku (Bonus!)


---

# Networking

```
docker network ls
```

### Type of network

1. Bridge Network
2. None Network
3. Host Network
4. Overlay Network

---

![Image](https://miro.medium.com/max/2656/1*WKiEgPXO8XXppoqgr7ZVQA.png)

https://morioh.com/p/07e61c20c234

---

### 1. Bridge network

```
docker run -itd --name network-app-01 alpine
```

```
docker run -itd --name network-app-02 alpine
```

```
docker network inspect bridge
```

---

![Image](https://miro.medium.com/max/1182/1*DBNLOEZJN3M-Q0LXOuSnew.png)
https://medium.com/@somprasongd/docker-networking-59b6637de3df

---

#### Check ip of container

```
docker exec -it network-app-01 ifconfig
```

```
docker exec -it network-app-02 ifconfig
```

```
docker exec -it network-app-02 ping {ip}
```

---

![Image](https://miro.medium.com/max/2080/1*DQ7oHTHIgHOvfTqk58NR-A.png)

https://morioh.com/p/07e61c20c234

---
#### Linking by container name

```
docker rm network-app-01 -f
```

```
docker rm network-app-02 -f
```

```
docker run -itd --name network-app-01 alpine
```

```
docker run -itd --name network-app-02 --link network-app-01 alpine
```

```
docker exec -it network-app-02 ping network-app-01
```

---

#### Create own network

```
docker rm network-app-01 -f
```

```
docker rm network-app-02 -f
```

```
docker network create --driver bridge app-network
```

```
docker network ls
```

```
docker run -itd --name network-app-01 --network app-network alpine
```

```
docker run -itd --name network-app-02 --network app-network alpine
```

```
docker exec -it network-app-02 ping network-app-01
```


---

### 2. None network

```
docker run -itd --name app-none-network --network none alpine
```

```
docker exec -it app-none-network ping 1.1.1.1
```

```
docker exec -it network-app-02 ping 1.1.1.1
```

---

### 3. Host network

```
docker run -itd --name app-host-network --network host alpine
```

```
docker network inspect host
```

```
docker exec -it app-host-network ifconfig
```

---

### 4. Overlay network

![Image](https://miro.medium.com/max/1284/1*rk_5tzWwrSKcHCKwhYMH0Q.png)

---

## Lab 10. Golang api

```
docker build ???
```

```
docker run ???
```

```
docker images
```

---

### Let's build multistage

```
docker build -f Dockerfile.multi -t app-go:0.2 .
```

---

![Image](images/03.04.png)

---


![Image](images/03.05.png)

---


![Image](images/03.06.png)

---


![Image](images/03.07.png)

---

# Checkpoint

- [x] Basic docker volume
- [x] Docker network
- [x] Docker multistage build
- [ ] Real-world example project
- [ ] Deploy container to Heroku (Bonus!)

---

## Real-world example project

![Image](https://miro.medium.com/max/4800/0*6T9WJA6nqmF-t2r1.gif)

[:Ref](https://medium.com/bb-tutorials-and-thoughts/dockerizing-vue-app-with-nodejs-backend-33645f0f50ec)

---

## [Github](https://github.com/nitipatl/vuejs-nodejs-example)

```
git clone https://github.com/nitipatl/vuejs-nodejs-example.git
```

---

```
docker build -t vue-node-image:0.1 .
```

---

```
docker run -it -p  3080:3080 --name vue-node-ui vue-node-image:0.1
```

---

# Checkpoint

- [x] Basic docker volume
- [x] Docker network
- [x] Docker multistage build
- [x] Real-world example project
- [ ] Deploy container to Heroku (Bonus!)

---

## Deploy container to Heroku (Bonus!)

```
heroku login
```

```
heroku container:login
```

```
heroku create
```

---

### Push image to Heroku

```
heroku container:push web --app warm-dusk-59086
```

---

### Release app / run container 

```
heroku container:release web --app warm-dusk-59086
```

```
heroku open --app warm-dusk-59086
```

```
heroku logs --tail --app warm-dusk-59086
```

---
# Checkpoint

- [x] Basic docker volume
- [x] Docker network
- [x] Docker multistage build
- [x] Real-world example project
- [x] Deploy container to Heroku (Bonus!)
- [ ] [Auto deploy with Github (Bonus!)](https://dev.to/heroku/deploying-to-heroku-from-github-actions-29ej)

---

## 12 Factors app

[Link](images/03.08.png)

---