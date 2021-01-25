# node-vs-golang-performance

This project contains two Lambda functions that perform a merge sort 
algorithm one in NodeJS and the other is in GoLang to represent a performance
comparison between the two runtimes.

## Using this example:

- Use AWS CLI to create a resource bucket.
    ```bash
    $ S3_BUCKET=<your bucket name> make init
    ```

- Use SAM CLI to build the application.
    ```bash
    $ make build
    ```

- Use SAM CLI to package the application.
    ```bash
    $ S3_BUCKET=<your bucket name> make package
    ```

- Use SAM CLI to deploy the application.
    ```bash
    $ STACK_NAME=<your stack name> S3_BUCKET=<your bucket name> make deploy
    ```

- Use AWS CLI to delete the application.
    ```bash
    $ STACK_NAME=<your stack name> S3_BUCKET=<your bucket name> make delete
    ```
