# Micro Service Example
This example uses [micro service framework](https://github.com/0-haha/microsvs) and graphql.
The demonstrates sql operations for a user such as insertion, deletion, update, and query.
In addition, the example includes an example for security check before calling graphql.
Using this template, one can quickly develop a backend server.

### How to run
- Start docker compose
    ```bash
    docker-compose up
    ```

- Initial Mongo
    ```bash
    # login mongo
    mongo --port 27017 --host=localhost --authenticationDatabase=admin \
        -p password --username test
    ```

    ```bash
    # setup replication with single master
    rs.initiate({_id: "rs0", members: [{_id: 0, host: "localhost:27017"}] })
    ```
