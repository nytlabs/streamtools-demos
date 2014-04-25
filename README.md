# streamtools-demos

[Streamtools](https://github.com/nytlabs/streamtools) is a graphical toolkit for dealing with streams of data. Streamtools makes it easy to explore, analyse, modify and learn from streams of data. 

This set of demos gets you up and running with Streamtools in no time flat. 

## Download and run the demos

```code
mkdir -p $GOPATH/src/github.com/nytlabs
cd $GOPATH/src/github.com/nytlabs
git clone https://github.com/nytlabs/streamtools-demos.git
cd streamtools-demos
go get .
go build
./streamtools-demos
```

You'll find the list of demos at [http://localhost:8080/](http://localhost:8080/)

**IMPORTANT!** 

This demo requires you to have Streamtools installed already. To actually view the demos in Streamtools you'll need to have another browser window or tab open and pointed at [http://localhost:7070/](http://localhost:7070/)
