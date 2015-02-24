mqtt_golang_test
====
My MQTT & Golang study.

Environment settings for me...
----
    $ cd ~
    $ apt-get install golang mercurial
    
    $ vi ~/.bash_profile
    
      #add below lines...
      export GOPATH=$HOME/go
      export PATH=$PATH:$GOPATH/bin/
    
    $ . ~/.bash_profile
    
    $ mkdir go/src/github.com/yoggy
    $ cd go/src/github.com/yoggy
    $ git clone git@github.com:yoggy/mqtt_golang_test.git
    
    $ cd mqtt_golang_test
    $ go get
    
    $ cp config.json.template config.json
    $ vi config.json
    
    $ go run mqtt_sub.go
    
    $ go run mqtt_pub.go
    
    

