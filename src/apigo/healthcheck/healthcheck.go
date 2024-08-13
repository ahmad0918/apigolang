package healthcheck

import (
	Conn "apigolang/src/apigo/connection"
	"apigolang/src/apigo/models"
	"apigolang/src/config/db"
	"database/sql"
	"net/http"

	"github.com/alexcesaro/log/stdlog"
	"github.com/gin-gonic/gin"
)

type CheckStatus struct {
	Name            string `json:"name" example:"Service Apigolang"`
	Pass            bool   `json:"pass" example:"true"`
	ResponseMessage string `json:"responsiveness" example:"Success hit api golang"`
}

type Config struct {
	HealthPath  string
	Method      string
	StatusOK    int
	StatusNotOK int
}

type Check interface {
	Pass() Response
	Name(string) string
}

type Response struct {
	Pass            bool
	ResponseCode    int
	ResponseMessage string
}

var logkoe = stdlog.GetFromFlags()

func Controller(checks []Check, service string, config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		statuses := CheckStatus{
			Name:            service,
			Pass:            true,
			ResponseMessage: "success hit api golang",
		}

		httpStatus := config.StatusOK
		for _, check := range checks {
			response := check.Pass()
			statuses = CheckStatus{
				Name:            check.Name(service),
				Pass:            response.Pass,
				ResponseMessage: response.ResponseMessage,
			}

			if !response.Pass {
				httpStatus = config.StatusNotOK
				logkoe.Info("msg =", "ALL DATA FOUND", "detail =", statuses)
				c.JSON(httpStatus, statuses)
				return
			}
		}

		logkoe.Info("msg =", "ALL DATA FOUND", "detail =", statuses)
		response := models.SuccessResponse{
			Message:  "Operation successful",
			Response: http.StatusOK,
			Result:   statuses,
		}
		c.JSON(http.StatusOK, response)
	}
}

func Service(engine *gin.Engine, service string, config Config, checks []Check) {
	engine.Handle(config.Method, config.HealthPath, Controller(checks, service, config))
	return
}

type SqlCheck struct {
	Sql *sql.DB
}

func (s SqlCheck) Pass() Response {
	if s.Sql == nil {
		return Response{
			Pass:            false,
			ResponseCode:    http.StatusServiceUnavailable,
			ResponseMessage: "Cannot connect",
		}
	}
	if err := s.Sql.Ping(); err != nil {
		return Response{
			Pass:            false,
			ResponseCode:    http.StatusServiceUnavailable,
			ResponseMessage: "Cannot connect",
		}
	}
	return Response{
		Pass:            true,
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Successfully connected",
	}
}

func (s SqlCheck) Name(service string) string {
	return service
}

// DefaultConfig godoc
// @Summary Get Health Check Configuration
// @Description Retrieve the configuration for the health check endpoint, including the path, method, and status codes.
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse{result=healthcheck.CheckStatus} "Health check configuration"
// @Failure 500 {object} models.InternalErrorResponse "Internal Server Error"
// @Router /healthcheck/config  [get]
func DefaultConfig() Config {
	return Config{
		HealthPath:  "/api/golang/public/healthcheck/config",
		Method:      http.MethodGet,
		StatusOK:    http.StatusOK,
		StatusNotOK: http.StatusServiceUnavailable,
	}
}

type DBJson struct {
	DbStatus string `json:"dbStatus" example:"Database Golang is not ready"`
	DbName   string `json:"dbName" example:"Golang"`
}

// GetDBHealthCheck godoc
// @Summary Check Database Health Status
// @Description Retrieve the health status of the database connections, returning the status of each connection.
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse{result=[]healthcheck.DBJson} "All DB Connection Found and Ready"
// @Failure 503 {object} models.ServiceUnavailableResponse "Some Services are Down"
// @Router /healthcheck/database  [get]
func GetDBHealthCheck(c *gin.Context) {
	statArr := make([]DBJson, 0)
	status := http.StatusOK
	message := "All DB Connection Found and Ready"

	// DB BAF Kiosk
	cfgBafUser := db.CreateConnection()
	conn, err := Conn.NewPostgresConnection(cfgBafUser)
	dbStatusK := "DB Golang is ready"
	if err != nil || conn == nil {
		dbStatusK = "Database Golang is not ready"
	}

	if err != nil || conn == nil {
		status = http.StatusServiceUnavailable
		message = "Some Services are Down"
	}
	dbJson := DBJson{dbStatusK, "Golang"}
	statArr = append(statArr, dbJson)

	logkoe.Info("msg =", "All DB Connection Found and Ready", "detail =", statArr)
	c.JSON(status, gin.H{
		"response": status,
		"message":  message,
		"result":   statArr,
	})
}
