package handler

type HandlerCtx interface {
	JSON(code int, obj interface{})
}
