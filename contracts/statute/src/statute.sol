pragma solidity >=0.5.0 <=0.5.3;

contract Owned {

    /// `owner` is the only address that can call a function with this
    /// modifier
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    address public owner;

    /// @notice The Constructor assigns the message sender to be `owner`
    constructor() public {
        owner = msg.sender;
    }

    address newOwner = address(0);

    event OwnerUpdate(address _prevOwner, address _newOwner);

    ///change the owner
    function changeOwner(address _newOwner) public onlyOwner {
        require(_newOwner != owner);
        newOwner = _newOwner;
    }

    /// accept the ownership
    function acceptOwnership() public {
        require(msg.sender == newOwner);
        emit OwnerUpdate(owner, newOwner);
        owner = newOwner;
        newOwner = address(0);
    }
}


contract MeshBox is Owned {

    mapping(address => uint256) meshboxAddress;

    address[] meshboxList;

    function addAddress(address[] memory _owners, uint256 version) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            address _o = _owners[i];
            if (meshboxAddress[_o] == 0) {
                meshboxList.push(_o);
            }
            meshboxAddress[_o] = version;
        }
    }

    function getMeshboxList() view public returns (address[] memory) {
        return meshboxList;
    }


    function cleanMeshboxList(address addr) private {
        meshboxAddress[addr] = 0;
        for (uint i=0;i<meshboxList.length;i++) {
            if (meshboxList[i] == addr) {
                for (uint j=i;j<meshboxList.length-1;j++) {
                    meshboxList[j] = meshboxList[j+1];
                }
                i--;
                meshboxList.length -= 1;
            }
        }
    }

    function delAddress(address[] memory _owners) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            cleanMeshboxList(_owners[i]);
        }
    }

    function existAddress(address _owner) view public returns (uint256){
        return meshboxAddress[_owner];
    }
}


// Address and nodeid mapping
// used for reward txfee
contract Anmap is MeshBox {

    mapping(address => address[]) anmap;
    mapping(address => address) namap;

    //"0x70aEfe8d97EF5984B91b5169418f3db283F65a29",28,"0xa006360ed68f21443221354903f49dce733afaaeac3b35d420bb2154746c9592","0x6f31ccfa10b449531fd0fff78db52cc02edaabd2c5e9a6abb8fc1cd6df26f442"
    function bind(address nodePub, uint8 v, bytes32 r, bytes32 s) public {

        bytes32 hash = keccak256(abi.encodePacked(msg.sender));
        address signer = ecrecover(hash,v,r,s);

        require(nodePub == signer);
        //require(anmap[msg.sender] == address(0));
        require(namap[nodePub] == address(0));

        //anmap[msg.sender] = nodePub;
        namap[signer] = msg.sender;

        anmap[msg.sender].push(signer);

    }

    function unbindBySig(address nodePub, uint8 v, bytes32 r, bytes32 s) public {
        bytes32 hash = keccak256(abi.encodePacked(msg.sender));
        address signer = ecrecover(hash,v,r,s);
        require(nodePub == signer);
        address addr = namap[signer];
        require(addr != address(0));
        delete namap[signer];
        if (anmap[addr].length < 2) {
            delete anmap[addr];
        }else{
            for (uint i=0;i<anmap[addr].length;i++) {
                if (anmap[addr][i] == signer ){
                    for ( uint j = i;j<anmap[addr].length-1;j++ ){
                        anmap[addr][j] = anmap[addr][j+1];
                    }
                    anmap[addr].length -= 1;
                    break;
                }
            }
        }
    }

    function bindInfo(address addr) view public returns(address from, address[] memory nids) {
        from = namap[addr];
        if (from == address(0)) {
            from = addr;
            nids = anmap[addr];
        }else{
            nids = new address[](1);
            nids[0] = addr;
        }
    }

}
