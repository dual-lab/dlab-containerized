![dlab container](https://github.com/dual-lab/dlab-containerized)

# OpenSSL on Alpine Linux

Small image base on alpine linux to use openssl

- opnessl 1.1.1g - Alpine 3.12

## Installation

`docker/podman pull dlab/openssl:tag`

## Examples

### OpenSSL REPL

`docker/podman run -it dlab/openssl`

### Generate keys

`docker/podman run -it --rm dlab/openssl opnessl enc -aes-256-cbc -k <pass> -P -md sha256 -nosalt -pbkdf2`

### AES encrypt

`target=<directory containg the file to encrypt>`
`docker/podman run -it --rm -v $target:/export dlab/openssl aes-256-cbc -e -in <file> -out <file_enc> -K <key> -iv <iv>`

Where _key_ and _iv_ are generated during the keys step.

Read the OpenSSL [documentation](https://www.openssl.org/docs/) for further informations.

### AES decrypt

`target=<directory containg the file to encrypt>`
`docker/podman run -it --rm -v $target:/export dlab/openssl aes-256-cbc -d -in <file_enc> -out <file> -K <key> -iv <iv>`

Where _key_ and _iv_ are generated during the keys step.

Read the OpenSSL [documentation](https://www.openssl.org/docs/) for further informations.

### interactive shell

`docker/podman run -it --entrypoint /bin/ash dlab/openssl`


