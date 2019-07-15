# Spectrum

![光谱logo2(1).png](https://upload-images.jianshu.io/upload_images/528413-0c926281c1d94539.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/440)





Spectrum is an [Ethereum-compatible](https://github.com/ethereum/go-ethereum) project. It uses a new consensus and new block reward for SmartMesh ecosystem devices and IOT. And you can view the transactions on the [
BlockChain Browser Address](https://spectrum.pub).


Since the list of signers is 17, it is recommended that the confirmation number of general transfer transaction block be set to 17 (one round), and that of exchange block be set to 34 (two rounds).

## List of Chain ID's:
| Chain(s)    |  CHAIN_ID  | 
| ----------  | :-----------:| 
| mainnet     | 20180430     | 
| testnet     | 3            | 
| devnet      | 4            | 

## Warning

We suggest that the GasPrice should not be less than 18Gwei, otherwise the transaction may not be packaged into the block.

## Build the source 

Building Spectrum requires both a Go (version 1.9.2 or later) and a C compiler. You can install them using your favourite package manager.

### Ubuntu:

	1）Install git
		$ apt-get install git
		$ apt-get install golang
	2）Install golang in the directory of {GOHOME}
		$ cd {GOHOME}
		$ tar -zxvf go1.9.2.linux-amd64.tar
		$ mkdir {GOHOME}/gopath
		$ export GOAPTH={GOHOME}/gopath 
		$ export GOROOT=${GOHOME}/go 
		$ export PATH=$GOROOT/bin:$PATH
	3）download the source on the git
		$ cd {GOPATH}
		$ mkdirs src\github.com\SmartMeshFoundation\
		$ cd src\github.com\SmartMeshFoundation\
		$ git clone https://github.com/SmartMeshFoundation/Spectrum.git
	4）compile the source code
		$ cd Spectrum
		$ make smc
	5）run the program
		$ ./build/bin/smc console

### Mac


	1）download the latest *Spectrum* source code from the github 
	
        	$ git clone https://github.com/SmartMeshFoundation/Spectrum.git

	2）Install by using source code

        	$ cd Spectrum
        	$ make smc

		If any error happens during the compiling and prompt message shows " lacking of the header file that is related to the Mac OS", you may try to install the xcode command-line tool before executing above commands:
	
	        $ xcode-select --install

	3） startup
	
        	$ build/bin/smc console

### Windows

	1）First of all, you need to install a package management software for Windows named "chocolatey". Please refer to https://chocolatey.org for Installation method.
	2）Install git, golang, mingw by using chocolatey
		c:\Users\xxx> choco install git
		c:\Users\xxx> cholo install golang 
		c:\Users\xxx> cholo install mingw

	3）Set environment variables {GOPATH} of golang：
		c:\Users\xxx> mkdir {GOPATH}
		c:\Users\xxx> set "GOPATH={GOPATH}"
		c:\Users\xxx> set "Path={GOPATH}\bin;%Path%"
	4）download the source on the git： 
		c:\Users\xxx> cd {GOPATH}
		{GOPATH}> mkdir src\github.com\SmartMeshFoundation\
		{GOPATH}> cd src\github.com\SmartMeshFoundation\
		...\SmartMeshFoundation> git clone https://github.com/SmartMeshFoundation/Spectrum.git
	5）compile source code：
		...\SmartMeshFoundation> cd Spectrum
		...\SmartMeshFoundation> go install -v ./cmd/...
	6）run the program：
		c:\Users\xxx> smc console

## Encrypt your nodekey

     $ ./build/bin/smc security --passwd
## Decrypt your nodekey

     $ ./build/bin/smc security --unlock
     
## Run fast node to test smc

    $ ./build/bin/smc console
    
## Create new account

    > personal.newAccount()

## Get your own miner id

    every node has it's own miner id, you can run getMiner() function to get that id:

    > tribe.getMiner() 

    then you will see below messages:
    {
        address: "0x00ab501f3fe4b2f71651764699ec5752598e679f",
        level: "Signer"
    }
    
## Bind your own miner id  to  wallet address

    Users can bind their block ID to a wallet address:

    > tribe.bind("account","password") 

## Deposit token to POC

    Users can become miner by deposit smt:

    > tribe.pocDeposit("account","password") 


## Start mining

    Users can start mining or resume mining:

    > tribe.pocStart("account","password") 


## Stop mining

    Users can stop mining:

    > tribe.pocStop("account","password") 
    
## Withdraw smt

    Users can withdraw smt:

    > tribe.pocWithdraw("account","password")   
    
# More functions   
    Users can input tribe to view:
    
    > tribe

