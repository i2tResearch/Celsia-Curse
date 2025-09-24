# Explore Docker Images

```shell
docker run -it --rm wernight/funbox bash
```

# Ejemplo CMATRIX

```bash
docker pull alpine
```

Crearemos el Dockerfile 

```bash
vim Dockerfile 
```

Dentro de este comenzaremos la edición del archivo 

```docker
FROM alpine 
```

Construimos la imagen 

```bash
docker build -t njse22/cmatrix .
```

Ejecutamos la imagen 

```bash
docker run --rm -it njse22/cmatrix
```

Dentro del contenedor 

```bash
hostname

apk update

git clone https://github.com/spurin/cmatrix.git

apk update

apk add git

git clone https://github.com/spurin/cmatrix.git

cd cmatrix/

ls -l

autoreconf -i

apk add autoconf

autoreconf -i

apk add automake

autoreconf -i

echo $?

./configure LDFLAGS="-static"

apk add alpine-sdk

./configure LDFLAGS="-static"

apk add ncurses-dev ncurses-static

mkdir -p /usr/lib/kbd/consolefonts /usr/share/consolefonts

./configure LDFLAGS="-static"

make

ls -l ./cmatrix

./cmatrix
```

Despues de validar podemos crear nuestro dockerfile desde el historial de comandos 

```bash
FROM alpine

WORKDIR /cmatrix

RUN apk update
RUN apk add git
RUN git clone https://github.com/spurin/cmatrix.git .
RUN apk add autoconf
RUN apk add automake
RUN autoreconf -i
RUN apk add alpine-sdk
RUN apk add ncurses-dev ncurses-static
RUN mkdir -p /usr/lib/kbd/consolefonts /usr/share/consolefonts
RUN ./configure LDFLAGS="-static"
RUN make
CMD ["./cmatrix"]
```

Luego podemos correr la imagen 

```bash
docker run --rm -it njse22/cmatrix
```

## Ejemplo publicar en docker hub

Validar el login: 

```bash
 docker login -u njse22
```

```bash
docker buildx build --no-cache --platform linux/amd64 . -t njse22/cmatrix --push

## Nota para arm usar linux/arm64/v8
```
