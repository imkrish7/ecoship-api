package main

import (
	"log"
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/imkrish7/ecoship-api/application"
	"github.com/imkrish7/ecoship-api/controllers"
	"github.com/imkrish7/ecoship-api/gen/proto/auth/v1/authv1connect"
	repositories "github.com/imkrish7/ecoship-api/repositories/auth"
)

func AppRoutes(app *application.Application) *chi.Mux {
	r := chi.NewRouter()

	authRepo := repositories.NewAuthRepository(app.DB)

	auth := controllers.NewAuthServer(app, authRepo)

	// interceptor := connect.WithInterceptors(interceptors.AuthInterceptors())

	authPath, authServiceHandler := authv1connect.NewAuthServiceHandler(auth)

	reflector := grpcreflect.NewStaticReflector(
		authv1connect.AuthServiceName,
	)

	grpcPath, grpcHandler := grpcreflect.NewHandlerV1(reflector)
	grpcv1Path, grpcv1Handler := grpcreflect.NewHandlerV1Alpha(reflector)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS", "POST", "PATCH", "PUT"},
		AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRF-Token", "Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Grpc-Timeout",
			"X-Grpc-Web",
		},
		ExposedHeaders: []string{"Link", "Content-Encoding",
			"Connect-Content-Encoding",
			"Grpc-Status",
			"Grpc-Message"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Mount(authPath, http.StripPrefix("/api", authServiceHandler))

		r.Mount(grpcPath, http.StripPrefix("/api", grpcHandler))
		r.Mount(grpcv1Path, http.StripPrefix("/api", grpcv1Handler))
	})
	log.Print("routing...")
	log.Print(r)
	return r
}
