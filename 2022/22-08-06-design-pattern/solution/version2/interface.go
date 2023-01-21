package version2

type IGroupLabelStrategyService interface {
	processBiz(dto *Parameter) bool
	getType() string
}
