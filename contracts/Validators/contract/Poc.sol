// SPDX-License-Identifier: MIT
pragma solidity 0.8.1;

// import "@openzeppelin/contracts@4.5.0/token/ERC20/utils/SafeERC20.sol";

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract POC {

    struct minerInfo {
        address owner;//该矿工的owner地址
        uint256 amount;//抵押金额
        uint256 stop_block;//暂停出块时的区块号
        uint256 recover_block;//可恢复出块的区块号，用于被共识算法拉黑记录，如果是用户主动暂停则该值为0
        uint256 node_type;// 1 poa 2 pos
    }

    //矿工地址对应的状态信息
    mapping(address=>minerInfo) minerMap;

    //正常出块矿工地址列表
    address[] normalList;

    //normalList下标索引
    mapping (address => uint256) normalListIndex;

    //停止出块矿工地址列表
    address[] stopList;

    //stopList下标索引
    mapping (address => uint256) stopListIndex;

    //取回质押资金的最小等待区块数
    uint256 public constant withdrawWaitNumber = 20160;

    //质押金额
    uint256 public constant depositAmount = 5000000 ether;

    //meshToken
    using SafeERC20 for IERC20;
    IERC20 public meshToken;


    //存款
    event Deposit(address indexed miner, address indexed owner, uint256 amount);

    //暂停出块
    event Stop(address indexed miner);

    //启动出块
    event Start(address indexed miner);

    //提现
    event Withdraw(address indexed miner, address indexed owner, uint256 amount);


    modifier onlyPosNode(address minerAddress) {
        require(minerMap[minerAddress].node_type == 2, "POS node only");
        _;
    }

    modifier onlyPoaNode(address minerAddress) {
        require(minerMap[minerAddress].node_type == 1, "POA node only");
        _;
    }

    // 质押mesh，成为一个POS节点
    function deposit(bytes32 _r, bytes32 _s, uint8 _v) external {

        meshToken.safeTransferFrom(msg.sender, address(this), depositAmount);

        //交易发起人是否得到矿工授权
        bytes32 h = keccak256(abi.encodePacked(msg.sender));

        //获取矿工地址
        address minerAddress = ecrecover(h,_v,_r,_s);

        addPosNode(minerAddress, depositAmount);
    }


    //添加POS节点
    function addPosNode(address minerAddress, uint256 amount) internal {
        addNode(minerAddress, 2, amount);
    }


    //添加POA节点
    function addPoaNode(address minerAddress) internal {
        addNode(minerAddress, 1, 0);
    }


    //添加节点
    function addNode(address minerAddress, uint256 nodeType, uint256 amount) private {

        require (minerAddress != address(uint160(0)), "Invalid address");

        //每个节点只能质押一次
        require (minerMap[minerAddress].node_type == 0, "Duplicate addresses");

        //记录节点信息
        minerMap[minerAddress] = minerInfo(msg.sender, amount, 0, 0, nodeType);

        //添加到正常出块列表
        normalList.push(minerAddress);

        //添加索引
        normalListIndex[minerAddress] = normalList.length - 1;

        //添加事件日志
        emit Deposit(minerAddress, msg.sender, amount);

    }


    //用户主动暂停出块
    function stop(address minerAddress) external onlyPosNode(minerAddress) {

        //是否是地址owner
        require(minerMap[minerAddress].owner == msg.sender, "Owner only");

        require(addStop(minerAddress) == 1, "Stop error");
    }

    //停止出块并停用该账户，在 recoverBlock 区块号之后可以启动
    function stopAndDisable(address minerAddress, uint256 recoverBlock) internal onlyPosNode(minerAddress){
        addStop(minerAddress);
        minerMap[minerAddress].recover_block = recoverBlock;
    }


    //结束暂停, 恢复出块
    function start(address minerAddress) external onlyPosNode(minerAddress){

        //是否是地址owner
        require(minerMap[minerAddress].owner == msg.sender, "Owner only");

        require(addStart(minerAddress) == 1, "Start error");
    }


    //添加暂停出块入口 -1 不存在此地址 0 此地址已经在暂停列表 1 设置成功
    function addStop(address minerAddress) private returns(int8) {

        //是否存在此地址
        if (minerMap[minerAddress].node_type == 0){
            return -1;
        }

        //是否已经在暂停出块
        if (minerMap[minerAddress].stop_block != 0){
            return 0;
        }

        //修改状态
        minerMap[minerAddress].stop_block = block.number;

        //获取位置下标
        uint256 index = normalListIndex[minerAddress];

        //使用最后一个地址覆盖
        normalList[index] = normalList[normalList.length - 1];

        // normalList.length -= 1;
        normalList.pop();

        //更新normalList索引
        delete normalListIndex[minerAddress];
        if (index < normalList.length){
            normalListIndex[normalList[index]] = index;
        }

        //加入暂停出块列表
        stopList.push(minerAddress);

        //更新stopList索引
        stopListIndex[minerAddress] = stopList.length - 1;

        //添加事件日志
        emit Stop(minerAddress);

        return 1;
    }


    //添加启动出块入口 -1 不存在此地址 0 此地址已经在出块列表或不被允许启动 1 设置成功
    function addStart(address minerAddress) private returns(int8) {

        //是否存在此地址
        if (minerMap[minerAddress].node_type == 0){
            return -1;
        }

        //黑名单时间是否已过
        require(minerMap[minerAddress].recover_block < block.number, "Waiting for recovery");

        //是否已经在正常出块
        if (minerMap[minerAddress].stop_block == 0){
            return 0;
        }

        //修改状态
        minerMap[minerAddress].stop_block = 0;
        minerMap[minerAddress].recover_block = 0;

        //获取位置下标
        uint256 index = stopListIndex[minerAddress];

        //使用最后一个地址覆盖
        stopList[index] = stopList[stopList.length - 1];
        stopList.pop();

        //更新stopList索引
        delete stopListIndex[minerAddress];
        if (index < stopList.length){
            stopListIndex[stopList[index]] = index;
        }

        //加入出块列表
        normalList.push(minerAddress);

        //更新normalList索引
        normalListIndex[minerAddress] = normalList.length - 1;

        //添加事件日志
        emit Start(minerAddress);

        return 1;
    }


    //移除POS节点（质押资金提现）
    function removePosNode(address minerAddress) external onlyPosNode(minerAddress) {

        //是否是地址owner
        require (minerMap[minerAddress].owner == msg.sender, "Owner only");

        //是否处于暂停状态
        require (minerMap[minerAddress].stop_block > 0, "Must stop");

        //暂停等待期是否完成
        require (block.number - minerMap[minerAddress].stop_block > withdrawWaitNumber, "Waiting period");

        //黑名单时间是否已过
        require (minerMap[minerAddress].recover_block < block.number, "Waiting for recovery");

        removeNode(minerAddress);
    }

    //移除POA节点
    function removePoaNode(address minerAddress) internal onlyPoaNode(minerAddress) {
        removeNode(minerAddress);
    }


    //移除节点
    function removeNode(address minerAddress) private {

        uint256 amount = minerMap[minerAddress].amount;

        //删除记录
        delete minerMap[minerAddress];

        //获取位置下标
        uint256 index = stopListIndex[minerAddress];

        //使用最后一个地址覆盖
        stopList[index] = stopList[stopList.length - 1];
        stopList.pop();

        //更新stopList索引
        delete stopListIndex[minerAddress];
        if (index < stopList.length){
            stopListIndex[stopList[index]] = index;
        }

        //添加事件日志
        emit Withdraw(minerAddress, msg.sender, amount);

        //质押资金返回
        if(amount > 0){
            meshToken.safeTransfer(msg.sender, amount);
        }
    }


    // 获取矿工信息，当amount 或 minerOwner为空时 表示不存在记录，number为暂停块号，大于0则已经处于暂停状态
    function getValidatorState(address _addr) public view returns(uint256, uint256, uint256, address, uint256) {

        uint256 amount = minerMap[_addr].amount;
        uint256 stopBlock = minerMap[_addr].stop_block;
        uint256 recoverBlock = minerMap[_addr].recover_block;
        address minerOwner = minerMap[_addr].owner;
        uint256 nodeType = minerMap[_addr].node_type;

        return (amount, stopBlock, recoverBlock, minerOwner, nodeType);
    }


    //获取所有暂停出块地址
    function getStopList() public view returns(address[] memory) {
        return stopList;
    }


    //获取所有正常出块地址
    function getNormalList() public view returns(address[] memory) {
        return normalList;
    }

    //获取所有信息
    function getAllValidators() public view returns(
        address[] memory,
        uint256[] memory,
        uint256[] memory,
        uint256[] memory,
        uint256[] memory,
        address[] memory
    ){
        uint256 slen = stopList.length;
        uint256 nlen = normalList.length;
        uint256 len = slen + nlen;

        //地址列表
        address[] memory minerList;
        minerList = new address[](len);

        //质押资金列表
        uint256[] memory amountList;
        amountList = new uint256[](len);

        //暂停区块号列表
        uint256[] memory stopBlockList;
        stopBlockList = new uint256[](len);

        //节点类型列表
        uint256[] memory typeList;
        typeList = new uint256[](len);

        //恢复时间列表
        uint256[] memory recoverBlockList;
        recoverBlockList = new uint256[](len);

        //owner列表
        address[] memory ownerList;
        ownerList = new address[](len);

        for (uint256 i = 0; i < slen; i++){
            minerList[i]  = stopList[i];
            amountList[i] = minerMap[stopList[i]].amount;
            stopBlockList[i] = minerMap[stopList[i]].stop_block;
            typeList[i] = minerMap[stopList[i]].node_type;
            recoverBlockList[i]  = minerMap[stopList[i]].recover_block;
            ownerList[i]  = minerMap[stopList[i]].owner;
        }

        for (uint256 j = 0; j < nlen; j++ ){
            minerList[slen + j]  = normalList[j];
            amountList[slen + j] = minerMap[normalList[j]].amount;
            stopBlockList[slen + j] = minerMap[normalList[j]].stop_block;
            typeList[slen + j] = minerMap[normalList[j]].node_type;
            recoverBlockList[slen + j]  = minerMap[normalList[j]].recover_block;
            ownerList[slen + j]  = minerMap[normalList[j]].owner;
        }
        return (minerList, amountList, stopBlockList, typeList, recoverBlockList, ownerList);
    }
}
