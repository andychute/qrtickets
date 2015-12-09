# Installation #
* Make sure you have the go_appengine SDK installed and located within your $PATH
* Get the `qrticket` source

`go get -u github.com/capnfuzz/qrtickets`

*  Copy app.sample.yaml to app.yaml and update directives accordingly.  Change the X,Y,and D values to match your ECDSA P229 Private Key.  Change the HTTP_AUTH environment variable to a good string which you will require in your HTTP requests to admin-only calls

```
cd $GOPATH/github.com/capnfuzz/qrtickets/
cp app.sample.yaml app.yaml
emacs app.yaml
```


* Test through google app engine
`goapp serve`