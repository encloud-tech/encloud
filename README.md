# Encloud

A encloud made in Go framework. It is able to upload and download encrypted data to filecoin.

## Development server
 
 ```bash
go mod tidy

# enable encloud cli command
go install .
```

## Steps To Execution
1) Generate key pair to encrypt & decrypt dek. Run below command to root path of the project to generate key pair
    > go run . generate-key-pair
2) Upload encrypted data to filecoin server. This command encrypted your data using dek. It also encrypt your dek using generated public key. 

    > go run . upload -p `<GENERATED_PUBLIC_KEY>` -f `<UPLOAD_FILE_PATH>` 

    Read public key from path option allowing to pass public key filepath instead of publickey.

    > go run . upload -p `<GENERATED_PUBLIC_KEY_PATH>` -f `<UPLOAD_FILE_PATH>` -r true
3) List your uploaded data. 

    > go run . list -p `<GENERATED_PUBLIC_KEY>`

    Read public key from path option allowing to pass public key filepath instead of publickey.

    > go run . list -p `<GENERATED_PUBLIC_KEY_PATH>` -r true
4) Retreive data from filecoin server. This command decrypt your data using dek. It also decrypt your dek using generated private key. 

    > go run . retrieve-by-cid -p `<GENERATED_PUBLIC_KEY>` -k `<GENERATED_PRIVATE_KEY>` -u `<UUID>`

    Read key from path option allowing to pass key path instead of key value. For public key pass r flag with true value and for private key pass o flag with true value

    > go run . retrieve-by-cid -p `<GENERATED_PUBLIC_KEY_PATH>` -k `<GENERATED_PRIVATE_KEY_PATH>` -u `<UUID>` -r true -o true

## Share content
1) Share your content to other user using your cid and dek.

    > go run . share -e `<EMAIL>` -p `<GENERATED_PUBLIC_KEY>` -k `<GENERATED_PRIVATE_KEY>` -u `<UUID>`

    Read key from path option allowing to pass key path instead of key value. For public key pass r flag with true value and for private key pass o flag with true value

    > go run . share -e `<EMAIL>` -p `<GENERATED_PUBLIC_KEY_PATH>` -k `<GENERATED_PRIVATE_KEY_PATH>` -u `<UUID>` -r true -o true

2) Retrieve shared content from other user using your cid and dek.

    > go run . retrieve-shared-content -c `<RECEIVED_CID_OF_YOUR_EMAIL>` -d `<RECEIVED_DEK_FILE_PATH>`