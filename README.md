# Spectrum

Spectrum is an [Ethereum-compatible](https://github.com/SmartMeshFoundation/Spectrum) project. It uses a new consensus and new block reward for SmartMesh ecosystem devices and IOT. 

Spectrum is still in the development stage, and this version is just for testing.

## Build the source 

For prerequisites and detailed build instructions please read the [Installation Instructions](https://github.com/SmartMeshFoundation/Spectrum/wiki/Building-Specturm) on the wiki.

Building geth requires both a Go (version 1.9.2 or later) and a C compiler. You can install them using your favourite package manager.
    
## Run fast node to test geth

    $ ./build/bin/geth console

## Create new account

    > personal.newAccount()

## View the miner nodes

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

## View the block mining history

    > tribe.getHistory()
    
    then you will see the following message:
    [{
        🔨: "9497 -> 0x76A45989Dd96f58D5cF3d36153068D7B347f7b55"
    }, {
        🔨: "9496 -> 0x00aB501f3Fe4b2f71651764699EC5752598E679f"
    }, {
        🔨: "9495 -> 0x7B06dd132c089034157f1E1AAcda75787DF1e0c5"
    }, {
        🔨: "9494 -> 0xc22D53456ABd14Da347517a4B47ea24866B8E3Ae"
    }, {
        🔨: "9493 -> 0xAd4c80164065a3c33dD2014908c7563eFf88Ab49"
    }, {
        🔨: "9492 -> 0x3a5fBaC6CA913599C5fde8c1638dB58d01De8A48"
    }, {
        🔨: "9491 -> 0x76A45989Dd96f58D5cF3d36153068D7B347f7b55"
    }, {
        🔨: "9490 -> 0x00aB501f3Fe4b2f71651764699EC5752598E679f"
    }, {
        🔨: "9489 -> 0x7B06dd132c089034157f1E1AAcda75787DF1e0c5"
    }, {
        🔨: "9488 -> 0xc22D53456ABd14Da347517a4B47ea24866B8E3Ae"
    }, {
        🔨: "9487 -> 0xAd4c80164065a3c33dD2014908c7563eFf88Ab49"
    }]
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
