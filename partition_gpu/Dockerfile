# Copyright 2017 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM --platform=$BUILDPLATFORM golang:1.24-bullseye as builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /go/src/github.com/GoogleCloudPlatform/container-engine-accelerators
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o gpu_partitioner partition_gpu/partition_gpu.go
RUN chmod a+x /go/src/github.com/GoogleCloudPlatform/container-engine-accelerators/gpu_partitioner

FROM us.gcr.io/gke-release/gke-distroless/bash:gke_distroless_20250407.00_p0@sha256:b903ad51976ccd68a817f445c5b6df8bf3655bfd1b30ca989472f7a99f930fc5
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/container-engine-accelerators/gpu_partitioner /usr/bin/gpu_partitioner
CMD ["/usr/bin/gpu_partitioner", "-logtostderr"]
