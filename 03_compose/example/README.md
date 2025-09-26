# Ejercicio con Docker Compose

Partes de un archivo .yaml: 

```bash
services: # servicios que vamos a crear dentro de nuestro proyectos 
    service_name: # Nombre del servicio especifico puede ser:
                  # Backend, Fronted un Proxy, una base de datos

        build: # como vamos a contruirla, puede ser con una imagen 
               # de docker o desde un docker file 
        networks: # definimos la red a la que vamos a cenectar nuestro 
                  # servicio 
        ports: # definimos puertos internos y externos del servicio 

        depends_on: # definimos si nuestro servicio depende de otro 

        deploy: # definimos caracteristicas para el desplieque de 
                # nuestro servicio, como recursos 
                # y replicas 

        volumes: # definimos un espacio de disco donde vamos a almacenar 
                 # nuestra información 
        enviroment: # variables de entorno del servicio
```

## Definamos nuestras imagenes con los Dockerfile

Definamos la imagen del proxy 

```bash
FROM nginx:mainline-alpine
COPY init.sh /init.sh
RUN apk update && apk add --no-cache openssl
COPY ./nginx.conf /etc/nginx/nginx.conf
ENTRYPOINT ["/bin/sh", "init.sh"]
CMD ["nginx", "-g", "daemon off;"]
```

Definimos la imagen del backend 

```bash
FROM golang:1.25.1-alpine
WORKDIR /app
COPY . .
RUN go mod init main && go mod tidy
RUN go build -o server .
CMD ["/app/server"]
```

## Definamos como administrar nuestra aplicación con Docker compose

Creación del proxy server: 

```bash
  proxy:
    container_name: nginx_proxy 
    build: 
      context: proxy
      dockerfile: Dockerfile
    restart: always
    networks:
      - app
    ports:
      - 80:80
    links:
      - backend
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 500M
```

Creación de la base de datos: 

```bash
  db:
    container_name: postgres_db 
    image: postgres
    ports:
      - "5432:5432"
    restart: always
    environment:
      POSTGRES_PASSWORD: password
      POSTGRES_USER: postgres
      APP_DB_USER: postgres
      APP_DB_PASS: password
      APP_DB_NAME: postgres
    volumes:
      - ./db:/docker-entrypoint-initdb.d/
    networks:
      - app
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 500M
```

Definición del backed: 

```bash
  backend:
    build: 
      context: golang
      dockerfile: Dockerfile
    restart: always
    expose:
      - "8080"
    depends_on:
      - db
    networks:
      - app
    environment:
      RDS_HOSTNAME: db
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 500M
      mode: replicated
      replicas: 4
      placement:
        max_replicas_per_node: 1
      endpoint_mode: dnsrr
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure
```

Definición de la red: 

```bash
networks:
  app:
    ipam:
      driver: default
      config:
        - subnet: 192.168.200.0/24
```
