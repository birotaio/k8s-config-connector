// Copyright 2023 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package iam -var YAML_workload_identity_pool blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/workload_identity_pool.yaml

package iam

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/workload_identity_pool.yaml
var YAML_workload_identity_pool = []byte("info:\n  title: Iam/WorkloadIdentityPool\n  description: The Iam WorkloadIdentityPool resource\n  x-dcl-struct-name: WorkloadIdentityPool\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a WorkloadIdentityPool\n    parameters:\n    - name: workloadIdentityPool\n      required: true\n      description: A full instance of a WorkloadIdentityPool\n  apply:\n    description: The function used to apply information about a WorkloadIdentityPool\n    parameters:\n    - name: workloadIdentityPool\n      required: true\n      description: A full instance of a WorkloadIdentityPool\n  delete:\n    description: The function used to delete a WorkloadIdentityPool\n    parameters:\n    - name: workloadIdentityPool\n      required: true\n      description: A full instance of a WorkloadIdentityPool\n  deleteAll:\n    description: The function used to delete all WorkloadIdentityPool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many WorkloadIdentityPool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    WorkloadIdentityPool:\n      title: WorkloadIdentityPool\n      x-dcl-id: projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: A description of the pool. Cannot exceed 256 characters.\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: Whether the pool is disabled. You cannot use a disabled pool\n            to exchange tokens, or use existing tokens to access resources. If the\n            pool is re-enabled, existing tokens grant access again.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: A display name for the pool. Cannot exceed 32 characters.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the pool.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: WorkloadIdentityPoolStateEnum\n          readOnly: true\n          description: 'Output only. The state of the pool. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETED\n")

// 3454 bytes
// MD5: 251cef3655ba232ae1adaba39bc26327