module github.com/sudhin-az/api-gateway

go 1.24.1

require (
    github.com/gin-gonic/gin v1.9.0
    google.golang.org/grpc v1.74.2
)
replace github.com/sudhin-az/FOOD-ORDERING/user-service => ../user-service
replace github.com/sudhin-az/FOOD-ORDERING/menu-service => ../menu-service
replace github.com/sudhin-az/FOOD-ORDERING/order-service => ../order-service