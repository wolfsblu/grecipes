## Build Steps
1. Generate the API server based on the OpenAPI spec with
    ```
    go generate
    ```
2. Build the frontend application 
    ```
    npm --prefix app install
    npm --prefix app run build
    ```
3. Build the project
    ```
    go build -o ./tmp/main
    ```

## Running the Application
Assuming you want to run the binary directly from the repository:
1. Apply any pending database migrations
    ```shell
    atlas migrate apply --dir "file://db/migrations" --url "sqlite://tmp/db.sqlite" 
    ```
2. Start the server
    ```
    DB_PATH=tmp/db.sqlite ./tmp/main
    ```
3. Open the frontend at http://localhost:8080
