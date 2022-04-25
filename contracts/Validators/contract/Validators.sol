// SPDX-License-Identifier: MIT
pragma solidity 0.8.1;

import "./Poc.sol";
import "./Anmap.sol";
// import "@openzeppelin/contracts@4.5.0/utils/math/SafeMath.sol";
// import "@openzeppelin/contracts@4.5.0/access/Ownable.sol";

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract Validators is POC,Anmap,Ownable{

	uint256 constant epoch = 21;//出块周期，21个出块人为一个周期
	uint256 constant disableNumber = 2880;//黑名单锁定区块时间
    bool public initialized;//初始化标记

	modifier onlyMiner() virtual {
		require(msg.sender == block.coinbase, "Miner only");
		_;
	}


	//当前出块节点可调用
	function punishValidator(address minerAddress) external onlyMiner {

		//只能禁用POS节点
		if(minerMap[minerAddress].node_type == 2){
			stopAndDisable(minerAddress, block.number + disableNumber);
		}
	}

	//在当前出块周期的最后一个区块充值签名人列表
	function getNewValidators(uint256 randNum) external view returns(address[] memory){

		uint256 validatorLen;

		if(normalList.length >= epoch){
			validatorLen = epoch;
		} else {
			validatorLen = normalList.length;
		}

		address[] memory validators;
        validators = new address[](validatorLen);

		uint256 index = randNum % normalList.length;
		uint256 total;

		//添加出块人列表
		for(uint256 i = index; i < normalList.length && total < validatorLen; i++){
			validators[total] = normalList[i];
			total++;
		}

		//如果数组尾部出块人不足，则用数组头部补充
		for(uint256 i = 0; i < index && total < validatorLen; i++){
			validators[total] = normalList[i];
			total++;
		}

		return validators;
	}

	//添加POA节点
	function ownerAddPoaNode(address[] calldata minerAddressList) external onlyOwner {
		for(uint256 i = 0; i < minerAddressList.length; i++){
			addPoaNode(minerAddressList[i]);
		}
	}

	//移除POA节点
	function ownerRemovePoaNode(address[] calldata minerAddressList) external onlyOwner {
		for(uint256 i = 0; i < minerAddressList.length; i++){
			removePoaNode(minerAddressList[i]);
		}
	}

    //初始化POA节点和Owner
    function initialize(address[] calldata minerAddressList, address newOwner) external {
    	require(initialized == false);
        initialized = true;

        for(uint256 i = 0; i < minerAddressList.length; i++){
			addPoaNode(minerAddressList[i]);
		}
		meshToken = IERC20(0x0000000000000000000000000000000000002000);
        _transferOwnership(newOwner);
    }
}

