== Installation ==

# Make sure you have the go_appengine SDK installed and located within your $PATH
# Get the qrticket source

 go get -u github.com/capnfuzz/qrtickets

# Copy app.sample.yaml to app.yaml and update directives accordingly

 cd $GOPATH/github.com/capnfuzz/qrtickets/
 cp app.sample.yaml app.yaml
 emacs app.yaml

# Test through google app engine
 
 goapp serve