// +build ignore

/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2019 MinIO, Inc.
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

package main

import (
	"fmt"
	"log"

	minio "github.com/SherifEldeeb/miniogo"
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY and my-bucketname are
	// dummy values, please replace them with original values.

	// Requests are always secure (HTTPS) by default. Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().

	// New returns an Amazon S3 compatible client object. API compatibility (v2 or v4) is automatically
	// determined based on the Endpoint value.
	s3Client, err := minio.New("s3.amazonaws.com", "YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", true)
	if err != nil {
		log.Fatalln(err)
	}

	// s3Client.TraceOn(os.Stderr)

	// Get object lock configuration.
	mode, validity, unit, err := s3Client.GetBucketObjectLockConfig("my-bucketname")
	if err != nil {
		log.Fatalln(err)
	}

	if mode != nil {
		fmt.Printf("%v mode is enabled for %v %v for bucket 'my-bucketname'\n", *mode, *validity, *unit)
	} else {
		fmt.Println("No mode is enabled for bucket 'my-bucketname'")
	}
}
