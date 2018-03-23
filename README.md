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


## Protocol Design
There are number of issue that are involved in protocol design:

* It is broadcast or point to point. Broadcast protocol must use UDP or local multicast. For p2p use TCP or UDP.
* Is it useful to be stateless or stateful? if is better for one side to maintain state when compared to the other side.
* Is the transport protocol reliable or unreliable?
* Are replies needed? if they are then what happens when one does not arrive. (timeouts)
* what data format do you want? MIME or byte encoding are common possibilities
* Is the communication needed bursty or steady stream? Ethernet and internet are good at bursty traffic. For steady stream needed in case of video.
if the steady stream is needed what about the QoS (quality of service).
* Are multiple streams with sync required? does that data need to be sync with anything.
* Are you building a standalone application or a lib for others to use.


#### Message Format
In client server interaction with messages we have two parts:
1. **message type** (this can be either integers or strings - HTTP has int codes to have message type)
2. **message content** - this is the exact message

#### Data Format

**Byte Format** - in byte format the message content is represented as a series of bytes.
* The first byte identifies the message type.
* based on the message type the handler for handling the message will be choosen
* other bytes in the message will conform to the content according to the pre-defined format (json, gob, asn1 etc)

Advantages of byte format is that the data is compact and fast where as the disadvantage is that it is difficult to debug.

```
    handleClient(conn) {
        while(true){
            byte b = conn.readByte()
            switch(b) {
                case MSG1:
                        ...
                case MSG2:
                        ...
            }
        }
    }
```

**Character Format** - here everything that is sent is a character if possible. Integer 234 is sent as 3 characters '2', '3', '4'.

In character format:
* A message is sequence of one or more lines. The first word in the message is the indicator of the message type.
* String handling functions are used to determine the message type and decode message.
* rest of the first line and other lines are data.
* all handling happens line by line.

```

     handleClient(conn) {
        line = conn.readLine()
        if line.startsWith(..){
            ...
        } else if line.startsWith(...){
            ...
        }
      }

```

Character sets are not easy handle because of the encoding the data may have.


# Character sets and Encoding
A **character** is a unit of information that roughly corresponds to a symbol of a natural language, such as a letter, number or punctuation mark.

**Character Code** - is mapping of a character to an integer, e.g. in ASCII code set 'a' is denoted by 97 and 'A' denoted by 65.
this code is still not the one that you see in text files or TCP packets, there are few encoding left.

**Character encoding** - To communicate or store the characters on disk or over the network the characters are encoded in binary format. All encoding will
in terms of 8, 32, 64 or 128 formats. So the type of encoding used gives us the code.

**Transport encoding** - this is encoding that is done over an above character encoding that allows for transport of data over the network.

The information of type of encoding is sent in headers for any protocol. The header has infromation like the one given below

```
Content-Type: text/html; charset=ISO-8859-4
Content-Encoding: gzip
Transfer-Encoding: chunked
```

Here we see that the content encoding is done using gzip where as transfer encoding is chunked.
* **chunked** - data is sent in a series of chunks and then separated by \r\n
* **compress** - a format that uses the LZW algorithm to compress the content
* **deflate** - uses the zlib structure and a deflate algorithm for compression.
* **gzip** - uses LZ coding with 32 bit CRC
* **identity** - no compression or modification.


#### ASCII
This character set is small it holds all the english aplhabets and numbers plus some punctuations and control characters it is represented in 7 bit format.

#### ISO-8859
This character set is an octet 8 bits format that adds 128 characters to the set which allows for support of more characters.
But still there are ISO-8859-1 to ISP-8859-10 sets that support different languages.

#### Unicode
When it came to languages like english the character set is small so that ISO or ASCII can be used easily, but when it comes to
languages like Chinese there are 20,000 characters 5000 in common use there is a need to use 2 byte character set code.
Even Unicode has become small to hold all characters so we have the following:
* UTF-32 (4 byte encoding)
* UTF-16 encodes in 2 bytes with 2 bytes for overflow.
* UTF-8 uses 1 to 4 bytes for encoding per character depending on the character code.


# Security

Security has Functions and Levels and these correspond to the OSI architecture.

#### Functions
Principle functions of the security system are:
1. Authentication
2. Data Integrity
3. Confidentiality
4. Notaization/signature
5. Access Control
6. Assurance/availability

#### Levels
* Peer entity authentication
* Data origin authentication
* Access control service
* Connection confidentiality
* Connectionless confidentiality
* Selected field confidentiality
* Traffic flow confidentiality
* Connection integrity with recovery
* Connection integrity without recovery
* Connection integrity selective field
* Connectionless integrity selective field
* Non repudiation at origin
* Non repudiation of receipt

### Data Integrity
data integrity is all about insuring the data is not tampered with. This is usually done
by calculating a hash value (simple number) out of the bytes of the data.

For security purpose the hashing algorithm needs to be good so that the hackers cannot crack the code and tamper with the data.
There are a series of cryptographic algorithms that are used for calculating the hash value.

Golang supports MD4, MD5, RIPEMD-160, SHA1, SHA224, SHA256, SHA384 and SHA512
