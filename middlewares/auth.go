package middlewares

import (
	"net/http"
	"strings"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
)

// Secret key for JWT (sebaiknya disimpan di env)
var secretKey = []byte("secret_jwt_key")

// AuthMiddleware untuk validasi JWT
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
	}

	// Pastikan token memiliki prefix "Bearer "
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}
	tokenString := tokenParts[1]

	// Parse JWT token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_jwt_key"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Debugging - Print JWT claims
	fmt.Println("DEBUG JWT Claims =", claims)

	// Ambil user ID
	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}
	userID := uint(userIDFloat)

	// Ambil is_admin dari token
	isAdmin, adminOk := claims["is_admin"].(bool)
	if !adminOk {
		fmt.Println("DEBUG: is_admin bukan bool, coba baca sebagai float64")
		isAdminFloat, floatOk := claims["is_admin"].(float64)
		isAdmin = floatOk && isAdminFloat == 1
	}

	// Debugging - Print hasil parsing
	fmt.Println("DEBUG: user_id =", userID, "| is_admin =", isAdmin)

	// Simpan ke Fiber context
	c.Locals("user_id", userID)
	c.Locals("is_admin", isAdmin)

	return c.Next()
}

// AdminMiddleware memastikan user memiliki is_admin = 1
func AdminMiddleware(c *fiber.Ctx) error {
    isAdmin, ok := c.Locals("is_admin").(bool)
    fmt.Println("DEBUG AdminMiddleware: is_admin = ", isAdmin, "ok = ", ok)

    if !ok || !isAdmin {
        return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Admins only"})
    }
    return c.Next()
}
