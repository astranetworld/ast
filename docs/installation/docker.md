# Docker

There are two ways to obtain a astDocker image:

1. [GitHub](#github)
2. [Building it from source](#building-the-docker-image)

Once you have obtained the Docker image, proceed to [Using the Docker
image](#using-the-docker-image).

> **Note**
>
> astrequires Docker Engine version 20.10.10 or higher due to [missing support](https://docs.docker.com/engine/release-notes/20.10/#201010) for the `clone3` syscall in previous versions.
## GitHub

ast docker images for both x86_64 and ARM64 machines are published with every release of ast on GitHub Container Registry.

You can obtain the latest image with:

```bash
docker pull astranet/ast
```

Or a specific version (e.g. v0.1.0) with:

```bash
docker pull astranet/ast:0.1.0
```

You can test the image with:

```bash
docker run --rm astranet/ast:0.1.0-amd64
```

If you see the latest ast release version, then you've successfully installed ast via Docker.

## Building the Docker image

To build the image from source, navigate to the root of the repository and run:

```bash
make images
```

The build will likely take several minutes. Once it's built, test it with:

```bash
docker run astranet/ast:local --version
```

## Using the Docker image

There are two ways to use the Docker image:
1. [Using Docker](#using-plain-docker)
2. [Using Docker Compose](#using-docker-compose)

### Using Plain Docker

To run ast with Docker, execute:

```
docker run -p 6060:6060 -p 61016: 61016 -p 61015: 61015/udp -v astdata:/home/ast/data astranet/ast:local --metrics --metrics.addr '0.0.0.0' 
```

The above command will create a container named ast and a named volume called astdata for data persistence. It will also expose port 61016 TCP and 61015 UDP for peering with other nodes and port 6060 for metrics.

It will use the local image ast:local. If you want to use the DockerHub Container Registry remote image, use astranet/ast with your preferred tag.

### Using Docker Compose

To run ast with Docker Compose, execute the following commands from a shell inside the root directory of this repository:

```bash
docker-compose -f docker-compose.yml up -d 
# or make up
```

The default `docker-compose.yml` file will create three containers:

- ast
- Prometheus
- Grafana


Grafana will be exposed on `localhost:3000` and accessible via default credentials (username and password is `admin`):

## Interacting with ast inside Docker

To interact with ast, you must first open a shell inside the ast container by running:

```bash
docker exec -it ast sh
```

**If ast is running with Docker Compose, replace ast with ast-ast-1 in the above command.**

Inside the ast container, refer to the [CLI docs](../cli/cli.md) documentation to interact with ast.