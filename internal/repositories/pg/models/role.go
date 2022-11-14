package models

type Role struct {
	Name            string `gorm:"primaryKey"`
	Description     string
	RolePermissions []RolesPermission `gorm:"foreignKey:RoleName;references:Name"`
}

type RolesPermission struct {
	ID           string `gorm:"primaryKey"`
	RoleName     string
	PermissionID string
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
}
