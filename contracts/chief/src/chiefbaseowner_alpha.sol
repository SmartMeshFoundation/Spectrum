pragma solidity >=0.5.0 <=0.5.3;

contract ChiefBaseInterface{
    function pocChangeOwner(address newOwner, uint256 num)   public;

}
/*
该合约有两个主要功能:
1. 依据规则增减常委会节点
2. 辅助下一个版本硬分叉升级Chief合约,到时候我们希望只升级Chief,以及ChiefBase合约,但是不升级poc合约.
*/
contract ChiefBaseOwner {
    ChiefBaseInterface base;
    constructor(address chiefBase) public {
        base=chiefBase;
    }
    function UpgradeToNewChiefVersion( address newChiefBase,uint256 num) public{
        //msg.sender保证不存在,这样在硬分叉的指定块强制调用这个合约.
        require(msg.sender==address(0xffffffffffffffffffffffffffffffffffffffff));
        base.pocChangeOwner(newChiefBase,num);
        base=newChiefBase;
    }
}