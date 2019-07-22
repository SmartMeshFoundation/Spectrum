pragma solidity >=0.5.0 <0.6.0;

contract POC_1_0_0 {

    struct minerInfo {
        address owner;//该矿工的owner地址
        uint256 amount;//抵押金额
        uint256 stop_block;//暂停出块时的区块号
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

    //黑名单列表
    address[] blackList;

    //blackList下标索引
    mapping (address => uint256) blackListIndex;

    //锁定名单列表
    address[] lockList;

    //lockList下标索引
    mapping (address => uint) lockListIndex;


    //合约owner地址
    address[] owner;

    //初始最小存放金额
    uint256 public initMinDeposit;

    //启动挖矿块号
    uint256 public initBlockNumber;

    //存款减半间隔区块数
    uint256 public depositHalveLimitNumber;

    //提现最小等待区块数
    uint256 public withdrawWaitNumber;

    //新的合约owner生效区块数
    uint256 public newOwnerEffectiveNumber;


    //存款
    event Deposit(address indexed miner, address indexed owner, uint256 amount);

    //暂停出块
    event Stop(address indexed miner);

    //启动出块
    event Start(address indexed miner);

    //提现
    event Withdraw(address indexed miner, address indexed owner, uint256 amount);

    modifier onlyOwner() {

        uint256 len = owner.length;
        address currentOwner;

        //最新owner
        currentOwner = owner[len - 1];

        //检查最新owner是否未生效
        if (len > 1 && block.number < newOwnerEffectiveNumber){
            currentOwner = owner[len - 2];//使用上一个owner
        }

        require(msg.sender == currentOwner);
        _;
    }

    constructor(address _addr, uint256 _amount, uint256 _waitNumber, uint256 _limitNumber) public {

        initMinDeposit = _amount;
        owner.push(_addr);
        withdrawWaitNumber = _waitNumber;
        depositHalveLimitNumber = _limitNumber;
        initBlockNumber = block.number;
    }


    //每经过 depositHalveLimitNumber 区块数以后，质押金额即可减半,减半三次后不再减半
    function minDepositAmount() public view returns(uint256) {

        // 减半次数
        uint256 num = (block.number - initBlockNumber) / depositHalveLimitNumber;

        if (num > 3){
            num = 3;
        }
        return initMinDeposit / (2 ** (num));
    }


    //存款进入出块人列表
    function deposit(bytes32 _r, bytes32 _s, uint8 _v) public payable {

        //满足最小存款金额
        require (msg.value >= minDepositAmount());

        //交易发起人是否得到矿工授权
        bytes32 h = keccak256(abi.encodePacked(msg.sender));

        //获取矿工地址
        address minerAddress = ecrecover(h,_v,_r,_s);

        require (minerAddress != address(uint160(0)));

        //只能质押一次
        require (minerMap[minerAddress].amount == 0);

        //记录出块信息
        minerMap[minerAddress] = minerInfo(msg.sender, msg.value, 0);

        //添加到正常出块列表
        normalList.push(minerAddress);

        //添加索引
        normalListIndex[minerAddress] = normalList.length - 1;

        //添加事件日志
        emit Deposit(minerAddress, msg.sender, msg.value);
    }


    //暂停出块 准备提现
    function stop(address minerAddress) public {

        //是否是地址owner
        require (minerMap[minerAddress].owner == msg.sender);

        require(addStop(minerAddress) == 1);
    }


    //结束暂停出块 恢复出块
    function start(address minerAddress) public {

        //是否是地址owner
        require (minerMap[minerAddress].owner == msg.sender);

        require(addStart(minerAddress) == 1);
    }



    //owner设置暂停出块
    function ownerStop(address minerAddress) public onlyOwner returns(int8) {

        return addStop(minerAddress);
    }


    //添加暂停出块入口 -1 不存在此地址 0 此地址已经在暂停列表 1 设置成功
    function addStop(address minerAddress) private returns(int8) {

        //是否存在此地址
        if (minerMap[minerAddress].amount == 0){
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

        normalList.length -= 1;

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
        if (minerMap[minerAddress].amount == 0){
            return -1;
        }

        //是否在黑名单或者锁定状态
        if (blackAndLockStatus(minerAddress) > 0){
            return 0;
        }

        //是否已经在正常出块
        if (minerMap[minerAddress].stop_block == 0){
            return 0;
        }

        //修改状态
        minerMap[minerAddress].stop_block = 0;

        //获取位置下标
        uint256 index = stopListIndex[minerAddress];

        //使用最后一个地址覆盖
        stopList[index] = stopList[stopList.length - 1];
        stopList.length -= 1;

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


    //获取黑名单和锁定状态
    function blackAndLockStatus(address minerAddress) public view returns(int8) {

        int8 value = 0;

        //是否在黑名单内
        if (blackList.length > 0 && blackList[blackListIndex[minerAddress]] == minerAddress){
            value += 1;
        }

        //是否在锁定名单内
        if (lockList.length > 0 && lockList[lockListIndex[minerAddress]] == minerAddress){
            value += 2;
        }

        return value;
    }


    //owner 新增地址到黑名单列表
    function ownerPushBlackList(address minerAddress) public onlyOwner {

        //是否已经在黑名单列表
        if (blackList.length > 0 && blackList[blackListIndex[minerAddress]] == minerAddress){
            return;
        }

        //加入黑名单列表
        blackList.push(minerAddress);

        //更新黑名单列表索引
        blackListIndex[minerAddress] = blackList.length - 1;
    }


    //owner 从黑名单列表剔除一个地址 -1 不存在此地址 1 设置成功
    function ownerPopBlackList(address minerAddress) public onlyOwner returns(int8){

        //获取位置下标
        uint256 index = blackListIndex[minerAddress];

        //是否存在此地址
        if (blackList.length < 1 || blackList[index] != minerAddress){
            return -1;
        }

        //使用最后一个地址覆盖
        blackList[index] = blackList[blackList.length - 1];
        blackList.length -= 1;

        //更新blackList索引
        delete blackListIndex[minerAddress];
        if (index < blackList.length){
            blackListIndex[blackList[index]] = index;
        }

        return 1;
    }


    //owner 清空黑名单列表
    function ownerEmptyBlackList() public onlyOwner {

        //清空索引
        for(uint256 i = 0; i < blackList.length; i++){
            delete blackListIndex[blackList[i]];
        }

        //删除数据
        blackList.length = 0;

    }


    //owner 新增地址到锁定列表
    function ownerPushLockList(address minerAddress) public onlyOwner {

        //是否已经在锁定列表
        if (lockList.length > 0 && lockList[lockListIndex[minerAddress]] == minerAddress){
            return;
        }

        //加入锁定列表
        lockList.push(minerAddress);

        //更新锁定列表索引
        lockListIndex[minerAddress] = lockList.length - 1;

    }


    //owner 从锁定列表中删除一个地址 -1 不存在此地址 1 设置成功
    function ownerPopLockList(address minerAddress) public onlyOwner returns(int8){

        //获取位置下标
        uint256 index = lockListIndex[minerAddress];

        //是否存在此地址
        if (lockList.length < 1 || lockList[index] != minerAddress){
            return -1;
        }

        //使用最后一个覆盖
        lockList[index] = lockList[lockList.length - 1];
        lockList.length -= 1;

        //更新locklist索引
        delete lockListIndex[minerAddress];
        if (index < lockList.length){
            lockListIndex[lockList[index]] = index;
        }

        return 1;

    }


    //owner 设置启动挖矿块号
    function ownerSetInitBlockNumber (uint256 _number) onlyOwner public {
        initBlockNumber = _number;
    }



    //修改owner
    function changeOwner(address _newOwner, uint256 _number) onlyOwner public {

        //覆盖未生效的owner
        if (block.number < newOwnerEffectiveNumber){
            owner[owner.length - 1] = _newOwner;
        } else {
            owner.push(_newOwner);
        }
        newOwnerEffectiveNumber = _number;
    }


    //质押资金提现
    function withdraw(address minerAddress) public {

        //是否是地址owner
        require (minerMap[minerAddress].owner == msg.sender);

        //是否处于暂停状态
        require (minerMap[minerAddress].stop_block > 0);

        //暂停等待期是否完成
        require (block.number - minerMap[minerAddress].stop_block > withdrawWaitNumber);

        //是否在黑名单或者锁定状态
        require (blackAndLockStatus(minerAddress) == 0);

        uint256 amount = minerMap[minerAddress].amount;

        //删除记录
        delete minerMap[minerAddress];

        //获取位置下标
        uint256 index = stopListIndex[minerAddress];

        //使用最后一个地址覆盖
        stopList[index] = stopList[stopList.length - 1];
        stopList.length -= 1;

        //更新stopList索引
        delete stopListIndex[minerAddress];
        if (index < stopList.length){
            stopListIndex[stopList[index]] = index;
        }

        //添加事件日志
        emit Withdraw(minerAddress, msg.sender, amount);

        //资金返回
        msg.sender.transfer(amount);
    }


    //提出大于最小质押金额的资金
    function withdrawSurplus(address minerAddress) public {
        //是否是地址owner
        require (minerMap[minerAddress].owner == msg.sender);

        uint256 amount = minerMap[minerAddress].amount - minDepositAmount();

        require (amount > 0);

        minerMap[minerAddress].amount = minerMap[minerAddress].amount - amount;

        //添加事件日志
        emit Withdraw(minerAddress, msg.sender, amount);

        //资金返回
        msg.sender.transfer(amount);
    }



    // 获取矿工信息，当amount 或 minerOwner为空时 表示不存在记录，number为暂停块号，大于0则已经处于暂停状态
    function minerStatus(address _addr) public view returns(uint256, uint256, address) {

        uint256 amount = minerMap[_addr].amount;
        uint256 number = minerMap[_addr].stop_block;
        address minerOwner = minerMap[_addr].owner;

        return (amount, number, minerOwner);
    }


    //获取所有暂停出块地址
    function getStopList() public view returns(address[] memory) {
        return stopList;
    }


    //获取所有正常出块地址
    function getNormalList() public view returns(address[] memory) {
        return normalList;
    }


    //获取所有黑名单地址
    function getBlackList () public view returns(address[] memory) {
        return blackList;
    }


    //获取所有锁定名单地址
    function getLockList () public view returns(address[] memory) {
        return lockList;
    }


    //获取所有owner地址
    function getOwnerList() public view returns(address[] memory) {
        return owner;
    }


    //获取所有信息
    function getAll() public view returns(
        address[] memory,
        uint256[] memory,
        uint256[] memory,
        address[] memory,
        uint256[] memory
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
        uint256[] memory blockList;
        blockList = new uint256[](len);

        //owner列表
        address[] memory ownerList;
        ownerList = new address[](len);

        //黑名单状态
        uint256[] memory blackStatusList;
        blackStatusList = new uint256[](len);

        for (uint256 i = 0; i < slen; i++){
            minerList[i]  = stopList[i];
            amountList[i] = minerMap[stopList[i]].amount;
            blockList[i]  = minerMap[stopList[i]].stop_block;
            ownerList[i]  = minerMap[stopList[i]].owner;
            if (blackList.length > 0 && blackList[blackListIndex[stopList[i]]] == stopList[i]){
                blackStatusList[i] = 1;
            } else {
                blackStatusList[i] = 0;
            }
        }

        for (uint256 j = 0; j < nlen; j++ ){
            minerList[slen + j]  = normalList[j];
            amountList[slen + j] = minerMap[normalList[j]].amount;
            blockList[slen + j]  = minerMap[normalList[j]].stop_block;
            ownerList[slen + j]  = minerMap[normalList[j]].owner;
            if (blackList.length > 0 && blackList[blackListIndex[normalList[j]]] == normalList[j]){
                blackStatusList[slen + j] = 1;
            } else {
                blackStatusList[slen + j] = 0;
            }
        }
        return (minerList, amountList, blockList, ownerList, blackStatusList);
    }
}
