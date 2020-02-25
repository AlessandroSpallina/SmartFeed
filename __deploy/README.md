Here you can find all needed steps to deploy this project on docker-compose

## Development environment

Run

```sh

    sh build-dev.sh
    docker-compose -f docker-compose-dev.yml up
```

## Test with curl!
```
  # Register new user
  curl --location --request POST 'http://localhost:3000/u/register' --header 'Content-Type: application/json' --data-raw '{"username":"ale","password":"spa"}'


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
