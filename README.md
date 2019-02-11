# Spectrum

![光谱logo2(1).png](https://upload-images.jianshu.io/upload_images/528413-0c926281c1d94539.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/440)





[![](https://travis-ci.org/wangbaojin/Spectrum.svg?branch=master)](https://travis-ci.org/wangbaojin/Spectrum)

Spectrum is an [Ethereum-compatible](https://github.com/SmartMeshFoundation/Spectrum) project. It uses a new consensus and new block reward for SmartMesh ecosystem devices and IOT. 

Since the list of signers is 17, it is recommended that the confirmation number of general transfer transaction block be set to 17 (one round), and that of exchange block be set to 34 (two rounds).

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

## ChainId
   mainnet: 20180430 

## Encrypt your nodekey

     $ ./build/bin/smc security --passwd
## Decrypt your nodekey

     $ ./build/bin/smc security --unlock
     
## Run fast node to test smc

    $ ./build/bin/smc console
    
## Create new account

    > personal.newAccount()

## View the miner nodes

* tribe.getStatus()
* then you will see the following message:
<pre><code>{
  "epoch": 5760,
  "number": 39601,
  "signerLevel": "None",
  "signerLimit": 17,
  "signers": [
    {
      "address": "0x3a5fbac6ca913599c5fde8c1638db58d01de8a48",
      "score": 3
    },
    {
      "address": "0xad4c80164065a3c33dd2014908c7563eff88ab49",
      "score": 3
    },
    {
      "address": "0xc22d53456abd14da347517a4b47ea24866b8e3ae",
      "score": 3
    },
    {
      "address": "0x7b06dd132c089034157f1e1aacda75787df1e0c5",
      "score": 3
    },
    {
      "address": "0x00ab501f3fe4b2f71651764699ec5752598e679f",
      "score": 3
    }
  ],
  "totalSinner": 2,
  "volunteerLimit": 70,
  "volunteers": []
}</code></pre>

that tell you there are two miners in the testnet.

## View the block mining history in console

* tribe.getHistory(11,false)
    
* then you will see the following message (console format):
<pre><code>[{
    🔨: "39708 -> 0x7B06dd132c089034157f1E1AAcda75787DF1e0c5"
}, {
    🔨: "39707 -> 0xc22D53456ABd14Da347517a4B47ea24866B8E3Ae"
}, {
    🔨: "39706 -> 0xAd4c80164065a3c33dD2014908c7563eFf88Ab49"
}, {
    🔨: "39705 -> 0x3a5fBaC6CA913599C5fde8c1638dB58d01De8A48"
}, {
    🔨: "39704 -> 0x00aB501f3Fe4b2f71651764699EC5752598E679f"
}]</code></pre>

that tell the block number and miner's account that generate that block.

## get your own miner account

    every node has it's own miner account, you can run getMiner() function to get that account:

    > tribe.getMiner() 

    then you will see below messages:
    {
        address: "0x00ab501f3fe4b2f71651764699ec5752598e679f",
        balance: 2001223531052513000,
        level: "Signer"
    }
    that will show your miner account and the balance of miner account in Wei unit.

## becoming a miner 

    if you want to become a miner, you will send 1 smt to the miner account, after this transfer,
    you will run miner.start() to join into the miners group.

    > miner.start()
