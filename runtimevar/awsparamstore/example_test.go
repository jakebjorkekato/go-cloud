// Copyright 2018 The Go Cloud Development Kit Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsparamstore_test

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/awsparamstore"
)

func ExampleOpenVariable() {
	// This example is used in https://gocloud.dev/howto/runtimevar/#awsps-ctor

	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	sess, err := session.NewSession(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Construct a *runtimevar.Variable that watches the variable.
	v, err := awsparamstore.OpenVariable(sess, "cfg-variable-name", runtimevar.StringDecoder, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()
}

func Example_openVariableFromURL() {
	// This example is used in https://gocloud.dev/howto/runtimevar/#awsps

	// import _ "gocloud.dev/runtimevar/awsparamstore"

	// runtimevar.OpenVariable creates a *runtimevar.Variable from a URL.

	// Variables set up elsewhere:
	ctx := context.Background()

	v, err := runtimevar.OpenVariable(ctx, "awsparamstore://myvar?region=us-west-1&decoder=string")
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()
}
