pragma solidity >=0.7.0 <0.9.0;

// testnet : 0xa44B4ff09825f8c21fc4ad7FAA125a0d6238d0fd
contract evm_tests {

    function test_selfbalance()public view returns (uint256) {
        return getThis().balance;
    }

    function test_shl() public view returns (bool) {
        int256 i = 100;
        //i = i << 2;
        assembly{
            i := shl( 2, i )
        }
        return i == 400;
    }

    function test_shr() public view returns (bool) {
        int256 i = 100;
        //i = i >> 2;
        assembly{
            i := shr( 2, i )
        }
        return i == 25;
    }

    function test_sar() public view returns (bool) {
        int256 i = -100;
        assembly{
            i := sar( 2, i )
        }
        return i == -25;
    }

    function test_extcodehash() public view returns (bool) {
        address a = getThis();
        bytes32 codeHash;
        assembly {
            codeHash := extcodehash(a)
        }
        return codeHash != 0x0;
    }

    function test_chainid() external view returns (uint256) {
        uint256 id;
        assembly {
            id := chainid()
        }
        return id;
    }

    function test_create2() public payable {
        address _owner = getThis();
        uint _salt = getNumber();
        bytes memory bytecode = abi.encodePacked(type(TestContract).creationCode, abi.encode(_owner));
        address addr;
        assembly {
            addr := create2(
                callvalue(),
                add(bytecode, 0x20),
                mload(bytecode),
                _salt)
            if iszero(extcodesize(addr)) {
                revert(0, 0)
            }
        }
        emit Create2_ret( addr, _owner, _salt, addr == getAddress(_owner,_salt));
    }


    constructor() payable {}

    event Create2_ret(address addr, address owner, uint salt, bool ret);

    function getThis() public view returns (address){
        return address(this);
    }

    function getNumber() public view returns (uint256){
        return block.number;
    }

    function getAddress(address _owner ,uint _salt) public view returns (address) {
        bytes memory bytecode = abi.encodePacked(type(TestContract).creationCode, abi.encode(_owner));
        bytes32 hash = keccak256(
            abi.encodePacked(bytes1(0xff), address(this), _salt, keccak256(bytecode))
        );
        return address(uint160(uint(hash)));
    }
}

contract TestContract {
    address public owner;

    constructor(address _owner) payable {
        owner = _owner;
    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }

}