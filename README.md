# iso20022

Go implementation of the ISO 20022 standard

This package consumes valid iso20022 XML strings and returns Go structs to be processed by an iso20022 application.


# Setup
go get github.com/figassis/bankiso  
go get github.com/davecgh/go-spew/spew  --for debugging purposes only. Not required if not building

cd $GOPATH/src/github.com/figassis/bankiso
go build  


# Testing

./bankiso -f bankiso -f msg-examples/payments/pain.009.001.04/Business\ example3\ pain.009.001.04.xml

# Usage in a project
Take a look at https://github.com/figassis/bankiso/blob/master/main.go

Each iso20022 schema results in a different struct, and therefore a different go type.
I've decided to use interfaces, type assertions and type switches to provide a sane way to handle all messages.

I’ve created an interface ISOMessage, which all iso20022 (and later 8583) messages implement, so they can all be returned from the same function.

The package maintains a library of struct types that it picks according to the message code it reads from the XML string.

A function then unmarshals the XML into the chosen struct and returns an ISOMessage.

You can then use a type switch to type assert the message and call the correct handling function.


## License
[Q Public License 1.0 (QPL-1.0)](https://en.wikipedia.org/wiki/Q_Public_License)

Copyright &copy; Assis Ngolo
