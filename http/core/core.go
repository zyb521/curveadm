/*
*  Copyright (c) 2023 NetEase Inc.
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
 */

/*
* Project: Curveadm
* Created Date: 2023-04-04
* Author: wanghai (SeanHai)
 */

package core

import (
	"fmt"
	"strconv"

	"github.com/opencurve/curveadm/internal/errno"
	"github.com/opencurve/pigeon"
)

const (
	// http method
	HTTP_POST = "POST"
	HTTP_GET  = "GET"

	// method
	METHOD = "method"
)

func Exit(r *pigeon.Request, err error) bool {
	var status int
	if err == nil {
		status = 200
		r.SendJSON(pigeon.JSON{
			"errorCode": "0",
			"errorMsg":  "success",
		})
	} else {
		code := err.(*errno.ErrorCode)
		if code.IsHttpErr() {
			status = code.HttpCode()
		} else {
			status = 503
		}
		r.SendJSON(pigeon.JSON{
			"errorCode": strconv.Itoa(code.GetCode()),
			"errorMsg":  fmt.Sprintf("desc: %s; clue: %s", code.GetDescription(), code.GetClue()),
		})
	}
	return r.Exit(status)
}

func Default(r *pigeon.Request) bool {
	r.Logger().Warn("unupport request uri", pigeon.Field("uri", r.Uri))
	return Exit(r, errno.ERR_UNSUPPORT_REQUEST_URI)
}

func ExitSuccessWithData(r *pigeon.Request, data interface{}) bool {
	r.SendJSON(pigeon.JSON{
		"data":      data,
		"errorCode": "0",
		"errorMsg":  "success",
	})
	return r.Exit(200)
}
