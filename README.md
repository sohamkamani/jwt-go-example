# JWT Go example

Example application that implements JWT based authentication. Read the blog post [here](https://sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/)

To run this application, build and run the Go binary:

```sh
go build
./jwt-go-example
```

Now, using any HTTP client with support for cookies (like [Postman](https://www.getpostman.com/apps), or your web browser) make a sign-in request with the appropriate credentials:

```
POST http://localhost:8000/signin

{"username":"user1","password":"password1"}
```

You can now try hitting the welcome route from the same client to get the welcome message:

```
GET http://localhost:8000/welcome
```

Hit the refresh route, and then inspect the clients cookies to see the new value of the `token` cookie:

```
POST http://localhost:8000/refresh
```

You can find the working source code for this example [here](https://github.com/sohamkamani/jwt-go-example).