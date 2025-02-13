package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/png"
	"net/http"

	"web/csv"
	"web/graphs"
	"web/utils"

	"github.com/labstack/echo/v4"
)

var Dataset = utils.Dataset{nil, nil}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func Data(c echo.Context) error {
	return c.Render(http.StatusOK, "data.html", Dataset)
}

func Set_dataset(c echo.Context) error {
	file, err := c.FormFile("dataset-file")
	if err != nil {
		return c.Render(500, "error.html", err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return c.Render(500, "error.html", err.Error())
	}
	defer src.Close()
	data, err := csv.Read_CSV(src)
	if err != nil {
		return c.Render(500, "error.html", err.Error())
	}

	Dataset.Inputs = csv.Create_slice(data)
	return c.Render(http.StatusOK, "data.html", Dataset)
}

func Browse_dataset(c echo.Context) error {
	s, _ := c.FormParams()

	selected_vars := s["headers"]
	final_inputs := make([]utils.Inputs, len(selected_vars))

	for i, selected := range selected_vars {
		for _, in := range Dataset.Inputs {
			if in.Header == selected {
				final_inputs[i] = in
				continue
			}
		}
	}
	Dataset.Inputs = final_inputs

	Dataset_with_plot := struct {
		Inputs          []utils.Inputs
		Base64PlotImage string
	}{Inputs: Dataset.Inputs, Base64PlotImage: ""}

	return c.Render(http.StatusOK, "inputs.html", Dataset_with_plot)
}

func Create_plotter(c echo.Context) error {
	independent_vars := c.FormValue("indep-val")
	dependent_vars := c.FormValue("dep-val")

	feature_slice := make([]utils.Inputs, 2)

	for _, in := range Dataset.Inputs {
		if in.Header == independent_vars {
			feature_slice[0] = in
		}
		if in.Header == dependent_vars {
			feature_slice[1] = in
		}
	}

	if len(feature_slice[0].Values) < 1 {
		return c.Render(500, "error.html", "No input values")
	}

	points := graphs.ToPlotter(utils.String_to_float_slice(feature_slice[0].Values), utils.String_to_float_slice(feature_slice[1].Values))
	p := graphs.CreateGraphFromData(fmt.Sprintf("%s-%s", independent_vars, dependent_vars), independent_vars, dependent_vars, "scatter", points)

	// Renderowanie wykresu do obrazu PNG w pamięci
	img, err := graphs.ImageFromPlot(p, 450, 300)
	if err != nil {
		return err
	}

	// Konwersja do formatu PNG
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	var Base64Encoding string
	Base64Encoding += "data:image/png;base64,"
	Base64Encoding += toBase64(buf.Bytes())

	Dataset_with_plot := struct {
		Inputs          []utils.Inputs
		Base64PlotImage template.URL
	}{Inputs: Dataset.Inputs, Base64PlotImage: template.URL(Base64Encoding)}

	return c.Render(http.StatusOK, "inputs.html", Dataset_with_plot)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
