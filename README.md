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

## Give it a try ##

* Start your local development environment
`goapp serve`

* Create a test signature to use within your application by visiting http://localhost:8080/gensig

* Take the content from the "Marshaled Private Key" and smash it into your app.yaml file in the same sort of structure as outlined within app.sample.yaml

* Set a "X-Ticket-Auth" header within your browser and set the value to match the one specified in the HTTP_AUTH environment variable within your app.yaml file

* Add a test event by visiting http://localhost:8080/testadd in your browser

* A datastore key for the event will be displayed in your browser.  Copy the key and visit http://localhost:8080/api/v1/events/{THEDATASTOREKEY}/tickets/add to add a ticket

* A QR code will be displayed, you can test that this code works and claims properly by visiting http://localhost:8080/api/v1/tickets/{THECONTENTENCODEDINTHEBARCODE}/claim