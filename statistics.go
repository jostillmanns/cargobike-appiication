package main

import (
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

var (
	colorTransparent    = drawing.Color{R: 241, G: 241, B: 241, A: 0}
	colorWhite          = drawing.Color{R: 241, G: 241, B: 241, A: 255}
	colorMariner        = drawing.Color{R: 60, G: 100, B: 148, A: 255}
	colorLightSteelBlue = drawing.Color{R: 182, G: 195, B: 220, A: 255}
	colorPoloBlue       = drawing.Color{R: 126, G: 155, B: 200, A: 255}
	colorSteelBlue      = drawing.Color{R: 73, G: 120, B: 177, A: 255}
	colorLast           = drawing.Color{R: 121, G: 140, B: 177, A: 255}
	barWidth            = 100
	barSpacing          = 10
	strokeWidth         = 0.1
	fontSize            = float64(7)
)

func plotSurveyA() chart.StackedBarChart {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	return chart.StackedBarChart{
		Title:      "Quarterly Sales",
		TitleStyle: chart.Hidden(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 0,
			},
		},
		Width:  300,
		Height: 500,
		XAxis: chart.Style{
			FontSize: fontSize,
		},
		YAxis: chart.Style{
			FontSize: fontSize,
		},
		BarSpacing: barSpacing,
		Bars: []chart.StackedBar{
			{
				Name:  "Würden Sie in der Zukunft wieder ein Lastenrad nutzen?",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "74.4% Ja",
						Value: 74.4,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "18.6%",
						Value: 18.6,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "5.3%",
						Value: 5.3,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Würden Sie mittelfristig oder sofort ein Lastenrad kaufen?",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "",
						Value: 100 - 20.2 - 14.5 - 25.7 - 20.5 - 17.2,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorTransparent,
							FontColor:   colorTransparent,
						},
					},
					{
						Label: "20.2% Ja",
						Value: 20.2,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "14.5%",
						Value: 14.5,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "25.7%",
						Value: 25.7,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "20.5%",
						Value: 20.5,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "17.2%",
						Value: 17.2,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorLast,
							FontColor:   colorWhite,
						},
					},
				},
			},
		},
	}
}

func plotSurveyB() chart.StackedBarChart {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	return chart.StackedBarChart{
		Title: "Wie hätten Sie ohne ein Lastenrad-Sharing-Service den Transport bewältigt?",
		TitleStyle: chart.Style{
			FontSize: fontSize,
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top: 100,
			},
		},
		Width:  600,
		Height: 500,
		XAxis: chart.Style{
			FontSize: fontSize,
		},
		YAxis:      chart.Hidden(),
		BarSpacing: barSpacing,
		Bars: []chart.StackedBar{
			{
				Name:  "Zu Fuss",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "",
						Value: 100 - 3.3,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: 0,
							FillColor:   colorTransparent,
							FontColor:   colorTransparent,
							StrokeColor: colorTransparent,
						},
					},
					{
						Label: "3.3%",
						Value: 3.3,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Öffentliche Verkehrsmittel",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "%",
						Value: 100 - 9.6,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: 0,
							FillColor:   colorTransparent,
							StrokeColor: colorTransparent,
							FontColor:   colorTransparent,
						},
					},
					{
						Label: "9.6%",
						Value: 9.6,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: strokeWidth,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Mit dem Fahrrad",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "",
						Value: 100 - 27.7,
						Style: chart.Style{
							FontSize:    fontSize,
							StrokeWidth: 0,
							FillColor:   colorTransparent,
							StrokeColor: colorTransparent,
							FontColor:   colorTransparent,
						},
					},
					{
						Label: "27.7%",
						Value: 27.7,
						Style: chart.Style{
							StrokeWidth: strokeWidth,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    fontSize,
						},
					},
				},
			},
			{
				Name:  "Mit dem Auto",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "4.3% Geliehen",
						Value: 4.3,
						Style: chart.Style{
							StrokeWidth: strokeWidth,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    fontSize,
						},
					},
					{
						Label: "25% Carsharing",
						Value: 25,
						Style: chart.Style{
							StrokeWidth: strokeWidth,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
							FontSize:    fontSize,
						},
					},
					{
						Label: "16.1% Eigenes Auto",
						Value: 16.1,
						Style: chart.Style{
							StrokeWidth: strokeWidth,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
							FontSize:    fontSize,
						},
					},
				},
			},
			{
				Name:  "Ich hätte den Transport nicht gemacht.",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "",
						Value: 100 - 12.8,
						Style: chart.Style{
							FillColor:   colorTransparent,
							StrokeColor: colorTransparent,
							FontColor:   colorTransparent,
							FontSize:    fontSize,
						},
					},
					{
						Label: "12.8%",
						Value: 12.8,
						Style: chart.Style{
							StrokeWidth: strokeWidth,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    fontSize,
						},
					},
				},
			},
		},
	}
}
