package vql

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"www.velocidex.com/golang/velociraptor/acls"
	acl_proto "www.velocidex.com/golang/velociraptor/acls/proto"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	"www.velocidex.com/golang/velociraptor/logging"
	"www.velocidex.com/golang/vfilter"
)

const (
	ACL_MANAGER_VAR = "$acl"
)

type ACLManager interface {
	CheckAccess(permission acls.ACL_PERMISSION, args ...string) (bool, error)
}

// NullACLManager is an acl manager which allows everything. This is
// currently used on the client and on the command line where there is
// no clear principal or ACL controls.
type NullACLManager struct{}

func (self NullACLManager) CheckAccess(
	permission acls.ACL_PERMISSION, args ...string) (bool, error) {
	return true, nil
}

// ServerACLManager is used when running server side VQL to control
// ACLs on various VQL plugins.
type ServerACLManager struct {
	principal string
	Token     *acl_proto.ApiClientACL
}

func (self *ServerACLManager) CheckAccess(
	permission acls.ACL_PERMISSION, args ...string) (bool, error) {
	return acls.CheckAccessWithToken(self.Token, permission, args...)
}

// NewRoleACLManager creates an ACL manager with only the assigned
// roles. This is useful for creating limited VQL permissions
// internally.
func NewRoleACLManager(role string) ACLManager {
	policy := &acl_proto.ApiClientACL{}
	acls.GetRolePermissions(nil, []string{role}, policy)
	return &ServerACLManager{Token: policy}
}

func NewServerACLManager(
	config_obj *config_proto.Config,
	principal string) ACLManager {
	policy, err := acls.GetEffectivePolicy(config_obj, principal)
	if err != nil {
		logger := logging.GetLogger(config_obj, &logging.FrontendComponent)
		logger.WithFields(logrus.Fields{
			"user":  principal,
			"error": err,
		}).Error("Unable to get policy")
		policy = &acl_proto.ApiClientACL{}
	}

	return &ServerACLManager{principal: principal, Token: policy}
}

// Check access through the ACL manager in the scope.  NOTE: This
// assumes it is not possible for a user to mask the ACL manager in
// the scope! There is currently no way to create an acl manager type
// from within VQL so this is a safe assumption - if a user was to
// override the ACL_MANAGER_VAR with something else this will lock
// down the entire VQL ACL system and deny all permissions.
func CheckAccess(scope *vfilter.Scope, permission acls.ACL_PERMISSION) error {
	manager_any, pres := scope.Resolve(ACL_MANAGER_VAR)
	if !pres {
		return errors.New(fmt.Sprintf("Permission denied: %s", permission))
	}

	manager, ok := manager_any.(ACLManager)
	if !ok {
		return errors.New(fmt.Sprintf("Permission denied: %s", permission))
	}

	perm, err := manager.CheckAccess(permission)
	if !perm || err != nil {
		return errors.New(fmt.Sprintf("Permission denied: %s", permission))
	}

	return nil
}

func CheckFilesystemAccess(scope *vfilter.Scope, accessor string) error {
	switch accessor {

	// These accessor are OK to use at any time.
	case "data":
		return nil

	default:
		return CheckAccess(scope, acls.FILESYSTEM_READ)
	}
}

// Get the principal that is running the query if possible.
func GetPrincipal(scope *vfilter.Scope) string {
	manager_any, pres := scope.Resolve(ACL_MANAGER_VAR)
	if !pres {
		return ""
	}

	manager, ok := manager_any.(*ServerACLManager)
	if !ok {
		return ""
	}

	return manager.principal
}
