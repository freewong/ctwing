package Aep_nb_device_management

import (
	"net/http"

	"github.com/freewong/ctwing/apis/core"
)

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func BatchCreateNBDevice(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_nb_device_management/batchNBDevice"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200828140355"

	application := appKey
	key := appSecret

	return core.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func BatchCancelDevices(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_nb_device_management/cancelledDevices"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20211009093738"

	application := appKey
	key := appSecret

	return core.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func DeleteDeviceByImei(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_nb_device_management/deleteDeviceByImei"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20220226071405"

	application := appKey
	key := appSecret

	return core.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
