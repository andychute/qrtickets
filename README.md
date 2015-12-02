# Installation #
1. Make sure you have the go_appengine SDK installed and located within your $PATH
2. Get the qrticket source

```bash
go get -u github.com/capnfuzz/qrtickets
```

3. Copy app.sample.yaml to app.yaml and update directives accordingly

```bash
 cd $GOPATH/github.com/capnfuzz/qrtickets/
 cp app.sample.yaml app.yaml
 emacs app.yaml
```

4. Test through google app engine
```bash
 goapp serve
```