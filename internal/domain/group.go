package domain

import "context"

type Group struct {
	Code        string
	Name        string
	Description string
	Internal    bool
}

type Role struct {
	Code        string
	Name        string
	Description string
	Internal    bool
}

type Resource struct {
	Code        string
	Name        string
	Description string
	Internal    bool
}

type RWX struct {
	R bool
	W bool
	X bool
}

type Permission struct {
	Allow RWX
	Deny  RWX
}

type RoleResourcePermission struct {
	RoleCode        string
	ResourcePattern string
	Permission      Permission
}

type RoleWildCardPermission struct {
	RoleCode        string
	ResourcePattern string
	Permission      Permission
}

type SecurityServices interface {
	// get group by code
	GetGroup(ctx context.Context, code string) (*Group, error)
	// get all group except deleted
	GetAllGroup(ctx context.Context) ([]*Group, error)
	// get role by code
	GetCode(ctx context.Context, code string) (*Role, error)
	// get all role except deleted
	GetAllRole(ctx context.Context) ([]*Role, error)
	// get roles assigned to a group
	GetRolesOnGroups(ctx context.Context, groups []string) ([]string, error)
	// get resource by code
	GetResource(ctx context.Context, code string) (*Resource, error)
	// get all resources except deleted
	GetAllResource(ctx context.Context) ([]*Resource, error)
	// calculate permissions on resource for the roles and applies allow/deny RWX logic
	GetGrantedPermissions(ctx context.Context, resource string, roles []string) (*RWX, error)
	// check if roles have permission on resource
	CheckPermissions(ctx context.Context, resource string, roles []string, requestedPermissions []string) (bool, error)
	// returns permission on resource / roles setup explicitly
	GetExplicitPermissions(ctx context.Context, resource []string, roles []string) ([]*RoleResourcePermission, error)
	// returns wildcard permission on roles
	GetWildCardPermissions(ctx context.Context, roles []string) ([]*RoleWildCardPermission, error)
}

type SecurityStorage interface {
	// GetGroup retrieves a group by code
	GetGroup(ctx context.Context, code string) (*Group, error)
	// GetGroups retrieves all not deleted groups
	GetGroups(ctx context.Context) ([]*Group, error)
	// GetRole retrieves a role by code
	GetRole(ctx context.Context, code string) (*Role, error)
	// GetAllRoles retrieves all not deleted roles
	GetAllRoles(ctx context.Context) ([]*Role, error)
	// GetAllRoleCodes retrieves all role codes
	GetAllRoleCodes(ctx context.Context) ([]string, error)
	// GetResource retrieves a resource by code
	GetResource(ctx context.Context, code string) (*Resource, error)
	// GetAllResources retrieves all not deleted resources
	GetAllResources(ctx context.Context) ([]*Resource, error)
	// ResourceExplicitPermissionsExists checks if there are explicit (no wildcard) permissions on the resource
	ResourceExplicitPermissionsExists(ctx context.Context, code string) (bool, error)
	// GetRoleCodesForGroups retrieves role codes for groups
	GetRoleCodesForGroups(ctx context.Context, groups []string) ([]string, error)
	// GroupsWithRoleExists checks if there are groups with assigned role
	GroupsWithRoleExists(ctx context.Context, role string) (bool, error)
	// GetPermissions retrieves permissions granted to roles on resource
	GetPermissions(ctx context.Context, resource string, roles []string) ([]*Permission, error)
	// GetWildcardPermissions retrieves wildcard permissions granted to roles on resource
	GetWildcardPermissions(ctx context.Context, resource string, roles []string) ([]*Permission, error)
}
