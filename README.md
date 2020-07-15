# test-rig-example

Example of the `sysl test-rig` feature.

## Usage

```bash
$ make
```

Make generates sources in `/cmd` and `/gen`, and a `docker-compose.yml` in the root directory. Running `docker-compose up` builds the binaries and runs containers for both the PostgreSQL database backend and `dbfront`, the database frontend.
