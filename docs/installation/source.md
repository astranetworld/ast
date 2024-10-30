# Build from Source

You can build ast Linux, macOS, Windows, and Windows WSL2.

> **Note**
>
> ast does not work on Windows WSL1.

## Dependencies

First, Instructions for installing Go are available at the [Go installation page](https://golang.org/doc/install) and necessary bundles can be downloaded from the [Go download page](https://golang.org/dl/).
Note: Please download version 1.2, as some libraries do not yet support higher versions.
## Build ast

With go and the dependencies installed, you're ready to build ast. First, clone the repository:

```plaintext
git clone https://github.com/WeAreAmaze/ast
cd ast
```

Then, install ast into your PATH directly via:

```plaintext
go build -o ./build/bin ./cmd/ast
# or make ast (Linux, MacOS)
```
Now, via the command line, the binary will be accessible as ast and resides in ./build/bin folder.

Compilation may take around 1 minute. If `ast --help` displays the [command-line documentation](../cli/cli.md).

If you run into any issues, please check the [Troubleshooting](#troubleshooting) section, or reach out to us on [Telegram](https://t.me/N42t).

## Update ast

You can update ast to a specific version by running the commands below.

The ast directory will be the location you cloned ast to during the installation process.

${VERSION} will be the version you wish to build in the format vX.X.X.

```bash
cd ast
git fetch
git checkout ${VERSION}
go build -o ./build/bin ./cmd/ast
# or make ast  (Linux, MacOS)
```

