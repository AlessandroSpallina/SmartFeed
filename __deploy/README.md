Here you can find all needed steps to deploy this project on docker-compose

## Development environment

1. Install docker and docker-compose
2. Run:

```sh

    sh build-dev.sh
    docker-compose -f docker-compose-dev.yml up
```

## Test with postman!
1. Install [Postman](https://www.postman.com/)
2. Get the postman method collection [HERE](https://www.getpostman.com/collections/2be9d338f6a5b032acb4)
3. Import the collection in postman
4. Choose methods to test by GUI.

## Test with curl (not working right now)!
```

  # Ping
  curl --location --request GET 'http://localhost:3000/ping' --header 'Content-Type: application/json'

  # Register new user
  curl --location --request POST 'http://localhost:3000/u/register' --header 'Content-Type: application/json' --form 'username=sk3' --form 'password=la'


  ----------------
  # Show uploaded metadatas
  curl --user user@a.a:user --location --request GET 'http://osa.localhost/fms/' --header 'Content-Type: application/json'

  # Upload file metadata
  curl --user user@a.a:user --location --request POST 'http://osa.localhost/fms/' --header 'Content-Type: application/json' --data-raw '{"filename" : "README.md","author" : "localhost gang"}'

  # Upload the file (be careful, it's a POST /{previously-returned-id}
  curl --user user@a.a:user --location --request POST 'http://osa.localhost/fms/2' --form 'file=@./README.md'

  # Get uploaded file url (be careful, it's a GET/{previously-returned-id}
  curl --user user@a.a:user --location --request GET 'http://osa.localhost/fms/2' --header 'Content-Type: application/json'

```
