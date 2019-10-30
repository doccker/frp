// Copyright 2016 doccker, doccker@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consts

var (
	// proxy status
	Idle    string = "idle"
	Working string = "working"
	Closed  string = "closed"
	Online  string = "online"
	Offline string = "offline"

	// proxy type
	TcpProxy   string = "tcp"
	UdpProxy   string = "udp"
	HttpProxy  string = "http"
	HttpsProxy string = "https"
	StcpProxy  string = "stcp"
	XtcpProxy  string = "xtcp"
)
