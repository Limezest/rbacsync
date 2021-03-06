/*
   Copyright 2018 GM Cruise LLC

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package groups

import (
	"fmt"

	"github.com/pkg/errors"
	rbacv1 "k8s.io/api/rbac/v1"
)

// ErrNotFound us used by Grouper implementations to signal a group is not
// found. Can be used as the cause in errors.Wrap if more context is required.
//
// Use IsNotFound to detect this condition.
var ErrNotFound = fmt.Errorf("not found")

// IsNotFound returns true if the case of the error is ErrNotFound.
func IsNotFound(err error) bool {
	return errors.Cause(err) == ErrNotFound
}

// Grouper returns the members of a group, as RBAC Subjects, given a group
// name.
//
// If the group is unknown to the grouper, the error will return true for
// IsNotFound.
type Grouper interface {
	Members(group string) ([]rbacv1.Subject, error)
}

// Empty is a grouper that returns not found for all groups.
var Empty emptyGrouper

type emptyGrouper struct{}

func (e emptyGrouper) Members(group string) ([]rbacv1.Subject, error) {
	return nil, ErrNotFound
}
