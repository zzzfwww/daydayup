package version2

import "testing"

func TestProcessBiz(t *testing.T) {
	ProcessBiz(&Parameter{
		Type:   EnterPrise,
		ReqNum: 1,
		PageNo: 2,
	})
}
