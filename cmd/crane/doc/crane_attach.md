## crane attach

Add an attachment to an image

```
crane attach IMAGE FILE [flags]
```

### Options

```
  --artifact-type string   value for config.mediaType field (default "application/vnd.oci.image.config.v1+json")
  -h, --help               help for attach
  --media-type string      value for layers[0].mediaType field (default "application/vnd.oci.image.layer.v1.tar+gzip")
```

### Options inherited from parent commands

```
      --allow-nondistributable-artifacts   Allow pushing non-distributable (foreign) layers
      --insecure                           Allow image references to be fetched without TLS
      --platform platform                  Specifies the platform in the form os/arch[/variant][:osversion] (e.g. linux/amd64). (default all)
  -v, --verbose                            Enable debug logs
```

### SEE ALSO

* [crane](crane.md)	 - Crane is a tool for managing container images

