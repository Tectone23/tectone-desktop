package main

import "context"

type registerer interface {
	register(context.Context)
}
