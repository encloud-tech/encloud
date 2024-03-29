## encloud shared

Retrieve shared content

### Synopsis

Retrieve shared content from other users using your CID, DEK type and DEK

```
encloud shared [flags]
```

### Options

```
  -c, --cid string       CID of shared file to retrieve
  -d, --dek string       Path to DEK file
  -h, --help             help for shared
  -n, --name string      Name of retrieved file
  -s, --storage string   Path to store retrieved file under
  -t, --type string      DEK type (default "chacha20")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.cobra.yaml)
      --viper           use Viper for configuration (default true)
```

### SEE ALSO

* [encloud](encloud.md)	 - encloud is a CLI tool for on-boarding sensitive data to Filecoin.

###### Auto generated by spf13/cobra on 28-Jul-2023
