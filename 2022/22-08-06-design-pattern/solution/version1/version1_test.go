package version1

import "testing"

func TestLogicObj_IsMarketHit(t *testing.T) {
	obj := &LogicObj{}
	obj.IsMarketHit(&Parameter{
		PageSize: 1,
		PageNo:   1,
		ReqNum:   1,
		Type:     MarketList,
	})
}
