package version2

/*
//一个接口
interface IGroupLabelStrategyService {

    //这个方法对应策略实现类的具体实现
    boolean processBiz(Parameter dto);

    //这个方法就是策略类的类型，也就是对应```if...else```条件判断的类型
    String getType();
}

//企业客群类型的策略实现类
EnterpriseGroupLablelStrategyServiceImpl implements IGroupLabelStrategyService{

    //对应企业客群类型的条件分支里面的实现
    boolean processBiz(Parameter dto){

         //开关关闭不请求
       if(isEnterpriseSwitchClose){
          return false;
       }

        //请求只有一条记录的话
        if(dto.reqNum==1){
            //调用大数据的点查接口
            return singleRemoteEOIinvoke(dto);

            //请求超过一条的话
        }else if(dto.reqNum>1){

            //调用远程大数据批量接口
            return batchRemoteEOIinvoke(dto);
        }

     }

    //对应企业类型
    String getType(){
       return "enterprise";
    }
}

//市场营销类型的策略实现类
MarketListGroupLablelStrategyServiceImpl implements IGroupLabelStrategyService{

     //对应市场营销类型的条件分支里面的实现
     boolean processBiz(Parameter dto){

       //开关关闭不请求
       if(isMarketListSwitchClose){
          return false;
       }

        //请求只有一条记录的话
        if(dto.reqNum==1){
            //调用营销点查接口
            return singleRemoteMarketinvoke(dto);

          //请求超过一条的话
        }else if(dto.reqNum>1){
            //调用营销批量接口
            return batchRemoteMarketinvoke(dto);
        }

      }

      String getType(){
         return "market_list";
       }
}
*/

const (
	EnterPrise = "enterprise"
	MarketList = "market_list"
)

var (
	IsEnterpriseSwitchClose = false
	IsMarketListSwitchClose = false
	LogicMap                = map[string]IGroupLabelStrategyService{}
)

func init() {
	LogicMap[EnterPrise] = new(EnterpriseGroupLabelStrategyServiceImpl)
	LogicMap[MarketList] = new(MarketListGroupLablelStrategyServiceImpl)
}

type Parameter struct {
	PageSize int
	PageNo   int
	ReqNum   int
	Type     string
}

func ProcessBiz(dto *Parameter) bool {
	impl := LogicMap[dto.Type]
	if impl != nil {
		return impl.processBiz(dto)
	}
	return false
}

func SingleRemoteEOIInvoke(dto *Parameter) bool {
	return false
}

func BatchRemoteEOIInvoke(dto *Parameter) bool {
	return false
}

func SingleRemoteMarketInvoke(dto *Parameter) bool {
	return false
}

func BatchRemoteMarketInvoke(dto *Parameter) bool {
	return false
}
