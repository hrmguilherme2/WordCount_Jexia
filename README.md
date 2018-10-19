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
          `export PATH=$PATH:$HOME/go/bin:$GOPATH/bin`

<b>Step 2:</b>
 
Download and copy the project in src folder, this folder is where golang was installed.

Open terminal, type and copy:
 `open $HOME/go/src`



## Usage
Open terminal and navigate to the WordCount folder.

 `cd $HOME/go/src/WordCount`
 
 Type the following commands in terminal:

-_Will install Dep project management tool:_

`$ make deps` 

-_Will build the project and receive all dependencies:_

`$ make build`

-_Will run the program:_

`$ make run`
   
_The program accept 2 parameters where <b>-input</b> is the file where will be processed and  <b>-output</b> will be the file name after processed and generated!_
`$ go run main.go -input (Filename.txt) -output (Output filename.csv)`
    
-_Run Tests:_

`$ make test`
  
  
  Should appear for you in terminal:
  - Bytes Read
  - Bytes Write
  - Execution Time
  
  >The result will generate an .CSV file in the same folder where you extracted this project!
  ![Alt Text](https://i.imgur.com/YRmyudV.png)

