package main

import "sync"

var (
	users = map[int]User{}
	posts = map[int]Post{}

	userMu = sync.RWMutex{}
	postMu = sync.RWMutex{}
)
