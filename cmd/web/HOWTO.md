# encloud Desktop Application - How To Use

The encloud Desktop Application is a cross-platform application (Windows, MacOS and Linux) built using the Wails framework. It combines a React based
frontend with the Golang based APIs into a seamless desktop experience for users who prefer a GUI. The Desktop App has 
an intuitive interface that is better suited for use by data analysts, data scientist and business analysts that prefer 
interacting with a UI.

The application has four basic flows -

1. Generate And Manage Key Encryption Keys
2. Upload Encrypted Files To Filecoin
3. Retrieve Encrypted Files From Filecoin
4. Retrieve Shared Files

## Generate And Manage Key Encryption Keys
<img src="../../.github/Screen 1.png" alt="Screen 1" width="50%" height="50%"/>

The Key Encryption Key is used to encrypt the Data Encryption Key generated for each file upload as part of the envelope encryption process. KEK is an assymetric key with a public and private key component. The user can easily generate their KEKs from within the Generate Key Section. They can choose between RSA and ECIES schemes. After generating, the user can set this generated key as the current key. The user should be careful in handling this key as the application tracks all file metadata in the local store using the public key. Replacing a key will result in the metadata not being visible. Keys can easily be set to recover metadata associated with them.  

## Upload Encrypted Files To Filecoin
<img src="../../.github/Screen 2.png" alt="Screen 1" width="50%" height="50%"/>

Once a KEK has been set by the user, they can start uploading their files to Filecoin. To do this they must select a DEK type - 
* AES 256 GCM - standard for file encryption recommended by NIST
* ChaCha20-Poly1035 - an improved algorithm that works better for larger files and is more secure than AES 256

Users can then select the file to upload and simply upload it. The app encrypts the file and on successful onboarding to Filecoin, creates an entry in the metadata store.

## Retrieve Encrypted Files From Filecoin
<img src="../../.github/Screen 3.png" alt="Screen 1" width="50%" height="50%"/>

Users can see a list of all files they have uploaded and the associated metadata from their local storage. The visible metadata is 
* UUID - Internal unique identifier for the file used to track it in the metadata 
* Filename - Name of the File
* FileSize - Size in bytes
* CID - CID of the file in the Filecoin network
* Retrieve/Share - Button to perform actions; Retrieve opens the file detail page where the user can download the file; Share opens an pop-up where an email can be entered to share the file with

## Retrieve Shared Files 
<img src="../../.github/Screen 4.png" alt="Screen 1" width="50%" height="50%"/>

Users can retrieve encrypted files that have been shared with them via email. Users will receive an email with the file CID and the DEK type, along with the DEK as an attachment to the email. Users will need to use the CID and the DEK to retrieve the file from Filecoin, decrypt it and save it locally.

Users can use the form to retrieve the file by filling out:
* CID of the file from the shared email
* DEK type as per the shared email
* Path to the DEK attachment from the email 
* File Name for the file to be downloaded
* Path of the file to be downloaded
