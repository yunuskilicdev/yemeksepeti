package main

type KeyValueService interface {
	Get(key string) string
	Put(key string, value string)
}
