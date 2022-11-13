package models

type Role struct {
	Name            string `gorm:"primaryKey"`
	Description     string
	RolePermissions []RolePermission `gorm:"foreignKey:RoleName;references:Name"`
}

type RolePermission struct {
	ID           string `gorm:"primaryKey"`
	RoleName     string
	PermissionID string
	Permission   `gorm:"foreignKey:PermissionID;references:ID"`
}
