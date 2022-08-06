package version2

type EnterpriseGroupLabelStrategyServiceImpl struct{}

func (e *EnterpriseGroupLabelStrategyServiceImpl) processBiz(dto *Parameter) bool {
	//开关关闭不请求
	if IsEnterpriseSwitchClose {
		return false
	}

	//请求只有一条记录的话
	if dto.ReqNum == 1 {
		//调用大数据的点查接口
		return SingleRemoteEOIInvoke(dto)

		//请求超过一条的话
	} else if dto.ReqNum > 1 {

		//调用远程大数据批量接口
		return BatchRemoteEOIInvoke(dto)
	}
	return false
}

//对应企业类型
func (e *EnterpriseGroupLabelStrategyServiceImpl) getType() string {
	return EnterPrise
}

type MarketListGroupLablelStrategyServiceImpl struct{}

func (e *MarketListGroupLablelStrategyServiceImpl) processBiz(dto *Parameter) bool {
	//开关关闭不请求
	if IsMarketListSwitchClose {
		return false
	}

	//请求只有一条记录的话
	if dto.ReqNum == 1 {
		//调用大数据的点查接口
		return SingleRemoteMarketInvoke(dto)

		//请求超过一条的话
	} else if dto.ReqNum > 1 {

		//调用远程大数据批量接口
		return BatchRemoteMarketInvoke(dto)
	}
	return false
}

//对应企业类型
func (e *MarketListGroupLablelStrategyServiceImpl) getType() string {
	return MarketList
}
