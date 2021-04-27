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
	barWidth            = 120
)

func plotSurveyA() chart.StackedBarChart {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	return chart.StackedBarChart{
		Title:      "Quarterly Sales",
		TitleStyle: chart.Hidden(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 100,
			},
		},
		Width:      800,
		Height:     500,
		XAxis:      chart.Shown(),
		YAxis:      chart.Shown(),
		BarSpacing: 50,
		Bars: []chart.StackedBar{
			{
				Name:  "Würdest du in der Zukunft wieder ein Lastenrad nutzen?",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "74.4% Ja",
						Value: 74.4,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "18.6%",
						Value: 18.6,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "5.3%",
						Value: 5.3,
						Style: chart.Style{
							StrokeWidth: .01,
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
						Label: "20.2% Ja",
						Value: 20.2,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "14.5%",
						Value: 14.5,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "25.7%",
						Value: 25.7,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "20.5%",
						Value: 20.5,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "17.2%",
						Value: 17.2,
						Style: chart.Style{
							StrokeWidth: .01,
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
		Title:      "Ohne einen Lastenrad-Sharing-Service, wie hättest du den Transport bewältigt?",
		TitleStyle: chart.Hidden(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 100,
			},
		},
		Width:      800,
		Height:     500,
		XAxis:      chart.Shown(),
		YAxis:      chart.Shown(),
		BarSpacing: 50,
		Bars: []chart.StackedBar{
			{
				Name:  "Zu Fuss",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "",
						Value: 96.7,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorTransparent,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "3.3%",
						Value: 3.3,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    8,
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
						Value: 90.4,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorTransparent,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "9.6%",
						Value: 9.6,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    8,
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
						Value: 72.3,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorTransparent,
							FontColor:   colorWhite,
							FontSize:    8,
						},
					},
					{
						Label: "27.7%",
						Value: 27.7,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    8,
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
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    8,
						},
					},
					{
						Label: "25% Carsharing",
						Value: 25,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    8,
						},
					},
					{
						Label: "16.1% Eigenes Auto",
						Value: 16.1,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
							FontSize:    8,
						},
					},
				},
			},
		},
	}
}
