package http

type Handler func(*Request) LambdaResponseBuilder
