/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dubbo

// Response payload type enum
const (
	RESPONSE_WITH_EXCEPTION                  int32 = 0
	RESPONSE_VALUE                           int32 = 1
	RESPONSE_NULL_VALUE                      int32 = 2
	RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS int32 = 3
	RESPONSE_VALUE_WITH_ATTACHMENTS          int32 = 4
	RESPONSE_NULL_VALUE_WITH_ATTACHMENTS     int32 = 5
)