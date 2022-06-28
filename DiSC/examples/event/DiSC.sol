pragma solidity 0.4.21;


contract DiSC {

  uint256 public constant GX = 0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798;
  uint256 public constant GY = 0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8;
  uint256 public constant AA = 0;
  uint256 public constant BB = 7;
  uint256 public constant PP = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F;
  //Parameters of Secp256k1
  uint256 public HX=0;
  uint256 public HY=0;
  uint256 public bigr1=0;
  uint256 public bigr2=0;
  uint256 public priKey=0;
  uint256 public pubKeyX=0;
  uint256 public pubKeyY=0;
  byte [32*6] public tmp;


    uint256 constant public gx = 0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798;
    uint256 constant public gy = 0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8;
    uint256 constant public n = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F;
    uint256 constant public a = 0;
    uint256 constant public b = 7;

    /*
    function swap(uint256* a, uint256* b)public returns(){
        uint256 c = *a;
        *a = *b;
        *b = c;
    }*/

    function gcb(uint256 _a, uint256 _b)public pure returns(uint256){
        uint256 t;
        uint256 _c;
        if(a<b){
            t = _a;
            _a = _b;
            _b = t;
        }
        _c = _a%_b;
        while(_c!=0){
            //t = _a;
            _a = _b;
            _b = _c;
            _c = _a%_b;
        }
        return _b;
    }

    uint256[] xishu;
    function modInverse(uint256 _a)public returns(uint256){
        if(_a==1){
            return 1;
        }
        //必须要互素才有逆元
        if(PP%_a==0){
            return 0;//错误
        }
        
        uint256 divisend =  PP;
        uint256 divisor = _a;
        
        uint256 ind=0;
        uint256 c = divisend%divisor;
        while(c!=0){
            xishu.push(divisend/divisor);//整除
            ind = ind + 1;
            divisend = divisor;
            divisor = c;
            c = divisend%divisor;
        }
        uint256 res1=1;
        uint256 res2=PP-xishu[ind];//加法逆元
        ind--;
        while(ind!=0){
            c = res1;
            res1 = res2;
            res2 = c + res1*(PP-xishu[ind]);
            res2 = res2%PP;
            ind--;
        }
        delete xishu;
        return res2%PP;
    }

    function addPoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2) public returns(uint256, uint256){
        if(0==x1&&0==y2){
            return (x2,y2);
        }
        if(0==x2&&0==y2){
            return (x1,y1);
        }
        uint256 k;
        uint256 t;
        if(x1==x2){
            t = mulmod(x1,x1,PP);
            t = mulmod(t,3,PP);
            k = mulmod(2,y1,PP);
            k = modInverse(k);
            k = mulmod(k,t,PP);
        }else{//都是无符号整数，如果小怎么办。用逆元
            if(y2>y1){
                t=y2-y1;
                if(x2>x1){
                    k=x2-x1;
                    k = modInverse(k);
                    k = mulmod(t,k,PP);
                }else{
                    k=x1-x2;
                    k = PP-k;//加法逆元
                    k = modInverse(k);
                    k = mulmod(t,k,PP);
                    
                }
            }else{
                t=y1-y2;
                t=PP-t;//加法逆元
                if(x2>x1){
                    k=x2-x1;
                    k = modInverse(k);
                    k = mulmod(t,k,PP);
                }else{
                    k=x1-x2;
                    k = PP-k;//加法逆元
                    k = modInverse(k);
                    k = mulmod(t,k,PP);
                    
                }
            }
        }

        uint256 x3 = mulmod(k,k,PP);
        t = addmod(x1,x2,PP);
        t = PP-t;//加法逆元
        x3 = x3+t;
        uint256 y3;
        //x3 = x3 -x1 -x2;
        if(x1>x3){
            y3 = mulmod(k,x1-x3,PP);
        }else{
            y3 = mulmod(k,PP-x3+x1,PP);
        }
        if(y3>y1){
            y3=y3-y1;
        }else{
            y3=y1-y3;
            y3=PP-y3;
        }
        return (x3,y3);
    }

    function kPoint(uint256 x, uint256 y, uint256 k)public returns(uint256, uint256){
        uint256 resx=0;
        uint256 resy=0;
        //uint256 tmpx=x;
        //uint256 tmpy=y;
        while(k!=0){
            if(k%2==1){
                (resx,resy) = addPoint(resx,resy,x,y);
            }
            (x,y) = addPoint(x,y,x,y);
            k = k/2;
        }
        return (resx,resy);
    }




    function _jAdd(
        uint256 x1, uint256 z1,
        uint256 x2, uint256 z2)
        public 
        pure
        returns(uint256 x3, uint256 z3)
    {
        (x3, z3) = (
            addmod(
                mulmod(z2, x1, PP),
                mulmod(x2, z1, PP),
                n
            ),
            mulmod(z1, z2, PP)
        );
    }

    function _jSub(
        uint256 x1, uint256 z1,
        uint256 x2, uint256 z2)
        public 
        pure
        returns(uint256 x3, uint256 z3)
    {
        (x3, z3) = (
            addmod(
                mulmod(z2, x1, PP),
                mulmod(PP - x2, z1, PP),
                PP
            ),
            mulmod(z1, z2, PP)
        );
    }

    function _jMul(
        uint256 x1, uint256 z1,
        uint256 x2, uint256 z2)
        public 
        pure
        returns(uint256 x3, uint256 z3)
    {
        (x3, z3) = (
            mulmod(x1, x2, PP),
            mulmod(z1, z2, PP)
        );
    }

    function _jDiv(
        uint256 x1, uint256 z1,
        uint256 x2, uint256 z2) 
        public 
        pure
        returns(uint256 x3, uint256 z3)
    {
        (x3, z3) = (
            mulmod(x1, z2, PP),
            mulmod(z1, x2, PP)
        );
    }

    function _inverse(uint256 val) public pure
        returns(uint256 invVal)
    {
        uint256 t = 0;
        uint256 newT = 1;
        uint256 r = PP;
        uint256 newR = val;
        uint256 q;
        while (newR != 0) {
            q = r / newR;

            (t, newT) = (newT, addmod(t, (PP - mulmod(q, newT, PP)), PP));
            (r, newR) = (newR, r - q * newR );
        }

        return t;
    }

    function _ecAdd(
        uint256 x1, uint256 y1, uint256 z1,
        uint256 x2, uint256 y2, uint256 z2) 
        public 
        pure
        returns(uint256 x3, uint256 y3, uint256 z3)
    {
        uint256 lx;
        uint256 lz;
        uint256 da;
        uint256 db;

        if (x1 == 0 && y1 == 0) {
            return (x2, y2, z2);
        }

        if (x2 == 0 && y2 == 0) {
            return (x1, y1, z1);
        }

        if (x1 == x2 && y1 == y2) {
            (lx, lz) = _jMul(x1, z1, x1, z1);
            (lx, lz) = _jMul(lx, lz, 3, 1);
            (lx, lz) = _jAdd(lx, lz, a, 1);

            (da,db) = _jMul(y1, z1, 2, 1);
        } else {
            (lx, lz) = _jSub(y2, z2, y1, z1);
            (da, db) = _jSub(x2, z2, x1, z1);
        }

        (lx, lz) = _jDiv(lx, lz, da, db);

        (x3, da) = _jMul(lx, lz, lx, lz);
        (x3, da) = _jSub(x3, da, x1, z1);
        (x3, da) = _jSub(x3, da, x2, z2);

        (y3, db) = _jSub(x1, z1, x3, da);
        (y3, db) = _jMul(y3, db, lx, lz);
        (y3, db) = _jSub(y3, db, y1, z1);

        if (da != db) {
            x3 = mulmod(x3, db, PP);
            y3 = mulmod(y3, da, PP);
            z3 = mulmod(da, db, PP);
        } else {
            z3 = da;
        }
    }

    function _ecDouble(uint256 x1, uint256 y1, uint256 z1) public pure
        returns(uint256 x3, uint256 y3, uint256 z3)
    {
        (x3, y3, z3) = _ecAdd(x1, y1, z1, x1, y1, z1);
    }

    function _ecMul(uint256 d, uint256 x1, uint256 y1, uint256 z1) public pure
        returns(uint256 x3, uint256 y3, uint256 z3)
    {
        uint256 remaining = d;
        uint256 px = x1;
        uint256 py = y1;
        uint256 pz = z1;
        uint256 acx = 0;
        uint256 acy = 0;
        uint256 acz = 1;

        if (d == 0) {
            return (0, 0, 1);
        }

        while (remaining != 0) {
            if ((remaining & 1) != 0) {
                (acx,acy,acz) = _ecAdd(acx, acy, acz, px, py, pz);
            }
            remaining = remaining / 2;
            (px, py, pz) = _ecDouble(px, py, pz);
        }

        (x3, y3, z3) = (acx, acy, acz);
    }

    function ecadd(
        uint256 x1, uint256 y1,
        uint256 x2, uint256 y2)
        public
        pure
        returns(uint256 x3, uint256 y3)
    {
        uint256 z;
        (x3, y3, z) = _ecAdd(x1, y1, 1, x2, y2, 1);
        z = _inverse(z);
        x3 = mulmod(x3, z, PP);
        y3 = mulmod(y3, z, PP);
    }

    function ecmul(uint256 x1, uint256 y1, uint256 scalar) public pure
        returns(uint256 x2, uint256 y2)
    {
        uint256 z;
        (x2, y2, z) = _ecMul(scalar, x1, y1, 1);
        z = _inverse(z);
        x2 = mulmod(x2, z, PP);
        y2 = mulmod(y2, z, PP);
    }


    // function point_hash( uint256[2] point )
    //     public pure returns(address)
    // {
    //     return address(uint256(keccak256(abi.encodePacked(point[0], point[1]))) & 0x00FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF);
    // }

    /**
    * hash(g^a + B^c)
    */
    function sbmul_add_mul(uint256 s, uint256[2] B, uint256 c)
        public pure returns(address)
    {
        uint256 Q = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141;
        s = (Q - s) % Q;
        s = mulmod(s, B[0], Q);

        return ecrecover(bytes32(s), B[1] % 2 != 0 ? 28 : 27, bytes32(B[0]), bytes32(mulmod(c, B[0], Q)));
    }

    //
    // Based on the original idea of Vitalik Buterin:
    // https://ethresear.ch/t/you-can-kinda-abuse-ecrecover-to-do-ecmul-in-secp256k1-today/2384/9
    //
    // function ecmulVerify(uint256 x1, uint256 y1, uint256 scalar, uint256 qx, uint256 qy) public pure
    //     returns(bool)
    // {
    //     address signer = sbmul_add_mul(0, [x1, y1], scalar);
    //     return point_hash([qx, qy]) == signer;
    // }

    function publicKey(uint256 privKey) public pure
        returns(uint256 qx, uint256 qy)
    {
        return ecmul(GX, GY, privKey);
    }

    // function publicKeyVerify(uint256 privKey, uint256 x, uint256 y) public pure
    //     returns(bool)
    // {
    //     return ecmulVerify(gx, gy, privKey, x, y);
    // }

    function deriveKey(uint256 privKey, uint256 pubX, uint256 pubY) public pure
        returns(uint256 qx, uint256 qy)
    {
        uint256 z;
        (qx, qy, z) = _ecMul(privKey, pubX, pubY, 1);
        z = _inverse(z);
        qx = mulmod(qx, z, PP);
        qy = mulmod(qy, z, PP);
    }


//注释是因为论文有提到，但实现的时候可以省略
struct RequestStruct{
      uint256 Ri;//Region
      uint256 ci;//collector id
      //uint256 pubKeyX;
      //uint256 pubKeyY;
      uint256 dni;//donation number
  }
  
  struct Proof {
	uint256 deltaX;
    uint256 deltaY;
    uint256 yX;    
    uint256 yY;     
    uint256 gX;    
    uint256 gY;     
    uint256 aa;    
}

  struct ResponseStruct{
      uint256 tp1;//id of Transponder1
      uint256 lc;//id of Logistics company
      uint256 idj1;//order ID idj1
      uint256 dnj;//donation number
      bytes32 clj;//checking list
      //bytes32 pkcoi;//pk of coi//no use
      // cm and tok
      uint256 comx;
      uint256 comy;
      uint256 tokx;
      uint256 toky;
      uint256 rj;//random number
      //RequestStruct req;
      uint256 dt;//date time
  }
  
  struct FeedbackStruct{
      uint256 tp2;
      uint256 lci;
      uint256 id;
      uint256 dnj;
      //uint256 pkcoix;
      //uint256 pkcoiy;
      uint256 clj;
      uint256 dt;
  }

//14*32
struct parameters {
    uint256 ReqRi;
	uint256 ReqCi;
	//uint256 ReqPkcoiX;
    //uint256 ReqPkcoiY;
	uint256 ReqDni;
	uint256 tp1;    
	uint256 lc;    
	uint256 idj1;
	uint256 dnj;  
	bytes32 clj;
    //这是捐赠者生成的
	uint256 comx;//pkx
    uint256 comy;//pky
	uint256 tokx;//
    uint256 toky;// 
	uint256 rj;  
	uint256 dt;   

    
}

  
  uint256 sum=0;
  uint256 ATokoutX=0;
  uint256 ATokoutY=0;
  uint256 AComoutX=0;
  uint256 AComoutY=0;
  uint256 ATokinX=0;
  uint256 ATokinY=0;
  uint256 ACominX=0;
  uint256 ACominY=0;
  ResponseStruct[] Res;
  
  //
  bytes32[] List;
  uint256 ListLength=0;
  //
  //Dis{}={pk,hash,cm,tok,dt}
  //uint256[] DisPkX;
  //uint256[] DisPkY;
  bytes32[] DisH;
  uint256[] DisCmX;
  uint256[] DisCmY;
  uint256[] DisTokX;
  uint256[] DisTokY;
  uint256[] DisDt;
  uint256[9] Auddition1;
  uint256[9] Auddition2;
  uint256 Aggx;
  uint256 Aggy;
/*
  uint256 DeltaX1;
  uint256 DeltaY1;
  uint256 H_R_SK_X1;
  uint256 H_R_SK_Y1;
  uint256 H_R_X1;
  uint256 H_R_Y1;
  uint256 Aa1;

  uint256 DeltaX2;
  uint256 DeltaY2;
  uint256 H_R_SK_X2;
  uint256 H_R_SK_Y2;
  uint256 H_R_X2;
  uint256 H_R_Y2;
  uint256 Aa2;
*/
  //FinalDis{}
  uint256[] FinalDisAggX;
  uint256[] FinalDisAggY;
  uint256[] FinalDisDt;
  bytes32[] FinalDisSig;
  
  


    //不确定是否会自动调用
  function DiSC() public{
      //Generate H
      //uint256 hx;
      //uint256 hy;
      //(hx, hy)=ecmul(
      //GX,
      //GY,
      //random
    //);
    //HX=hx;
    //HY=hy;
  }

  function AddOne(uint256 num)public returns (uint256){
      uint256 res = num;
      res = res + 1;
      bigr1 = res;
      return res;
  }
  
  function getNewGenerator(uint256 hx,
                        uint256 hy,
                        uint256 r1, 
                        uint256 r2, 
                        uint256 prik,
                        uint256 pubkX,
                        uint256 pubkY) public returns (uint256, uint256) {
      HX = hx;
      HY = hy;
      bigr1 = r1;
      bigr2 = r2;
      priKey = prik;
      pubKeyX = pubkX;
      pubKeyY = pubkY;
    return (HX,HY);
  }
  


    function intToStr(uint256 _i) internal pure returns (string memory) {
        if (_i == 0) return "0";

        uint256 j = _i;
        uint256 len;
        while (j > 0) {
            j /= 10;
            len++;
        }

        bytes memory bstr = new bytes(len);
        uint256 k = len - 1;
        while (_i > 0) {
            bstr[k] = byte(uint8( 48 + _i % 10));
            _i /= 10;
            k--;
        }
        return string(bstr);
    }


    //参数太多了，超出16个，前三个没用，就不传了
  function Response(//uint256 arr0,
                    //uint256 arr1,
                    //uint256 arr2,
                    uint256 arr3,
                    uint256 arr4,
                    uint256 arr5,
                    uint256 arr6,
                    uint256 arr7,
                    uint256 arr8,
                    uint256 arr9,
                    uint256 arr10,
                    uint256 arr11,
                    uint256 arr12,
                    uint256 arr13)  public {//returns(uint256[] memory, uint256[] memory) {
      //Create response here
    //RequestStruct memory req=RequestStruct(arr0,arr1,arr2);
    ResponseStruct memory res=ResponseStruct(arr3,arr4,arr5,arr6,bytes32(arr7),arr8,arr9,arr10,arr11,arr12,/*req,*/ arr13);
    Res.push(res);
    bytes32 h = keccak256(res.clj);
    DisH.push(h);//sha256 will throw "invalid opcode 0xfa" exception//太老了，没有这个指令
    //return h,res.dt;
  }





  //H=random*G
  function DistributePartOne() 
  public returns(bytes32[] memory,
                uint256[] memory,
                uint256[] memory,
                uint256[] memory,
                uint256[] memory,
                uint256[] memory){

          uint256 cm1x;
          uint256 cm1y;
          uint256 cm2x;
          uint256 cm2y;
          uint256 comx;
          uint256 comy;
          uint256 tokx;
          uint256 toky;
      for(uint i=0;i<Res.length;i++){
          ResponseStruct memory res=Res[i];
          
          ( cm1x, cm1y)=ecmul(GX,GY,res.dnj);//g^dnj
          ( cm2x, cm2y)=ecmul(HX,HY,res.rj);//h^rj
          ( comx, comy)=ecadd(cm1x,cm1y,cm2x,cm2y);//cm=g^dnj * h^rj
          ( tokx, toky)=ecmul(pubKeyX,pubKeyY,res.rj); //tok=pkcoi^rj
          sum+=res.dnj;
          //ATokout = sum(tok)
          if(ATokoutX==0 && ATokoutY==0){
              ATokoutX=tokx;
              ATokoutY=toky;
          }else{
             (ATokoutX,ATokoutY)=ecadd(ATokoutX,ATokoutY,tokx,toky);
          }
          if(AComoutX==0 && AComoutY==0){
              AComoutX=comx;
              AComoutY=comy;
          }else{
             (AComoutX,AComoutY)=ecadd(AComoutX,AComoutY,comx,comy);
          }

          if(ATokinX==0 && ATokinY==0){
              ATokinX=res.tokx;
              ATokinY=res.toky;
          }else{
             (ATokinX,ATokinY)=ecadd(ATokinX,ATokinY,res.tokx,res.toky);
          }
          if(ACominX==0 && ACominY==0){
              ACominX=res.comx;
              ACominY=res.comy;
          }else{
             (ACominX,ACominY)=ecadd(ACominX,ACominY,res.comx,res.comy);
          }
        


          
          
          //DisPkX.push(res.pkcoix);
          //DisPkY.push(res.pkcoiy);
          //DisH.push(keccak256(res.clj));//sha256 will throw "invalid opcode 0xfa" exception//太老了，没有这个指令
          DisCmX.push(comx);
          DisCmY.push(comy);
          DisTokX.push(tokx);
          DisTokY.push(toky);
          DisDt.push(res.dt);
          //genProof(res.rj);
          List.push(res.clj);//string array return
          ListLength++;
          //Dis{}={pk,hash,cm,tok,dt}+sigma
          //pk=(x,y),cm=(x,y),tok=(x,y)
          //Son{}={Agg,deltaX,deltaY,yX,yY,gX,gY,aa,dt,sigma}
      }
      //Agg
      

      (Aggx, Aggy)=ecmul(GX,GY,sum);//g^sum
      //FinalDisAggX.push(Aggx);
      //FinalDisAggY.push(Aggy);
      //genProof twice
      //uint256 tx;
      //uint256 ty;//变量超了
      cm2x = PP-Aggx;
      (cm1x,cm1y) = ecadd(AComoutX,AComoutY,cm2x,Aggy);
      genProof(cm1x,cm1y,true);
      Auddition1[0]=Aggx;
      Auddition1[1]=Aggy;
      Auddition2[0]=Aggx;
      Auddition2[1]=Aggy;
      (cm1x,cm1y) = ecadd(ACominX,ACominY,cm2x,Aggy);
      genProof(cm1x,cm1y,false);
      //Signature?Inside or Outsid Smart Contract?
      //Signature Not inplemented
      //return (DisH,DisCmX,DisCmY,DisTokX,DisTokY,DisDt);
      return (DisH,DisCmX,DisCmY,DisTokX,DisTokY,DisDt);
      //InternalCompilerError: Stack too deep,try using fewer variables, So we have to return values throgh two functions
      //如果局部变量太多，则函数栈帧太大了，以太坊虚拟机限制运行栈
      //如果改成全局变量
  }

/*
  function DistributePartTwo() 
  public returns(uint256[] memory,
                uint256[] memory,
                uint256[] memory){
        return (DisTokX,DisTokY,DisDt);
    }
    */

/*
  function AudditionPartOne()
  public view returns(uint256 memory,
                    uint256 memory,
                    uint256 memory, 
                    uint256 memory,
                    uint256 memory,
                    uint256 memory, 
                    uint256 memory){
      return(DeltaX1,DeltaY1,H_R_SK_X1,H_R_SK_Y1,H_R_X1,H_R_Y1,Aa1);
  }*/


  function AudditionPartOne()public view returns(uint256[9] memory){
      return Auddition1;
  }

  function AudditionPartTwo()public view returns(uint256[9] memory){
      return Auddition2;
  }

/*
  //To return string array List,we have to use "pragma experimental ABIEncoderV2", but EVM doesn't support 
  function AudditionPartTwo()
  public view returns(uint256 memory,
                    uint256 memory,
                    uint256 memory, 
                    uint256 memory,
                    uint256 memory,
                    uint256 memory, 
                    uint256 memory){
      return(DeltaX2,DeltaY2,H_R_SK_X2,H_R_SK_Y2,H_R_X2,H_R_Y2,Aa2);
  }*/
  
  //function DistributePartFour() public view returns(uint256[] memory,uint256[] memory,bytes32[] memory,uint256[] memory){
      //return(H_R_SK_X,H_R_SK_Y,H_R_X,H_R_Y);
  //}
  
  //To get the data of List, we mark the length of List in DiSC, see DistributePartOne
  //And we use function getList and change the index to access List
  //function getList(uint256 index) public view returns(uint256 memory){
      //return uint256(List[index]);
  //}
  


 function asciiToInteger(bytes32 x) public pure returns (uint256) {
    uint256 y;
    for (uint256 i = 0; i < 32; i++) {
        uint256 c = (uint256(x) >> (i * 8)) & 0xff;
        if (48 <= c && c <= 57)
            y += (c - 48) * 10 ** i;
        else
            break;
    }
    return y;
}





function genProof(uint256 AComX, uint256 AComY,bool bl)public{
    uint256 alpha = 1;
	uint256 deltaX;
    uint256 deltaY;
    uint256 yX;
    uint256 yY;
    (deltaX, deltaY)  =  ecmul(AComX, AComY, alpha);
    (yX, yY) = ecmul(AComX,AComY,priKey);
    //e = hash(ACom,y,delta)

    bytes32 tmpArr = bytes32(AComX);
    for(uint i=0;i<32;i++){ 
        tmp[i] = tmpArr[i];
    }
    tmpArr = bytes32(AComY);
    for(i=32;i<64;i++){ 
        tmp[i] = tmpArr[i-32];
    }
    tmpArr = bytes32(yX);
    for(i=64;i<96;i++){ 
        tmp[i] = tmpArr[i-64];
    }
    tmpArr = bytes32(yY);
    for(i=96;i<128;i++){ 
        tmp[i] = tmpArr[i-96];
    }
    tmpArr = bytes32(deltaX);
    for(i=128;i<160;i++){ 
        tmp[i] = tmpArr[i-128];
    }
    tmpArr = bytes32(deltaY);
    for(i=160;i<192;i++){ 
        tmp[i] = tmpArr[i-160];
    }
    bytes32 _hash = keccak256(tmp);
    uint256 bign = asciiToInteger(_hash);
    uint256 aa = priKey * bign;
    aa = (aa  +  alpha) % PP;
    //uint256 [7] memory uintarr;
    //证明两次，这些都要准备两份
    if(bl){
        Auddition1[2]=deltaX;
        Auddition1[3]=deltaY;//
        Auddition1[4]=AComX;
        Auddition1[5]=AComY;
        Auddition1[6]=yX;
        Auddition1[7]=yY;
        Auddition1[8]=aa;
    }else{
        Auddition2[2]=deltaX;
        Auddition2[3]=deltaY;//
        Auddition2[4]=AComX;
        Auddition2[5]=AComY;
        Auddition2[6]=yX;
        Auddition2[7]=yY;
        Auddition2[8]=aa;
    }
    
}



  function testAdd(uint256 x2,uint256 y2) public returns(uint256 , uint256 ){
       //(x2,y2) =  kPoint(GX,GY,3);
       (x2,y2) = addPoint(x2,y2,HX,HY);
        return(x2,y2);
  }

  function testMul(uint256 x2,uint256 y2) public returns(uint256 ,uint256 ){
      (x2,y2) = kPoint(x2,y2,101);
      return(x2,y2);
  }

  function Clear() public{
      //delete DisPkX;
      //delete DisPkY;
      delete DisH;
      delete DisCmX;
      delete DisCmY;
      delete DisTokX;
      delete DisTokY;
      delete DisDt;
      delete List;
      delete FinalDisDt;
      delete FinalDisSig;//
      delete FinalDisAggY;
      delete FinalDisAggX;
      /*
        delete DeltaX1;
        delete DeltaY1;
        delete H_R_SK_X1;
        delete H_R_SK_Y1;
        delete H_R_X1;
        delete H_R_Y1;
        delete Aa1;
        delete DeltaX2;
        delete DeltaY2;
        delete H_R_SK_X2;
        delete H_R_SK_Y2;
        delete H_R_X2;
        delete H_R_Y2;
        delete Aa2;
    */
      ListLength=0;
  }
  
}