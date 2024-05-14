package api

import (
	"github.com/MohammadAzhari/aqary_task/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewServer(q *db.Queries, pool *pgxpool.Pool) {
	s := &Server{
		q:    q,
		conn: pool,
	}

	s.initRoutes()
}

func (s *Server) initRoutes() {
	r := gin.Default()

	r.POST("/api/users", s.bindParamsToHandler(CreateUser))

	r.GET("/api/users", s.bindParamsToHandler(GetUsers))

	r.POST("/api/users/generateotp", s.bindParamsToHandler(GenerateOtp))

	r.POST("/api/users/verifyotp", s.bindParamsToHandler(VerifyOtp))

	r.Run()
}

func (s *Server) bindParamsToHandler(fn func(*gin.Context, *db.Queries, *pgxpool.Pool)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(ctx, s.q, s.conn)
	}
}
