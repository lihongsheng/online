package api

import "online/server/web/gen/restapi/operations"

// 重写 swagger 生成的handle ，提供真正的服务
func (s *APIManager) initAPIBase(base *operations.OnlineAPI) {
	// 真正提供服务的地方
	base.GetAuthListHandler = s.GetUserAuthList()
}
