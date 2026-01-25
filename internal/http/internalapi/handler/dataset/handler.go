package dataset

import (
	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/http/appctx"
	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/http/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (r *handler) createDatasetHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	var data api.Dataset
	if err := validator.BindAndValidate(cc, &data); err != nil {
		logrus.WithError(err).Error("failed to bind and validate token")
		return echo.NewHTTPError(500, err)
	}

	organizationID := cc.OrganizationID()

	dataset, err := r.datasetService.Create(ctx, organizationID, data)
	if err != nil {
		logrus.WithError(err).Error("failed to create dataset")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(dataset))
}

func (r *handler) listDatasetsHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	organizationID := cc.OrganizationID()

	datasets, err := r.datasetService.List(ctx, organizationID)
	if err != nil {
		logrus.WithError(err).Error("failed to list datasets")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(datasets))
}

func (r *handler) listDatasetFieldsHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	organizationID := cc.OrganizationID()
	datasetID := cc.Param("datasetId")

	fields, err := r.datasetService.ListFields(ctx, organizationID, datasetID)
	if err != nil {
		logrus.WithError(err).Error("failed to list dataset fields")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(fields))
}

func (r *handler) queryDatasetLogsHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	organizationID := cc.OrganizationID()
	datasetID := cc.Param("datasetId")

	var data api.ListLogsSchema
	if err := validator.BindAndValidate(cc, &data); err != nil {
		logrus.WithError(err).Error("failed to bind and validate token")
		return echo.NewHTTPError(500, err)
	}

	logs, err := r.datasetService.QueryLogs(ctx, organizationID, datasetID, data.FromDate, data.ToDate, data.Filter, data.Cursor)
	if err != nil {
		logrus.WithError(err).Error("failed to query dataset logs")
		return echo.NewHTTPError(500, err)
	}

	var response api.DatasetLogsResponse
	response.Data = logs
	if response.Data == nil {
		response.Data = make([]api.DatasetLog, 0)
	}

	response.Meta.TotalRowCount = len(logs)

	return cc.JSON(200, httputil.Response(response))
}

func (r *handler) getDatasetLogHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	organizationID := cc.OrganizationID()
	datasetID := cc.Param("datasetId")
	logID := cc.Param("logId")

	log, err := r.datasetService.GetLog(ctx, organizationID, datasetID, logID)
	if err != nil {
		logrus.WithError(err).Error("failed to get dataset log")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(log))
}
