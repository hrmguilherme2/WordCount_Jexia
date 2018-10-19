# WordCount Jexia.


## Installation
<b>Step 1:</b>

  OS X:

  Install Golang
    https://golang.org/doc/install?download=go1.11.1.darwin-amd64.pkg

   Once Golang installed, go to "Terminal" and type:
          `export PATH=$PATH:$HOME/go/bin:$GOPATH/bin`

<b>Step 2:</b>
 
Download and extract the project in src folder, this folder is where golang was installed.

Open terminal and type:
 `open $HOME/go/src`



## Usage
Open terminal and navigate to the WordCount folder.
 `cd $HOME/go/src/WordCount`
 
 Type the following commands in terminal:
 
   > _Will install Dep project management tool:_

   `make deps` 

   > _Will build the project and receive all dependencies:_

   `make build`

   > _Will run the program:_

   `make run`

  > _Run Tests:_

  `make test`
