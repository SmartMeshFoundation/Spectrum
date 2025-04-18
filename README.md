# Spectrum

![光谱logo2(1).png](https://upload-images.jianshu.io/upload_images/528413-0c926281c1d94539.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/440)





Spectrum is an [Ethereum-compatible](https://github.com/ethereum/go-ethereum) project. It uses a new consensus and new block reward for SmartMesh ecosystem devices and IOT. And you can view the transactions on the [
BlockChain Browser Address](https://spectrum.pub).


Since the list of signers is 17, it is recommended that the confirmation number of general transfer transaction block be set to 17 (one round), and that of exchange block be set to 34 (two rounds).

## List of Chain ID's:
| Chain(s)    |  CHAIN_ID  | 
| ----------  | :-----------:| 
| mainnet     | 20180430     | 
| testnet     | 2022         | 
| devnet      | 4            | 

## Warning

We suggest that the GasPrice should not be less than 18Gwei, otherwise the transaction may not be packaged into the block.

## Build the source 

Building Spectrum requires both a Go (version 1.15 or later) and a C compiler. You can install them using your favourite package manager. And you can view the detail installation and running steps on this [page](https://github.com/SmartMeshFoundation/Spectrum/wiki/Building-Specturm).

## Run node 

```bash
$ ./build/bin/smc console
```
    
## Create new account

```bash
# Users can create new account:
> personal.newAccount()
```

## Get your own miner id

```bash
# Every node has it's own miner id, you can run getMiner() function to get that id:
> tribe.getMiner() 
```
    
## Bind your own miner id to wallet address

```bash
# Users can bind their miner ID to a wallet address:
> tribe.bind("account","passwd") 
    
# Or Users can only generate binding signatures at the terminal:
> tribe.bindSign("account") 
```

## Deposit smt to POC

```bash
# Users can become miner by deposit smt:
> tribe.pocDeposit("account","passwd") 
```

## Start mining

```bash
# Users can start mining or resume it:
> tribe.pocStart("account","passwd") 
```

## Stop mining

```bash
Users can stop mining:
> tribe.pocStop("account","passwd") 
```

## Withdraw smt

```bash
# Users can withdraw smt:
> tribe.pocWithdraw("account","passwd")   
```

## More functions

```bash
# Users can input tribe to view:    
> tribe
```
    
## Security-related 
  
### Encrypt your nodekey

```bash
$ ./build/bin/smc security --passwd
```
     
### Decrypt your nodekey

```bash
$ ./build/bin/smc security --unlock
```

