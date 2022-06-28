pragma solidity 0.4.20;
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

    struct Refund_his{
        uint256 pkcoix;
        uint256 pkcoiy;
        uint256 dt;
        bytes32 hclj;
    }
    Refund_his[] refund_his_lis;

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

    function printLenOfStore(uint aaa) public view returns(uint,uint){
        uint bbbb = aaa;
        return(store_his_lis.length,bbbb);
    }


    function test_add_dp(address a,uint256 k) public{
        Pledger[a]=k;
    }

    function compareStrings(string memory a, string memory b) public pure returns (bool) {
        return (keccak256(a) == keccak256(b));
    }

    function Donate(uint256 comx,uint256 comy,uint256 tokx,uint256 toky,bytes32 hres,bytes32 hcl ,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        Donate_his memory d=Donate_his(comx,comy,tokx,toky,hres,hcl,dt);
        Transfer("Donate",hcl,0,msg.sender);
        donate_his_lis.push(d);
    }


    function Deliver1(uint256 pkcoix,uint256 pkcoiy, bytes32 hclj, uint256 dp,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        test_add_dp(msg.sender,10000000000);
        Deliver1_his memory d=Deliver1_his(pkcoix,pkcoiy,hclj,dp,dt);
        Pledger[msg.sender]-=dp;
        Transfer("Deliver1",hclj,dp,msg.sender);
        deliver1_his_lis.push(d);
    }

    function Store(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        Store_his memory ss =Store_his(pkcoix,pkcoiy,hclj,dt);
        Transfer("Store",hclj,0,msg.sender);
        store_his_lis.push(ss);
    }

    function Ship(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 dp,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        test_add_dp(msg.sender,10000000000);
        Ship_his memory ss=Ship_his(pkcoix,pkcoiy,hclj,dp);
        Pledger[msg.sender]-=dp;
        Transfer("Store",hclj,dp,msg.sender);
        ship_his_lis.push(ss);
    }

    function Distribute(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 cmqx,uint256 cmqy,uint256 tokqx,uint256 tokqy,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        uint256 dt=now;
        Dis_his memory d=Dis_his(pkcoix,pkcoiy,hclj,cmqx,cmqy,tokqx,tokqy,dt);
        Transfer("Distribute",hclj,0,msg.sender);
        dis_his_lis.push(d);
    }

    function Deliver2(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 dp,uint256 dt,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        test_add_dp(msg.sender,10000000000);
        Deliver2_his memory d=Deliver2_his(pkcoix,pkcoiy,hclj,dp,dt);
        Pledger[msg.sender]-=dp;
        Transfer("Deliver2",hclj,dp,msg.sender);
        deliver2_his_lis.push(d);
    }

    function Receive(bytes32 hfbi,uint256 dt,bytes32 hcl,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        Receive_his memory rr=Receive_his(hfbi,dt,hcl);
        //Refund(hcl);
        receive_his_lis.push(rr);
    }


    function Transfer(string memory txx,bytes32 hcl,uint256 dp,address from) public{
        if(compareStrings(txx,"Donate")){
            if(DonIsEmpty[hcl]==0){//empty
                 address[] memory aaaas=new address[](3);
                 don memory d=don(State.submitted,0,aaaas);
                 Don[hcl]=d;
                 DonIsEmpty[hcl]=1;
            }
        }

        else if(compareStrings(txx,"Deliver1")){

            //dp
            don memory dd=Don[hcl];
            if(dd.donstate==State.submitted){
                dd.donstate=State.delivery;
                dd.deposit=dd.deposit+dp;
                dd.DepositList[0]=from;

                DepositTable[from]=dp;
                Don[hcl]=dd;
            }
        }

        else if(compareStrings(txx,"Store")){
            don memory d2=Don[hcl];
            if(d2.donstate==State.delivery){
                d2.donstate=State.stored;
                d2.deposit=d2.deposit+dp;
                Don[hcl]=d2;
            }
        }

        else if(compareStrings(txx,"Ship")){
            //dp
            don memory d3=Don[hcl];
            if(d3.donstate==State.stored){
                d3.donstate=State.shipped;
                d3.deposit=d3.deposit+dp;
                d3.DepositList[1]=from;
                DepositTable[from]=dp;
                Don[hcl]=d3;
            }
        }

        else if(compareStrings(txx,"Deliver2")){
            //dp
            don memory d4=Don[hcl];
            if(d4.donstate==State.shipped){
                d4.donstate=State.delivered;
                d4.deposit=d4.deposit+dp;
                d4.DepositList[2]=from;
                DepositTable[from]=dp;
                Don[hcl]=d4;
            }
        }

    }

    function Refund(uint256 pkcoix,uint256 pkcoiy,uint256 dt, bytes32 hcl,bytes32 msgh,uint8 v,bytes32 r,bytes32 s) public{
        require(
            ecrecover(msgh, v, r, s)==msg.sender
        );
        don memory d=Don[hcl];
        Refund_his memory rr=Refund_his(pkcoix,pkcoiy,dt,hcl);
        refund_his_lis.push(rr);
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



}