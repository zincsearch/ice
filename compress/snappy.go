/* Copyright 2022 Zinc Labs Inc. and Contributors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package compress

import "github.com/golang/snappy"

// SnappyDecompress decompresses a block using snappy algorithm.
func SnappyDecompress(dst, src []byte) ([]byte, error) {
	return snappy.Decode(dst, src)
}

// SnappyCompress compresses a block using snappy algorithm.
func SnappyCompress(dst, src []byte) ([]byte, error) {
	return snappy.Encode(dst, src), nil
}
