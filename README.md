# Spectrum

SmartMesh Chain is an [Ethereum-based](https://github.com/SmartMeshFoundation/Spectrum) project. It uses a new consensus and new block reward for SmartMesh ecosystem devices and IOT. 

Spectrum is still in the development stage, and this version is just for testing.

## Building the source 

For prerequisites and detailed build instructions please read the [Installation Instructions](https://github.com/SmartMeshFoundation/Spectrum/wiki/Building-Specturm) on the wiki.

Building geth requires both a Go (version 1.7 or later) and a C compiler. You can install them using your favourite package manager. Once the dependencies are installed, run

    $ make all
    
or, to build the full suite of utilities:

    $ make all

## Run fast node to test geth

    $ ./build/bin/geth --testnet console

## Create new account

    > personal.newAccount()

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

## More works
    more functions is now under development, and will show you later.

