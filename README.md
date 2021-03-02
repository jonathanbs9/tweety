# Tweety

![N|Solid](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSV4Paq-ddrzOLM119LMfJTnkccjODj2k7RWw&usqp=CAU)


*Tweety es una ap red social similar a Twitter, donde el usuario se podrá registrar, loguearse, seguir a un determinado usuario y a la vez podrá revisar y chequear los tweets de la gente que está siguiendo.*

### Descripción && Features

*Perfil de usuario*

- Podremos ver nuestro propio perfil donde tendremos datos del usuario, avatar, bañera y un los tweets del usuario y podremos visitar el perfil de otros usuarios, pero solo se podrá editar el perfil de usuario de uno mismo.

*Sistema de Followers*

- Podremos seguir Y dejar de seguir a otros usuarios que estén registrados en la aplicación y tendremos una lista de usuarios para ver a quien estamos siguiendo en todo momento.

*Sistema de Tweets*

- Podremos mandar tweets en cualquier momento y desde cualquier página de nuestra aplicación y cuando visitemos el perfil de otro usuario podremos ver todos sus tweets.

*Buscador de usuarios*

- Podremos buscar usuarios por su nombre y filtrar la búsqueda entre usuarios que no estamos siguiendo o usuarios que estamos siguiendo.

*Feed de Tweets*

- Tendremos una pagina donde podremos ver los últimos tweets que han enviado los usuarios que estamos siguiendo.

### Tech's

Dillinger uses a number of open source projects to work properly:

* [Golang] - Backend
* [Git]  -  Control de versiones
* [Github] - Hosting de repositorios
* [MongoDB] -  Sistema de base de datos NoSQL. MongoDB guarda estructuras de datos BSON (una especificación similar a JSON) con un esquema dinámico, haciendo que la integración de los datos en ciertas aplicaciones sea más fácil y rápida.
* [Postman] -  Herramienta que se utiliza, para el testing de API REST.
* [JWT] - Creación de tokens de acceso que permiten la propagación de identidad y privilegios.
* [bCrypt] -  Encriptación de  passwords

### Run

Una vez descargado el repo de github

```sh
$ go build main.go
```

```sh
$ go run main.go
```

### Postman

```sh
localhost:8080
```
### Api documentation

- POST /signup
{
    "first_name" : string,
    "last_name" : string,
    "email": string,
    "password": string,
    "date_birth": "1992-06-21T23:06:00Z"
}

- POST /login
{
    "email": string,
    "password" : string   
}

- GET /profile?id={id}

- PUT /modifyprofile
{
    "first_name" : string,
    "last_name" : string,
    "date_birth": "1986-06-23T20:25:00Z",
    "location" : string,
    "biography": string ,
    "website": string
}

- POST /savetweet 
{
    "message" : string
}

- GET /tweets?id={id}&page={pageid}

- DELETE /deletetweet?id={id}

- POST /uploadavatar
{ 
  form-data:
  key : "avatar" => value: "avatar.png"
}

- POST /uploadbanner
{ 
  form-data:
  key : "banner" => value: "banner.png"
}
- POST /relation?id={id}

- DELETE /relation?id={id}

- GET /relation?id={id}

- GET /listusers?page={pageid}&type=new&search={string}

- GET /followtweets?page={pageid}

### Todo's

- Swagger documentaction
- Unit testing
- Heroku deployment

#
