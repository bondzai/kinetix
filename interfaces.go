// Copyright 2024 The strom Authors
// This file is part of the strom library.
//
// The strom library is licensed under the MIT License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// The strom library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// MIT License for more details.
//
// You should have received a copy of the MIT License
// along with the strom library. If not, see <https://opensource.org/licenses/MIT>.

// Package strom defines interfaces and utilities for Server-Sent Events (SSE).
package storm

import "context"

type EventSource interface {
	Stream(ctx context.Context) error
}
