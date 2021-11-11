pragma solidity 0.5.14;

import "./EllipticCurve.sol";

contract DiSC {

  uint256 public constant GX = 0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798;
  uint256 public constant GY = 0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8;
  uint256 public constant AA = 0;
  uint256 public constant BB = 7;
  uint256 public constant PP = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F;
  //Parameters of Secp256k1
  uint256 public HX=0;
  uint256 public HY=0;
  struct RequestStruct{
      uint256 Ri;//Region
      uint256 ci;//collector id
      bytes32 pkcoi;//public key
      uint256 dni;//donation number
      
  }
  
  struct ResponseStruct{
      uint256 tp1;//id of Transponder1
      uint256 lc;//id of Logistics company
      uint256 idj1;//order ID idj1
      uint256 dnj;//donation number
      string clj;//checking list
      //bytes32 pkcoi;//pk of coi
      uint256 pkcoix;
      uint256 pkcoiy;//pk point
      uint256 rj;//random number
      RequestStruct req;
      uint256 dt;//date time
  }
  
  struct FeedbackStruct{
      uint256 tp2;
      uint256 lci;
      uint256 id;
      uint256 dnj;
      uint256 pkcoix;
      uint256 pkcoiy;
      string clj;
      uint256 dt;
  }
  
  uint256 sum=0;
  ResponseStruct[] Res;
  
  //
  string[] List;
  uint256 ListLength=0;
  //
  //Dis{}={pk,hash,cm,tok,dt}
  uint256[] DisPkX;
  uint256[] DisPkY;
  bytes32[] DisH;
  uint256[] DisCmX;
  uint256[] DisCmY;
  uint256[] DisTokX;
  uint256[] DisTokY;
  uint256[] DisDt;
  //FinalDis{}
  uint256[] FinalDisAggX;
  uint256[] FinalDisAggY;
  uint256[] FinalDisDt;
  bytes32[] FinalDisSig;
  
  
  
  constructor(uint256 random) public{
      //Generate H
      (uint256 hx,uint256 hy)=EllipticCurve.ecMul(
      random,
      GX,
      GY,
      AA,
      PP
    );
    HX=hx;
    HY=hy;
  }
  
  function getNewGenerator(uint256 random) external pure returns (uint256, uint256) {
    return EllipticCurve.ecMul(
      random,
      GX,
      GY,
      AA,
      PP
    );
  }
  
  function Response(uint256 ReqRi,uint256 ReqCi,bytes32 ReqPkcoi,uint256 ReqDni,uint256 tp1,uint256 lc,uint256 idj1,uint256 dnj,string memory clj,uint256 pkcoix,uint256 pkcoiy,uint256 rj,uint256 dt) public{
      //Create response here
      RequestStruct memory req=RequestStruct(ReqRi,ReqCi,ReqPkcoi,ReqDni);
      ResponseStruct memory res=ResponseStruct(tp1,lc,idj1,dnj,clj,pkcoix,pkcoiy,rj,req,dt);
      Res.push(res);
  }
  //H=random*G
  function DistributePartOne() public returns(uint256[] memory,uint256[] memory,bytes32[] memory,uint256[] memory,uint256[] memory,uint256){
      for(uint i=0;i<Res.length;i++){
          ResponseStruct memory res=Res[i];
          (uint256 cm1x,uint256 cm1y)=EllipticCurve.ecMul(res.dnj,GX,GY,AA,PP);//g^dnj
          (uint256 cm2x,uint256 cm2y)=EllipticCurve.ecMul(res.rj,HX,HY,AA,PP);//h^rj
          (uint256 cmx,uint256 cmy)=EllipticCurve.ecAdd(cm1x,cm1y,cm2x,cm2y,AA,PP);//cm=g^dnj * h^rj
          (uint256 tokx,uint256 toky)=EllipticCurve.ecMul(res.rj,res.pkcoix,res.pkcoiy,AA,PP); //tok=pkcoi^rj
          sum+=res.dnj;
          
          DisPkX.push(res.pkcoix);
          DisPkY.push(res.pkcoiy);
          DisH.push(keccak256(abi.encodePacked(res.clj)));//sha256 will throw "invalid opcode 0xfa" exception
          DisCmX.push(cmx);
          DisCmY.push(cmy);
          DisTokX.push(tokx);
          DisTokY.push(toky);
          DisDt.push(res.dt);
          
          List.push(res.clj);//string array return
          ListLength++;
          //Dis{}={pk,hash,cm,tok,dt}
          //pk=(x,y),cm=(x,y),tok=(x,y)
      }
      //Agg
      (uint256 Aggx,uint256 Aggy)=EllipticCurve.ecMul(sum,GX,GY,AA,PP);//g^sum
      FinalDisAggX.push(Aggx);
      FinalDisAggY.push(Aggy);
      //Signature?Inside or Outsid Smart Contract?
      //Signature Not inplemented
      return (DisPkX,DisPkY,DisH,DisCmX,DisCmY,ListLength);
      //InternalCompilerError: Stack too deep,try using fewer variables, So we have to return values throgh two functions
  }
  //To return string array List,we have to use "pragma experimental ABIEncoderV2", but EVM doesn't support 
  function DistributePartTwo() public view returns(uint256[] memory,uint256[] memory,uint256[] memory){
      return(DisTokX,DisTokY,DisDt);
  }
  
  function getFinalDistribute() public view returns(uint256[] memory,uint256[] memory,bytes32[] memory,uint256[] memory){
      return(FinalDisAggX,FinalDisAggY,FinalDisSig,FinalDisDt);
  }
  
  //To get the data of List, we mark the length of List in DiSC, see DistributePartOne
  //And we use function getList and change the index to access List
  function getList(uint256 index) public view returns(string memory){
      return List[index];
  }
  
  function Clear() public{
      delete DisPkX;
      delete DisPkY;
      delete DisH;
      delete DisCmX;
      delete DisCmY;
      delete DisTokX;
      delete DisTokY;
      delete DisDt;
      delete List;
      delete FinalDisDt;
      delete FinalDisSig;
      delete FinalDisAggY;
      delete FinalDisAggX;
      
  }
  
  function Refund(){
      
  }
  
}
