# copee los archivos en su espacio de trabajo
ADD . /go/src/github.com/couponmanager

# copee el archivo de configuración
ADD . /go/config/config.json

# agrefue la ruta local a las variables de entorno
GOPATH /go/src/github.com/couponmanager

# obtenga godep para administrar las dependencias
go get github.com/tools/godep

# Instale godep dependecias
godep restore 

# corra el siguiente comando para obtener el archivo ejecutable
go install github.com/couponmanager

# ejecute dicho archivo
/go/bin/taskmanager

# de acuerdo a el archivo de configuracion el servicio queda expuesto en el puerto 8080.
EXPOSE 8080


carpetas:

-common: implementa algunas funciones de utilidad y proporciona lógica de inicialización para la aplicacion
-controllers: implementa los datos de los controladores de aplicaciones 
-service: implementa la logica de negocio
-models: describe el modelo de datos de la aplicacion
-routers: implementa los enrutadores de solicitud HTTP para la API RESTful