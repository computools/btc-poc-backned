# BTC backend

## Additional programs

1. [Taskfile](https://taskfile.dev/installation/) (Optional)
2. docker-compose or podman-compose
3. [Postman](https://www.postman.com/downloads/)

## How to run

1. Run `go mod download` command to download dependencies

   ```shell
   go mod download
   ```

2. Copy the [.env.example](.env.example) file to `.env` and change variables to what you need

   ```shell
   cp .env.example .env && nano .env
   ```

3. Start the postgres database

   > **<span style="color:#79b6c9">ⓘ NOTE:</span>** If you're starting the program by running the taskfile command, you
   can skip this step because DB will start up automatically, and it does not need additional configuration.

   ```shell
   export POSTGRES_PORT=5432 POSTGRES_USER=postgres POSTGRES_PASSWORD=postgres POSTGRES_DB=btc
   cd containers && docker-compose -f database.yml up -d && cd ..
   ```

   Or use next taskfile command:
   ```shell
   task db
   ```

4. Start the keycloak

   ```shell
   export KEYCLOAK_PORT=1141 KEYCLOAK_ADMIN=user KEYCLOAK_ADMIN_PASSWORD=bitnami
   cd containers && docker-compose -f keycloak.yaml up -d && cd ..
   ```

   Or use next taskfile command:
   ```shell
   task kc
   ```

5. Configure keycloak and paste KEYCLOAK_CLIENT_ID and KEYCLOAK_CLIENT_CREDENTIALS variables to [.env](.env) file
6. Run the following command to start server:

    ```shell
    go run main.go
    ```

   Or you can run `Run` configuration if you are using `Goland`

   Or you can run the following taskfile command:
    ```shell
    task run
    ```

7. Import [`BTC` collection](./requests/BTC.postman_collection.json) and [
   `BTC-Local` environment](./requests/BTC-Local.postman_environment.json) into `Postman` program

8. Send request throw `Postman`
