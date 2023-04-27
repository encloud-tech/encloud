# Configuration

Encloud Encryption and Storage CLI provides an ability to configure the various aspects of the application to suit the clients needs.

The following can be configured:

* Filecoin on-boarding and deal making mechanism (Estuary support offered currently)
* Metadata Storage Mechanism (BadgerDB and Couchbase support offered) 
* Email settings for sharing of DEKs for encrypted files on Filecoin

## Estuary 

Encloud Encryption and Storage CLI uses Estuary as a means to onboard and retrieve data from the Filecoin network. Estuary requires the client
to generate an API Key which can be requested [here](https://docs.estuary.tech/tutorial-get-an-api-key). While filling out the request [form](https://docs.estuary.tech/get-invite-key)
please mention "Encloud".

*Note: Estuary is under active development and hence unstable. Please see the latest estuary documentation or the [#ecosystem-dev](https://filecoinproject.slack.com/archives/C016APFREQK) channel on the Filecoin slack for relevant updates on API statuses.*

The Estuary API key needs to be configured under [config.yaml](../config.yaml) as follows under the `estuary` section:

```yaml
estuary:
  base_api_url: https://api.estuary.tech
  upload_api_url: https://edge.estuary.tech/api/v1
  gateway_api_url: https://edge.estuary.tech
  cdn_api_url: https://cdn.estuary.tech
  token: "XXXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
```

## Key Encryption Key

Encloud CLI Encryption and Storage CLI offers supports for both RSA and ECIES type asymmetric encryption schemes for the 
Key Encryption Key as part of the envelope encryption mechanism. 

This can be easily configured in the [config.yaml](../config.yaml) as follows under the `estuary` section:

```yaml
stat:
  kekType: ecies
```

Use `rsa` or `ecies`. Even though Encloud utilizes RSA-2048-OAEP, which is also used by major Web2 CSPs, there are known
vulnerabilities in its security and longevity. The KEK being the key encrypting all metadata, it is important that the KEK
follows best practices regarding longevity and security. ECIES encryption is considered more secure and offer better longevity.

ECIES is additionally supported by decentralized key custody solutions and users can leverage decentralized custody if they
choose ECIES scheme.

## Storage

### BadgerDB

BadgerDB is lightweight key-value store that can be used without any additional configuration. However, there are scalability issues while using
BadgerDB in production scenarios.

To use BadgerDb use the following configuration under [config.yaml](../config.yaml), under the `stat` section:

```yaml
stat:
  storageType: badgerdb
  badgerdb:
    path: badger.db
```

### Couchbase

Couchbase Server is an open source, distributed, NoSQL document-oriented engagement database. It exposes a fast key-value 
store with managed cache for sub-millisecond data operations, purpose-built indexers for fast queries and a powerful query engine for executing SQL-like queries.

Couchbase server can be utilized as a KV store for metadata. Please follow below instructions to install and setup couchbase server locally.

#### Installation

To install Couchbase Server please follow the instructions [here](https://docs.couchbase.com/server/current/install/install-intro.html).

##### Starting Couchbase Server

Once Couchbase Server has been installed simply navigate to where it has been installed and start "Couchbase Server".

To start Couchbase Server using Docker please see the documentation [here](https://docs.couchbase.com/server/6.0/getting-started/do-a-quick-install.html).

##### Accessing Couchbase Server

Couchbase Server can be accessed using
* [CLI](https://docs.couchbase.com/server/current/cli/cli-intro.html)
* [API](https://docs.couchbase.com/server/current/rest-api/rest-intro.html)
* An [administration (web) portal](https://docs.couchbase.com/server/current/getting-started/look-at-the-results.html)

#### Setup
Couchbase requires setting up a bucket to hold scopes, scopes that contain collection and collections that contain documents.
These need to be setup before using Couchbase as a store

##### Creating Bucket on Couchbase Server

To create a bucket on couchbase server please follow the instructions [here](https://docs.couchbase.com/server/current/manage/manage-buckets/create-bucket.html)

##### Manage scope and collection on Couchbase Server

To manage scope and collection of bucket on couchbase server please follow the instructions [here](https://docs.couchbase.com/server/current/manage/manage-scopes-and-collections/manage-scopes-and-collections.html)

##### Managing indexes on Couchbase Server
To manage and create primary or secondary indexes on couchbase server to fetch data please follow the instructions [here](https://docs.couchbase.com/server/current/manage/manage-indexes/manage-indexes.html)

##### Set credentials

Once Couchbase Server has been started and a bucket has been created then set the host, port, username, password and bucketName in **config.yaml** file.
A scope and a collection needs to be created within the bucket to store documents. These params can also be set in the config.

```yaml
stat:
  storageType: couchbase
  badgerdb:
    path: badger.db
  couchbase:
    host: localhost
    username: Administrator
    password: Encloud@2022
    bucket:
      name: encloud
      scope: file
      collection: metadata
```

## Email and sharing

We utilize emails to share DEKs from the client directly to the email of an entity they want to share the data with. It is
worth noting that once the DEK is shared with an entity they can download the data from Filecoin and decrypt it.

The following configs need to be made for emails under [config.yaml](../config.yaml), under the `email` section:

```yaml
email:
  server: "smtp.acme.io"
  port: 2525
  username: "XXXXXXXXXXXXX"
  password: "XXXXXXXXXXXXX"
  from: "noreply@acme.com"
```
