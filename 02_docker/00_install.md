# Instalación de docker (Oracle Linux 8)

Esta sección ha sido tomada y adaptada según la guía de instalación de 

[https://docs.docker.com/engine/install/rhel/](docker)

Eliminar dependecias previas y/o paquetes que puedan generar errores 

```bash
sudo dnf remove docker \
 docker-client \
 docker-client-latest \
 docker-common \
 docker-latest \
 docker-latest-logrotate \
 docker-logrotate \
 docker-engine \
 podman \
 runc
```

Agregar el repositorio con los paquetes 

```bash
sudo dnf -y install dnf-plugins-core
sudo dnf config-manager --add-repo https://download.docker.com/linux/rhel/docker-ce.repo
sudo dnf update
```

Instalar dependencias y paquetes 

```bash
sudo dnf install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

Dejar el servicio como enable con el systemd 

```bash
sudo systemctl enable --now docker
```

Habilitar al usuario actual para usar el comando `docker` sin permisos de ususario 

```bash
sudo usermod -aG docker $USER
newgrp docker
```


