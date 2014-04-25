# streamtools-demos

[Streamtools](https://github.com/nytlabs/streamtools) is a graphical toolkit for dealing with streams of data. Streamtools makes it easy to explore, analyse, modify and learn from streams of data. 

This set of demos gets you up and running with streamtools in no time flat. 

## Setup

The demos require a local instance of streamtools running on the default port (7070).

```code
mkdir -p $GOPATH/src/github.com/nytlabs
cd $GOPATH/src/github.com/nytlabs
git clone https://github.com/nytlabs/streamtools-demos.git
cd streamtools-demos
go get .
go build
./streamtools-demos
```

You'll find the list of demos at http://localhost:8080/
