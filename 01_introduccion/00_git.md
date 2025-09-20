# Comandos

1. Para crear un repositorio desde 0: 

```bash
echo "# CONTENT" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin <URL>
git push -u origin main
```

2. Para subir cambios por primera vez: 

```bash
git remote add origin <URL>
git branch -M main
git push -u origin main
```

3. Comandos basicos 

```bash
git add .                     # también git add --all 
git commit -m "COMENTARIO"    # Marcar el comentario 
git push                      # Subir los cambios 
git pull                      # Descargar cambios del repositorio 
git status                    # Validar el estado del repositorio 
```


