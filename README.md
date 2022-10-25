# Filecoin-encrypted-data-storage

A filecoin-encrypted-data-storage made in Go framework. It is able to upload and download encrypted data to filecoin.

## Steps To Execution
1) Generate key pair to encrypt & decrypt dek. Run below command to root path of the project to generate key pair
    > go run . generate-key-pair
2) Upload encrypted data to filecoin server. This command encrypted your data using dek. It also encrypt your dek using generated public key. 

    > go run . upload -p `<GENERATED_PUBLIC_KEY>` -f `<UPLOAD_FILE_PATH>` 
3) List your uploaded data. 

    > go run . list -p `<GENERATED_PUBLIC_KEY>`
4) Retreive data from filecoin server. This command decrypt your data using dek. It also decrypt your dek using generated private key. 

    > go run . retrieve-by-cid -p `<GENERATED_PUBLIC_KEY>` -k `<GENERATED_PRIVATE_KEY>` -u `<UUID>`       

