# Basic Commands

## Check ports listening

```bash
ss -n --all --tcp 
ss -nat
```

## Docker

1. Descargar una imagen de Docker Hub: 

```bash
docker pull <IMAGE_NAME>
```

2. Ejecutar una image

```bash
docker run -d -p <HOST_PORT:CONTAINER_PORT> <CONTAINER_NAME>
```

Equivalente

```bash
docker container run -d -p <HOST_PORT:CONTAINER_PORT> <CONTAINER_NAME>
```

3. Listar Imagenes 

```bash
docker image ls 
```

Equivalente

```bash
docker images
```

4. Crear y ejecutar un contenedor a partir de una imagen:

```bash
# docker run [opciones] nombre_del_contenedor nombre_de_la_imagen
docker run -p <HOST_PORT:CONTAINER_PORT> <CONTAINER_NAME> --name <CONTAINER_NAME> <CONTAINER-IMAGE>
```

5. Listar Contenedores en ejecución 

```bash
docker ps
```

6. Listar todos los contenedores (Incluyendo detenidos)

```bash
docker ps -a 
```

7. Detener contenedores

```bash
docker stop <ID_CONTAINER>
```

8. Iniciar contenedores

```bash
docker start <ID_CONTAINER>
```

9. Eliminar contenedores

```bash
docker rm <ID_CONTAINER> 
```

10. Construir una imagen desde un DockerFile

```bash
# docker build -t nombre_de_la_imagen [opciones y parametros] .
docker build -t <NAME> --no-cache .

# También 
docker build -t <NAME> --no-cache -f /path/to/Dockerfile
```

11. Etiquetar las imagenes 

```bash
docker tag <DOCKER_IMAGE_NAME> <NEW_TAG_NAME> 
```

12. Hacer Login en dockerhub 

```bash
docker login
```

13. Subir la imagen a dockerhub

```bash
docker push <NAME_IMAGE>
```

14. Descargar la imagen especifica

```bash
docker pull <CONTAINER_NAME:TAG>
```

## Docker Network y Volume

15. Crear una red personalizada:

```shell
docker network create nombre_de_la_red
```

16. Listar redes:

```shell
docker network ls
```

17. Crear un volumen:

```shell
docker volume create nombre_del_volumen
```

18. Listar volúmenes:

```shell
docker volume ls
```

## Interactuando con Contenedores

19. Ejecutar un comando en un contenedor en ejecución:

```shell
#docker exec -it nombre_del_contenedor comando
docker exec -it <NAME-CONTAINER> <COMMAND>
```

    20. Copiar archivos entre el host y un contenedor:

```shell
docker cp archivo.txt <CONTAINER_NAME>:/destination/path
```

## Docker Compose

21. Iniciar contenedores definidos en un archivo `docker-compose.yml`:

```shell
docker compose up
```

22. Escalar servicios en Docker Compose:

```shell
docker compose scale servicio=num_instancias
```

23. Detener y eliminar contenedores definidos en Docker Compose:

```shell
docker compose down
```

24. Ver logs de servicios en Docker Compose:

```shell
docker compose logs <SERVICES_NAME>
```

## Administración de Redes y Volúmenes

25. Eliminar una red:

```shell
docker network rm <NETWORK_NAME>
```

26. Eliminar un volumen:

```shell
docker volume rm <VOLUME_NAME>
```

27. Inspeccionar una red:

```shell
docker network inspect <NETWORK_NAME>
```

### Comandos para Información y Estadísticas

28. Ver detalles de un contenedor:

```shell
docker inspect <CONTAINER_NAME>
```

29. Obtener estadísticas de uso de recursos de un contenedor:

```shell
docker stats <CONTAINER_NAME>
```

30. Obtener estadísticas de uso de recursos de todos los contenedores:

```shell
docker stats
```
