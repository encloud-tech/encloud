# Configuration

Encloud Encryption and Storage CLI provides an ability to configure the various aspects of the application to suit the clients needs.

The following can be configured:

* Filecoin on-boarding and deal making mechanism (Estuary support offered currently)
* Metadata Storage Mechanism (BadgerDB support offered currently) 
* Email settings for sharing of DEKs for encrypted files on Filecoin

## Estuary 

Encloud Encryption and Storage CLI uses Estuary as a means to onboard and retrieve data from the Filecoin network. Estuary requires the client
to generate an API Key which can be requested [here](https://docs.estuary.tech/tutorial-get-an-api-key). 

The Estuary API key needs to be configured under [config.yml](../config.yml) as follows under the `estuary` section:

```yaml
estuary:
  base_api_url: "https://api.estuary.tech"
  download_api_url: "https://dweb.link/ipfs"
  shuttle_api_url: "https://shuttle-4.estuary.tech"
  token: "XXXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
```

## Email and sharing

We utilize emails to share DEKs from the client directly to the email of an entity they want to share the data with. It is 
worth noting that once the DEK is shared with an entity they can download the data from Filecoin and decrypt it.

The following configs need to be made for emails under [config.yml](../config.yml), under the `email` section:

```yaml
email:
  server: "smtp.acme.io"
  port: 2525
  username: "XXXXXXXXXXXXX"
  password: "XXXXXXXXXXXXX"
  from: "noreply@acme.com"
```

