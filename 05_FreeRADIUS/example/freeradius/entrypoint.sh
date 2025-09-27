#!/bin/bash

# Esperar a que PostgreSQL esté disponible
# (Aunque docker-compose 'depends_on' ayuda, esto es una doble verificación)
echo "Waiting for PostgreSQL..."
while ! nc -z $POSTGRES_HOST $POSTGRES_PORT; do
  sleep 1
done
echo "PostgreSQL started"

# Configurar el módulo SQL de FreeRADIUS
echo "Configuring FreeRADIUS SQL module..."
sed -i \
    -e "s/driver = \"rlm_sql_null\"/driver = \"rlm_sql_postgresql\"/"
    -e "s/dialect = \"sqlite\"/dialect = \"postgresql\"/"
    -e "s/server = \"localhost\"/server = \"${POSTGRES_HOST}\"/"
    -e "s/port = 5432/port = ${POSTGRES_PORT}/"
    -e "s/login = \"radius\"/login = \"${POSTGRES_USER}\"/"
    -e "s/password = \"radpass\"/password = \"${POSTGRES_PASSWORD}\"/"
    -e "s/radius_db = \"radius\"/radius_db = \"${POSTGRES_DB}\"/"
    /etc/raddb/mods-available/sql

# Habilitar el módulo SQL creando un enlace simbólico
ln -sf /etc/raddb/mods-available/sql /etc/raddb/mods-enabled/sql

# Descomentar la llamada al módulo 'sql' en el sitio por defecto para que se use en la autenticación y accounting
sed -i -e 's/#[[:space:]]*sql/\tsql/' /etc/raddb/sites-available/default

# Ejecutar el comando original pasado al contenedor (ej. freeradius -X)
exec "$@"
