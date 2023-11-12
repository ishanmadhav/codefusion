package server

func (s *Server) setupRoutes() {
	s.setupCodeRoutes()
	s.setupRoomRoutes()
}

// Contains basic user and auth routes
func (s *Server) setupUserRoutes() {

}

// Contains basic code routes, submit, fetch, etc.
func (s *Server) setupCodeRoutes() {
	s.App.Get("/code/:id", s.getCodeExecutionResultsByID)
	s.App.Post("/code", s.submitCode)
	s.App.Get("/codes", s.getAllCodes)
	s.App.Delete("/code/:id", s.deleteCodeByID)
}

// Interview room in which the collbaoritve editor exists is init here
// then we start with the web socket part
func (s *Server) setupRoomRoutes() {
	s.App.Get("/room/:id", s.getRoomByID)
	s.App.Post("/room", s.createRoom)
	s.App.Post("/room/codes/:roomID", s.GetCodesByRoomID)
}
