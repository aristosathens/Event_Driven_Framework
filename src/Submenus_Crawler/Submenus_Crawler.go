package Submenus

import ()

// ------------------------------------------- Definitions ------------------------------------------- //

type MenuType int

const (
	MAIN         MenuType = 1
	ADD_WEBSITES MenuType = 2
)

type Menu struct {
	menuType MenuType
	elements []MenuItem
}

type MenuItem struct {
	label string
	key   string
	response
}

// ------------------------------------------- Public ------------------------------------------- //
