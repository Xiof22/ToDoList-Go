package repository

import ()

type Repository interface {
	Create(title, description string)
}
