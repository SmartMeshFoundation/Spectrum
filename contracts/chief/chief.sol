// Copyright 2018 The Smartmesh Chain Authors
// This file is part of the SMChain library.
pragma solidity ^0.4.19;

contract TribeChief {

    address private owner;
    //配置 >>>>
    int epoch = 50; // TODO 调试环境设置为 50 , 生产环境为 30000
    int blockPeriod = 15; // 出块时间
    mapping(address => bool) genesisSigner; //创世签名节点
    mapping(int => address) recents; //当前 epoch 最新的块的签名者集合
    //配置 <<<<
    uint constant singerLimit = 101;

    //用偏移量来找 signer 对应的 score
    struct Signers {
        address[] signerList; //签名人列表
        int[] scoreList; //签名人分数
        //只用于查询，没有其他用途
        mapping(address => int) scoremap;
        uint256[] numberList; //日志，记录本次列表变动在哪个块上
    }

    //签名人集合
    Signers signers;

    //删除时，要把 signerlist 、 scorelist 、 scoremap 都删一遍
    function deleteSigner(uint index) private {
        Signers self = signers;
        uint len = self.signerList.length;
        if (index >= len) return;

        // scoremap >>>>
        delete self.scoremap[self.signerList[index]];
        // scoremap <<<<
        // signerList >>>>
        for (uint i = index; i<len-1; i++) {
            self.signerList[i] = self.signerList[i+1];
        }
        delete self.signerList[len-1];
        self.signerList.length--;
        // signerList <<<<
        // scoreList >>>>
        for (uint i = index; i<len-1; i++) {
            self.scoreList[i] = self.scoreList[i+1];
        }
        delete self.scoreList[len-1];
        self.scoreList.length--;
        // scoreList <<<<
        // numberList >>>>
        for (uint i = index; i<len-1; i++) {
            self.numberList[i] = self.numberList[i+1];
        }
        delete self.numberList[len-1];
        self.numberList.length--;
        // numberList <<<<

    }

    function pushSigner(address signer, int score) public returns (uint, bool){
        Signers self = signers;
        if (self.signerList.length >= singerLimit) {
            return (0, false);
        }
        self.signerList.push(signer);
        self.scoreList.push(score);
        self.numberList.push(block.number);
        self.score[signer] = score;
        return (self.signerList.length, true);
    }

    function TribeChief() public {
        owner = msg.sender;
        // 维护世界和平的任务就交给你们了

        address g1 = 0xecce1549c2c803996e69930c97594525294de68f;
        address g2 = 0x922316aefd06ae15ec9a797c88c695eaa4c3c32b;
        address g3 = 0xec61e6c0c17930fe1d8538de2c0b25644c542632;
        address g4 = 0xadb3ea3ad356199206ca817b04fd668cc5196df2;
        genesisSigner[g1] = true;
        genesisSigner[g2] = true;
        genesisSigner[g3] = true;
        genesisSigner[g4] = true;
        pushSigner(g1, 3);
        pushSigner(g2, 3);
        pushSigner(g3, 3);
        pushSigner(g4, 3);
    }

    //生效了
    modifier apply(address _addr) {
        require(signers.score[_addr] > 0);
        _;
    }

    function update() public apply(msg.sender) {
        int score = signers.scoreList[0];
        if (score > 0) {
            score = score - 1;
            signers.scoreList[0] = score;
            signers.numberList[0] = block.number;
        } else {
            delete signers.signerList[0];
            delete signers.scoreList[0];
        }
    }

    function getStatus() constant returns (address[], int[], uint256[]) {
        return (signers.signerList, signers.scoreList, signers.numberList);
    }

}
