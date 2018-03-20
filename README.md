# Golang Networking

The project will try and encapsulate how golang can be used in working with network protocols and related topics.

## Data Serialization

Protocols like TCP/UDP/IP allow for data to be moved over the network but once the data is transported there is a need
for us to understand how the data is structured. The process of data serialization is called marshalling and unmarshalling

There are various serialization specifications like the XDR and the ONC (Open Network Computing). Golang uses the gob
serialization.

#### Self Describing data
This data carries with itself the type information. The idea is that the marshaller must follow the same format that
the unmarshaller understands.

#### ASN.1
The ASN.1 format is supported by the encoding/asn1 package we can marshal normal data structures like string, int, datetime
along with support for structures as well. The marshalling and unmarshalling api is similar to yaml package.

#### JSON
the JSON support is standard and is present in the encoding/json package. The package in json encoding has two compoenents
1. encoder - the encoder is initialised by a io.Writer as input. it is used to write data to the stream.
2. decoder - the decoder is initialized by a io.Reader as input and it is used to read from a stream.

```
    /*
        conn is a connection from a tcp or udp client or server type
        either from net.Dial() or net.Listen() methods.
    */
    encoder, err := json.NewEncoder(conn)
    encoder.Encode(person)  // person is a structure
    ...

    decoder, err := json.NewDecoder(conn)
    var newPerson Person
    decoder.Decode(&newPerson)
```

when working with low level tcp or udp packages we can wrap the connection object in the encoder and decoder and read and write
to the connection that we establish.

#### Gob package
This is specific to the golang and is not a format supported by any other language. Gob can be used when you know the client
servers are both in golang. it supports all data type, structs but does not support channels and functions.  Also circular data
structures are also not very well supported. The api for the package is similar to the one of the json encoding package. Similar
encoder and decoders can be found.