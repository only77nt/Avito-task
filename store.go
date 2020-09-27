package main

type Store interface {
	TransformUrl()
	GetUrlFromDB()
}