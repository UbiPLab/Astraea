pragma solidity >=0.4.22 <0.7.0;
pragma experimental ABIEncoderV2;
contract DoSC{
    mapping(bytes32=>don) Don;
    mapping(bytes32=>uint8) DonIsEmpty;
    mapping(address=>uint256) Pledger; 
    mapping(address=>uint256) DepositTable;



    struct Donate_his{
        uint256 comx;
        uint256 comy;
        uint256 tokx;
        uint256 toky;
        bytes32 hres;
        bytes32 hcl;
        uint256 dt;
    }
    Donate_his[] donate_his_lis;
    struct Deliver1_his{
        uint256 pkcoix;
        uint256 pkcoiy;
        bytes32 hclj;
        uint256 dp;
        uint256 dt;
    }
    Deliver1_his[] deliver1_his_lis;
    struct Store_his{
        uint256 pkcoix;
        uint256 pkcoiy;
        bytes32 hclj;
        uint256 dt;
    }
    Store_his[] store_his_lis;
    struct Ship_his{
        uint256 pkcoix;
        uint256 pkcoiy;
        bytes32 hclj;
        uint256 dp;
    }
    Ship_his[] ship_his_lis;
    struct Dis_his{
        uint256 pkcoix;
        uint256 pkcoiy;
        bytes32 hclj;
        uint256 cmqx;
        uint256 cmqy;
        uint256 tokqx;
        uint256 tokqy;
        uint256 dt;
    }
    Dis_his[] dis_his_lis;
    struct Deliver2_his{
        uint256 pkcoix;
        uint256 pkcoiy;
        bytes32 hclj;
        uint256 dp;
        uint256 dt;
    }
    Deliver2_his[] deliver2_his_lis;
    struct Receive_his{
        bytes32 hfbi;
        uint256 dt;
        bytes32 hcl;
    }
    Receive_his[] receive_his_lis;
    struct don{
        State donstate;
        uint256 deposit;
        address[] DepositList;
    }

    enum State{
        submitted,
        delivery,
        stored,
        shipped,
        delivered,
        received
    }

    function test_add_dp(address a,uint256 k) private{
        Pledger[a]=k;
    }


    function Donate(uint256 comx,uint256 comy,uint256 tokx,uint256 toky,bytes32 hres,bytes32 hcl ,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender,
            "Incorrect"
        );
        Donate_his memory d=Donate_his(comx,comy,tokx,toky,hres,hcl,dt);
        Transfer("Donate",hcl,0,msg.sender);
        donate_his_lis.push(d);
    }


    function Deliver1(uint256 pkcoix,uint256 pkcoiy, bytes32 hclj, uint256 dp,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender,
            "Incorrect"
        );
        test_add_dp(msg.sender,10000000000);
        Deliver1_his memory d=Deliver1_his(pkcoix,pkcoiy,hclj,dp,dt);
        Pledger[msg.sender]-=dp;
        Transfer("Deliver1",hclj,dp,msg.sender);
        deliver1_his_lis.push(d);
    }

    function Store(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender,
            "Incorrect"
        );
        Store_his memory s=Store_his(pkcoix,pkcoiy,hclj,dt);
        Transfer("Store",hclj,0,msg.sender);
        store_his_lis.push(s);
    }

    function Ship(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 dp,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender,
            "Incorrect"
        );
        test_add_dp(msg.sender,10000000000);
        Ship_his memory s=Ship_his(pkcoix,pkcoiy,hclj,dp);
        Pledger[msg.sender]-=dp;
        Transfer("Store",hclj,dp,msg.sender);
        ship_his_lis.push(s);
    }

    function Distribute(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 cmqx,uint256 cmqy,uint256 tokqx,uint256 tokqy,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender,
            "Incorrect"
        );
        uint256 dt=now;
        Dis_his memory d=Dis_his(pkcoix,pkcoiy,hclj,cmqx,cmqy,tokqx,tokqy,dt);
        Transfer("Distribute",hclj,0,msg.sender);
        dis_his_lis.push(d);
    }

    function Deliver2(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 dp,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender,
            "Incorrect"
        );
        test_add_dp(msg.sender,10000000000);
        Deliver2_his memory d=Deliver2_his(pkcoix,pkcoiy,hclj,dp,dt);
        Pledger[msg.sender]-=dp;
        Transfer("Deliver2",hclj,dp,msg.sender);
        deliver2_his_lis.push(d);
    }

    function Receive(bytes32 hfbi,uint256 dt,bytes32 hcl,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender,
            "Incorrect"
        );
        Receive_his memory r=Receive_his(hfbi,dt,hcl);
        Refund(hcl);
        receive_his_lis.push(r);
    }


    function Transfer(string memory tx,bytes32 hcl,uint256 dp,address from) private{
        if(compareStrings(tx,"Donate")){
            if(DonIsEmpty[hcl]==0){//empty
                 address[] memory aaaas=new address[](3);
                 don memory d=don(State.submitted,0,aaaas);
                 Don[hcl]=d;
                 DonIsEmpty[hcl]=1;
            }
        }

        else if(compareStrings(tx,"Deliver1")){

            //dp
            don memory d=Don[hcl];
            if(d.donstate==State.submitted){
                d.donstate=State.delivery;
                d.deposit=d.deposit+dp;
                d.DepositList[0]=from;

                DepositTable[from]=dp;
                Don[hcl]=d;
            }
        }

        else if(compareStrings(tx,"Store")){
            don memory d=Don[hcl];
            if(d.donstate==State.delivery){
                d.donstate=State.stored;
                d.deposit=d.deposit+dp;
                Don[hcl]=d;
            }
        }

        else if(compareStrings(tx,"Ship")){
            //dp
            don memory d=Don[hcl];
            if(d.donstate==State.stored){
                d.donstate=State.shipped;
                d.deposit=d.deposit+dp;
                d.DepositList[1]=from;
                DepositTable[from]=dp;
                Don[hcl]=d;
            }
        }

        else if(compareStrings(tx,"Deliver2")){
            //dp
            don memory d=Don[hcl];
            if(d.donstate==State.shipped){
                d.donstate=State.delivered;
                d.deposit=d.deposit+dp;
                d.DepositList[2]=from;
                DepositTable[from]=dp;
                Don[hcl]=d;
            }
        }

    }

    function Refund(bytes32 hcl) private{
        don memory d=Don[hcl];
        if(DonIsEmpty[hcl]!=0){// not empty
            if(d.donstate==State.delivered){

                for(uint8 i=0;i<3;i++){
                    address aaaaa=d.DepositList[i];
                    uint256 dp=DepositTable[aaaaa];
                    Pledger[aaaaa]+=dp;
                    DepositTable[aaaaa]-=dp;
                }
                //Refund
                d.donstate=State.received;
                d.deposit=0;
                Don[hcl]=d;
                
            }
        }
    }


    function compareStrings(string memory a, string memory b) public view returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }

}
