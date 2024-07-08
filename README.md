## Burgers API

(Test assessment for a Golang Developer Position)

Create API (backend code + db):

* Search Burger by name
* List all Burgers by first letter
* Search ingredient by name
* Lookup full Burger details by id
* Lookup Ingredient by ID
* Lookup a random Burger
* Search by ingredient
* Creating a new Burger.

Once deployed, refer to http://{host}/swagger/index.html for details.

## Steps to build and run Back-End

Build and Run on Windows x64

```shell
cd app; if ($?) { go build -o ./../bin/cli/burgers-api.exe cmd/main/main.go}
```

```shell
cd ./../bin/cli; if ($?) { ./burgers-api.exe }
```

Build and Run on Linux

```shell
cd app/ && env GOOS=linux GOARCH=amd64 && go build -o ./../bin/cli/burgers-api cmd/main/main.go
```

```shell
cd ./../bin/cli && ./burgers-api
```

Stop the app with ^C.

## Dockerized app

Run the App

```shell
docker-compose -f ./docker-compose.yml  up -d
```

Stop the App

```shell
docker-compose -f ./docker-compose.yml down
```

## Generate Open API Docs

Windows:

```shell
swag init -g cmd/main/main.go --parseInternal --parseDependency --parseDepth 1  && swag fmt -g cmd/main/main.go
```

Linux:

```shell
~/go/bin/swag init -g cmd/main/main.go --parseInternal --parseDependency --parseDepth 1  && \ 
~/go/bin/swag fmt -g cmd/main/main.go 

## Steps to build Front-End

```shell
cd ./front-end/
```

```shell
npm install
```

for Development:

```shell
npm run dev
```

for Production:

```shell
npm run prod
```
