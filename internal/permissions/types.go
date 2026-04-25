package permissions

// PermissionConfig holds permission configuration
type PermissionConfig struct {
	AllowDelete         bool
	AllowModify         bool
	RequireConfirmation bool
	AdminOnly           bool
}
