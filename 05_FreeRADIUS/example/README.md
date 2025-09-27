# Configuración del FreeRadius con Docker Compose

0. Instalación de herramientas parapruebas: 

```bash
# Paquete de utilidades, para validar freeradiuos 
dnf install freeradius-utils

# Comando para validar auntenticación
# 1 = PAP (Password Authentication Protocol)
# radtest steve testing 192.168.1.7:1812 1 testing123
radtest <USER> <PASSWORD> <SERVER_IP>:1812 <PROTOCOL_TYPE> <SHARED_SECRED>
```

1. Creación del Dockerfile del freeradius 

```bash
FROM freeradius/freeradius-server:latest
RUN apt-get update && apt-get install -y freeradius-postgresql
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
CMD ["freeradius", "-X"]
```

2. Crear el archivo clients.conf: 

```bash
client testing {
    ipaddr      = 192.168.1.254   # IP del dispositivo que hará las peticiones
    secret      = testing123      # Secreto compartido entre el servidor y el cliente
    shortname   = home-router
}

# Cliente para pruebas locales con radtest
client localhost {
    ipaddr      = 127.0.0.1
    secret      = testing123
}

# Cliente para la red interna de Docker
client docker_network {
    ipaddr      = 172.0.0.0/8
    secret      = testing123
}

# Puedes definir clientes por subred
client localnet {
    ipaddr      = 10.0.0.0/8
    secret      = my_local_secret
}

# Cliente para permitir cualquier IP (SOLO PARA PRUEBAS)
# ADVERTENCIA: Inseguro para producción. Reemplazar con la IP específica del cliente.
client any_ip {
    ipaddr      = 0.0.0.0/0
    secret      = testing123
}

```

3. Creación del archivo authorize

```bash
# El formato es: nombre_usuario Atributo-de-chequeo := Valor
steve   Cleartext-Password := "testing"
        Service-Type = Framed-User,
        Framed-Protocol = PPP,
	# 172.0.0.0/8
        # Framed-IP-Address = 192.168.1.10,
        Framed-IP-Address = 172.0.0.0,
        Framed-IP-Netmask = 255.0.0.0

# Otro usuario más simple
user1   Cleartext-Password := "password123"
```

NOTAS: 

Los archivos de inicialización de las bases de datos fueron optenidos de: 

```bash
docker run --rm freeradius-temp cat
     /etc/freeradius/3.0/mods-config/sql/main/postgresql/schema.sql >
     postgres/init/init.sql
```


