<img src=".github/EnCloud_RGB-03.png" alt="180Protocol Logo" width="30%" height="30%"/>

encloud is the simplest way to onboard sensitive data to Web3. encloud fills a key gap in the decentralized web - privacy. Most decentralized
clouds like Filecoin, don't offer privacy natively rendering them unuseful for sensitive data.

**encloud solves this.**

encloud has three offerings

- The [encloud CLI](cmd/cli/) enables users to easily on-board sensitive data to the Filecoin network.
- The [encloud Desktop Application](cmd/web/) is a lightweight GUI built on top of the CLI and can be downloaded via the [encloud website](https://encloud.tech/).
  - [Readme](cmd/web/README.md)
  - [How to guide](cmd/web/HOWTO.md)
- The [encloud API](cmd/api/) is a REST API that can be used to integrate encloud into your web applications.

# Encloud CLI

encloud lets users manage encryption keys and onboard their encrypted data to Filecoin

- Generate their ECIES/RSA encryption keys
- Manage file and encryption metadata in a local or remote KV store
- Upload encrypted files to Filecoin
- Retrieve encrypted files from Filecoin and decrypt them
- Share encrypted files by transferring the DEK to a specified email

**Watch the encloud CLI demo in action!**

[![encloud CLI Demo](http://img.youtube.com/vi/R-j_533QZ08/0.jpg)](https://www.youtube.com/watch?v=R-j_533QZ08 "encloud CLI Demo")

Read [**here**](docs/DESIGN.md) for detailed **design and architecture**

## Prerequisites

- Golang 1.18 or higher
- Estuary API account and key. Read [more](docs/CONFIG.md).
- CouchbaseDB for metadata storage (optional)

## CLI Setup

```bash
# go module sync
go mod tidy

# enable encloud cli command
# make sure $GOPATH environment variable is set
go build -o $GOPATH/bin/encloud github.com/encloud-tech/encloud/cmd/cli
```

## Command reference

1. Generate ECIES secp256k1 OR RSA 2048 key pair (key encryption key or KEK) to encrypt & decrypt the AES-256 keys (data encryption key or DEK). Run below command from the root of the project to the ECIES/RSA key pair

   > encloud keygen

2. Upload encrypted data to Filecoin. This command encrypts the specified file using a newly generated DEK. The DEK is encrypted using the KEK and the metadata is stored on the local KV store.
   This command also provides multiple `DEK_TYPE` options to encrypt data - `aes` or `chacha20`.

   > encloud upload -p `<KEK_PUBLIC_KEY>` -f `<UPLOAD_FILE_PATH>` -t `<DEK_TYPE>`

   Read the KEK public key from a file path instead of raw text.

   > encloud upload -p `<KEK_PUBLIC_KEY_FILE_PATH>` -f `<UPLOAD_FILE_PATH>` -t `<DEK_TYPE>` -r true

3. List uploaded files and associated metadata. Metadata is used to query and retrieve the files from Filecoin.

   > encloud contents -p `<KEK_PUBLIC_KEY>`

   Read the KEK public key from a file path instead of raw text.

   > encloud contents -p `<KEK_PUBLIC_KEY_FILE_PATH>` -r true

4. Retrieve data from Filecoin with a specific UUID. This command decrypts encrypted data on Filecoin using the relevant DEK. The DEK is stored in encrypted form in the metadata and is itself decrypted first using the KEK Private Key.

   > encloud retrieve -p `<KEK_PUBLIC_KEY>` -k `<KEK_PRIVATE_KEY>` -u `<FILE_UUID>` -s `<DOWNLOAD_PATH>`

   Read the KEK public and private keys from a file path instead of raw text. For KEK public key pass `r` flag with `true` and for private key pass `o` flag with `true`

   > encloud retrieve -p `<KEK_PUBLIC_KEY_FILE_PATH>` -k `<KEK_PRIVATE_KEY_PATH>` -u `<FILE_UUID>` -s `<DOWNLOAD_PATH>` -r true -o true

5. Share your files with other users using the UUID and DEK.

   > encloud share -e `<EMAIL>` -p `<KEK_PUBLIC_KEY>` -k `<KEK_PRIVATE_KEY>` -u `<FILE_UUID>`

   Read the KEK public and private keys from a file path instead of raw text. For KEK public key pass `r` flag with `true` and for private key pass `o` flag with `true`

   > encloud share -e `<EMAIL>` -p `<KEK_PUBLIC_KEY_PATH>` -k `<KEK_PRIVATE_KEY_PATH>` -u `<FILE_UUID>` -r true -o true

6. Retrieve shared content from other users using your CID, DEK type and DEK.

   > encloud shared -c `<RECEIVED_CID_OF_YOUR_EMAIL>` -d `<RECEIVED_DEK_FILE_PATH>` -t `<RECEIVED_DEK_TYPE>`

7. List all keys along with file metadata stored in the local KV store

   > encloud keys

8. Update configurations for the application using a compatible yaml file

> encloud config -p `<CONFIG_YAML_PATH>`

## Support

- Please file an issue to get help or report a bug
- Storage Providers and Clients : we want to work with you to integrate our tooling and offer bespoke support, please reach
  out at [contact@encloud.tech](mailto:contact@encloud.tech)
- Also check out encloud's [Sentinel and Guardian](https://www.youtube.com/watch?v=JDB6T1_Rj2s&ab_channel=encloud) products for compute on encrypted data using TEEs

## License

[Apache 2.0](https://github.com/encloud-tech/encloud/blob/main/LICENSE)
