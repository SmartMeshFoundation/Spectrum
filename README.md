# SMChain

SmartMesh Chain is an [Ethereum-based](https://github.com/SmartMeshFoundation/SMChain) project. It uses a new consensus and new block reward for SmartMesh ecosystem devices and IOT. 

SMChain is still in the development stage, and this version is just for testing.

## Installation on Linux/macOS 

Building SmartMesh Chain requires both Go Language(Version 1.9.2 or later) and C Language. You can install them using your favourite package manager. Once the dependencies have been installed, run

    $ git clone https://github.com/SmartMeshFoundation/SMChain.git

    $ cd ./SMChain

    $ make geth

or build a full suite of utilities: make all：

    $ make all

## Run full node to test geth

    $ ./build/bin/geth --testnet --syncmode full console

## Connect to the testnet

    > admin.addPeer("enode://f9cf4fdc53ce5ecf84e7c26fbe262bddd7cbf3d17593328f74816a1c646a0ccfac9a85d81f7d51b59bc02dc8f0e8c5dada4db081efd79698a820faf9384773c0@49.51.11.222:60303")

    wait a few seconds, and you will see the block sync is begin.

## View the miner node

	> tribe.getStatus()

    then you will see the following message:
    {
        number: 243,
        signers: [{
            address: "0x4110bd1ff0b73fa12c259acf39c950277f266787",
            score: 3
        }, {
            address: "0x76d564fe610ddf349333acbfe4a58abfd1d3ca04",
            score: 3
        }],
        volunteers: []
    }

    that tell you there are two miners in the testnet.

## more works
    more functions is now under development, and will show you later.

