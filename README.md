# WordCount Jexia.
## Requirements

- Golang Dep project management
https://github.com/golang/dep

- Golang Glow MapReduce Framework
https://github.com/chrislusf/glow

## Installation
<b>Step 1:</b>

  OS X:

  Install Golang
    https://golang.org/doc/install?download=go1.11.1.darwin-amd64.pkg

   Once Golang installed, go to "Terminal" and type:
          `$ export PATH=$PATH:$HOME/go/bin:$GOPATH/bin`

<b>Step 2:</b>
 
Download and copy the project in src folder, this folder is where golang was installed.

Open terminal, type and copy:
 `$ open $HOME/go/src`



## Usage
Open terminal and navigate to the WordCount folder.

 `cd $HOME/go/src/WordCount`
 
 Type the following commands in terminal:
```rust
//Will install Dep project management tool
$ make deps
```


```rust
//Will build the project and receive all dependencies
$ make build
```

    
 ```rust
//Run Tests
$ make test
 ```
 
```rust
//The program accept 2 parameters where -input is the file where will be processed and -output will be the file name after processed and generated.
//Will run the program
$ make run
```

  
  Should appear for you in terminal:
  - Bytes Read
  - Bytes Write
  - Execution Time
  
  >The result will generate an .CSV file in the same folder where you extracted this project!
  ![Alt Text](https://i.imgur.com/YRmyudV.png)

