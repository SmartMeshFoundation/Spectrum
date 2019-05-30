pragma solidity ^0.5.3;

contract POC{
    

    struct minerInfo {
        address owner;//该矿工的出资人地址
        uint256 amount;//抵押金额
        uint256 stop_block;//暂停出块时停止出块块号
    }

    //矿工地址对应的状态信息
    mapping(address=>minerInfo) minerMap;


    //正常出块矿工地址列表
    address[] normalList;

    //停止出块矿工地址列表
    address[] stopList;
    
    //合约owner地址
    address[] owner;

    //最小存放金额
    uint256 public minDeposit;

    //提现最小等待区块
    uint256 public withdrawWaitNumber;

    //新的合约owner生效区块
    uint256 public newOwnerEffectiveNumber;

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

    constructor(address _addr, uint256 _amount, uint256 _number) public {
        minDeposit = _amount;
        owner.push(_addr);
        withdrawWaitNumber = _number;
    }
    

    //存款进入出块人列表    
    function deposit(bytes32 _r, bytes32 _s, uint8 _v) public payable {
        
        //存款金额至少100w
        require (msg.value >= minDeposit);

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

        //移出正常出块列表
        uint256 len = normalList.length;
        for(uint i = 0; i < len - 1; i++){
            if (minerAddress == normalList[i]){
                normalList[i] = normalList[len - 1];
                break;
            }
        }
        normalList.length -= 1;

        //加入暂停出块列表
        stopList.push(minerAddress);

        return 1;
    }


    //添加启动出块入口 -1 不存在此地址 0 此地址已经在出块列表 1 设置成功
    function addStart(address minerAddress) private returns(int8) {
        
        //是否存在此地址
        if (minerMap[minerAddress].amount == 0){
            return -1;
        }

        //是否已经在正常出块
        if (minerMap[minerAddress].stop_block == 0){
            return 0;
        }

        //修改状态
        minerMap[minerAddress].stop_block = 0;

        //移出暂停列表
        uint256 len = stopList.length;
        for(uint i = 0; i < len - 1; i++){
            if (minerAddress == stopList[i]){
                stopList[i] = stopList[len - 1];
                break;
            }
        }
        stopList.length -= 1;

        //加入出块列表
        normalList.push(minerAddress);

        return 1;
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
        
        uint256 amount = minerMap[minerAddress].amount;

        //删除记录
        delete minerMap[minerAddress];

        //移出暂停列表
        uint256 len = stopList.length;
        for(uint i = 0; i < len - 1; i++){
            if (minerAddress == stopList[i]){
                stopList[i] = stopList[len - 1];
                break;
            }
        }
        stopList.length -= 1;        

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


    //获取所有owner地址
    function getOwnerList() public view returns(address[] memory) {
        return owner;
    }
    

    //获取所有信息
    function getAll() public view returns(
        address[] memory,
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
        uint256[] memory blockList;
        blockList = new uint256[](len);

        //owner列表
        address[] memory ownerList;
        ownerList = new address[](len);

        for (uint256 i = 0; i < slen; i++){
            minerList[i]  = stopList[i];
            amountList[i] = minerMap[stopList[i]].amount;
            blockList[i]  = minerMap[stopList[i]].stop_block;
            ownerList[i]  = minerMap[stopList[i]].owner;
        }

        for (uint256 j = 0; j < nlen; j++ ){
            minerList[slen + j]  = normalList[j];
            amountList[slen + j] = minerMap[normalList[j]].amount;
            blockList[slen + j]  = minerMap[normalList[j]].stop_block;
            ownerList[slen + j]  = minerMap[normalList[j]].owner;
        }
        return (minerList, amountList, blockList, ownerList);
    }
}
