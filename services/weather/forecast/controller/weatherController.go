package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/steady-bytes/draft/pkg/chassis"

	v1 "github.com/AndrewSC208/weather-api/api/weather/forecast/v1"
	weatherV1Connect "github.com/AndrewSC208/weather-api/api/weather/forecast/v1/v1connect"
)

type (
	RpcHandler interface {
		chassis.RPCRegistrar
		weatherV1Connect.ForecastServiceHandler
	}

	weatherController struct {
		logger chassis.Logger
	}

	WeatherStationResponse struct {
		StationProperties *WeatherStationProperties `json:"properties"`
	}

	WeatherStationProperties struct {
		ForecastURL string `json:"forecast"`
	}

	ForecastResponse struct {
		ForecastProperties *ForecastProperties `json:"properties"`
	}

	ForecastProperties struct {
		Units   string            `json:"units"`
		Periods []*ForecastPeriod `json:"periods"`
	}

	ForecastPeriod struct {
		Number          int       `json:"number"`
		Name            string    `json:"name"`
		StartTime       time.Time `json:"startTime"`
		EndTime         time.Time `json:"endTime"`
		IsDaytime       bool      `json:"isDaytime"`
		Temperature     int       `json:"temperature"`
		TemperatureUnit string    `json:"temperatureUnit"`
		WindSpeed       string    `json:"windSpeed"`
		ShortForecast   string    `json:"shortForecast"`
	}
)

const (
	ErrFailedToGetPointsForCoordinates = "failed to get points for provided coordinates"
	ErrFailedToGetForecast             = "failed to get forecast from station"
	ErrFailedToCategorizeForecast      = "failed to categorize forecast"
	ErrInvalidTemperature              = "invalid temperature to categorize"

	ErrInvalidRequest                = "request is invalid"
	ErrFailedToAssembleHttpRequest   = "failed to assemble get request"
	ErrFailedToReadHttpResponseBody  = "failed to read http response body"
	ErrFailedToUnmarshalHttpResponse = "failed to unmarshal http response"
	ErrFailedHttpRequest             = "failed to execute http request"
)

func NewWeatherController(logger chassis.Logger) RpcHandler {
	return &weatherController{
		logger: logger,
	}
}

func (c *weatherController) RegisterRPC(server chassis.Rpcer) {
	pattern, handler := weatherV1Connect.NewForecastServiceHandler(c)
	server.AddHandler(pattern, handler, true)
}

func (c *weatherController) ShortForecast(ctx context.Context, request *connect.Request[v1.ShortForecastRequest]) (*connect.Response[v1.ShortForecastResponse], error) {
	c.logger.Info("short forecast")

	if err := request.Msg.Validate(); err != nil {
		c.logger.WithError(err).Error(ErrInvalidRequest)
		return nil, errors.New(ErrInvalidRequest)
	}

	c.logger.Info("get points for coordinates")
	points, err := c.getPoints(request.Msg.Latitude, request.Msg.Longitude)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetPointsForCoordinates)
		return nil, err
	}

	c.logger.Info("get forecast")
	forecast, err := c.getForecast(points.StationProperties.ForecastURL)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetForecast)
		return nil, err
	}

	c.logger.Info("categorize forecast")
	category, err := c.categorizeForecast(forecast)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToCategorizeForecast)
		return nil, err
	}

	return connect.NewResponse(&v1.ShortForecastResponse{
		ShortForecast:       forecast.ForecastProperties.Periods[0].ShortForecast,
		TemperatureCategory: category,
	}), nil
}

func (c *weatherController) getPoints(lat, long float32) (*WeatherStationResponse, error) {
	url := fmt.Sprintf("https://api.weather.gov/points/%.2f,%.2f", lat, long)

	body, err := c.httpGet(url)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetPointsForCoordinates)
		return nil, errors.New(ErrFailedToGetPointsForCoordinates)
	}

	station := &WeatherStationResponse{}
	if err := json.Unmarshal(body, station); err != nil {
		c.logger.WithError(err).Error(ErrFailedToUnmarshalHttpResponse)
		return nil, errors.New(ErrFailedToUnmarshalHttpResponse)
	}

	return station, nil
}

func (c *weatherController) getForecast(forecastURL string) (*ForecastResponse, error) {
	body, err := c.httpGet(forecastURL)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetForecast)
	}

	forecastRes := &ForecastResponse{}
	if err := json.Unmarshal(body, forecastRes); err != nil {
		c.logger.WithError(err).Error(ErrFailedToUnmarshalHttpResponse)
		return nil, errors.New(ErrFailedToUnmarshalHttpResponse)
	}

	return forecastRes, nil
}

func (c *weatherController) categorizeForecast(forecast *ForecastResponse) (v1.ForecastCharacterization, error) {
	if len(forecast.ForecastProperties.Periods) < 1 {
		c.logger.Error(ErrInvalidTemperature)
		return v1.ForecastCharacterization_INVALID, errors.New(ErrInvalidTemperature)
	}

	switch t := forecast.ForecastProperties.Periods[0].Temperature; {
	case t < 55:
		return v1.ForecastCharacterization_COLD, nil
	case t < 85:
		return v1.ForecastCharacterization_MODERATE, nil
	case t >= 85:
		return v1.ForecastCharacterization_HOT, nil
	default:
		return v1.ForecastCharacterization_INVALID, errors.New(ErrInvalidTemperature)
	}
}

func (c *weatherController) httpGet(url string) ([]byte, error) {
	client := &http.Client{}

	res, err := client.Get(url)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedHttpRequest)
		return nil, errors.New(ErrFailedHttpRequest)
	}

	if res.StatusCode != http.StatusOK {
		c.logger.WithError(err).Error(ErrFailedHttpRequest)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToReadHttpResponseBody)
		return nil, errors.New(ErrFailedToReadHttpResponseBody)
	}

	return body, nil
}
