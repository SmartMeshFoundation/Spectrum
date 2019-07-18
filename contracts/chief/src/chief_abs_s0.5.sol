pragma solidity >=0.5.0 <=0.5.3;

contract Chief {

    function update(address volunteer) public;

    function version() public view returns (string memory);

    function getSignerLimit() public view returns (uint);

    function getEpoch() public view returns (uint);

    function getVolunteerLimit() public view returns (uint);

    function getStatus() public view returns (
        address[] memory signerList,
        address[] memory blackList,
        uint[] memory scoreList,
        uint[] memory numberList,
        uint totalVolunteer,
        uint number
    );

    function filterVolunteer(address[] memory volunteers) public view returns (uint[] memory result);

    function getVolunteers() public view returns (
        address[] memory volunteerList,
        uint[] memory weightList,
        uint length
    );
}