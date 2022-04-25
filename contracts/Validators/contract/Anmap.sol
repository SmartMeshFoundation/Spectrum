// SPDX-License-Identifier: MIT
pragma solidity 0.8.1;

// Address and nodeid mapping
// used for reward txfee
contract Anmap {

    mapping(address => address[]) anmap;
    mapping(address => address) namap;

    //"0x70aEfe8d97EF5984B91b5169418f3db283F65a29",28,"0xa006360ed68f21443221354903f49dce733afaaeac3b35d420bb2154746c9592","0x6f31ccfa10b449531fd0fff78db52cc02edaabd2c5e9a6abb8fc1cd6df26f442"
    function bind(address nodePub, uint8 v, bytes32 r, bytes32 s) external {

        bytes32 hash = keccak256(abi.encodePacked(msg.sender));
        address signer = ecrecover(hash, v, r, s);

        require(nodePub == signer);
        //require(anmap[msg.sender] == address(0));
        require(namap[nodePub] == address(0));

        //anmap[msg.sender] = nodePub;
        namap[signer] = msg.sender;

        anmap[msg.sender].push(signer);

    }

    function unbindBySig(address nodePub, uint8 v, bytes32 r, bytes32 s) external {
        bytes32 hash = keccak256(abi.encodePacked(msg.sender));
        address signer = ecrecover(hash, v, r, s);
        require(nodePub == signer);
        address addr = namap[signer];
        require(addr != address(0));
        delete namap[signer];
        if (anmap[addr].length < 2) {
            delete anmap[addr];
        } else {
            for (uint i = 0; i < anmap[addr].length; i++) {
                if (anmap[addr][i] == signer) {
                    for (uint j = i; j < anmap[addr].length - 1; j++) {
                        anmap[addr][j] = anmap[addr][j + 1];
                    }
                    anmap[addr].pop();
                    break;
                }
            }
        }
    }

    function bindInfo(address addr) view external returns (address from, address[] memory nids) {
        from = namap[addr];
        if (from == address(0)) {
            from = addr;
            nids = anmap[addr];
        } else {
            nids = new address[](1);
            nids[0] = addr;
        }
    }

}
