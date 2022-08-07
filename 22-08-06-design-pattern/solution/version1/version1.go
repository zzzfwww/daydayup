package version1

/*
class Parameter{
    int pageSize;
    int pageNo;
    int reqNum；
    //其他参数。
}

//逻辑处理，是否命中客群
boolean isMarketHit(Parameter dto){
    //如果是企业客群类型
    if(dto.type == 'enterprise'){

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

            //调用大数据的批量接口
            return batchRemoteEOIinvoke(dto);
        }

        //如果是市场营销类型
    }else if(dto.type=='market_list'){

      //开关关闭不请求
       if(isMarketListSwitchClose){
          return false;
       }
        //请求只有一条记录的话
        if(dto.reqNum==1){
            //调用营销的点查接口
            return singleRemoteMarketinvoke(dto);

          //请求超过一条的话
        }else if(dto.reqNum>1){
            //调用营销的批量接口
            return batchRemoteMarketinvoke(dto);
        }
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
)

type Parameter struct {
	PageSize int
	PageNo   int
	ReqNum   int
	Type     string
}

type LogicObj struct{}

func (l *LogicObj) IsMarketHit(dto *Parameter) bool {
	if dto.Type == EnterPrise {
		// 开关关闭不请求
		if IsEnterpriseSwitchClose {
			return false
		}

		// 请求只有一条记录的话
		if dto.ReqNum == 1 {
			// 调用大数据的点查接口
			return SingleRemoteEOIInvoke(dto)

			// 请求超过一条的话
		} else if dto.ReqNum > 1 {
			// 调用大数据的批量接口
			return BatchRemoteEOIInvoke(dto)
		}
	} else if dto.Type == MarketList {
		// 如果是市场营销类型

		// 开关关闭不请求
		if IsMarketListSwitchClose {
			return false
		}
		// 请求只有一条记录的话
		if dto.ReqNum == 1 {
			// 调用营销的点查接口
			return SingleRemoteMarketInvoke(dto)

			// 请求超过一条的话
		} else if dto.ReqNum > 1 {
			// 调用营销的批量接口
			return BatchRemoteMarketInvoke(dto)
		}
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
