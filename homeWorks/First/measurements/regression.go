package measurements

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
	"strconv"
)

func linearRegression(arrX []int, arrY []int64) (float64, float64) {
	var arrYfloat64 []float64
	var arrXfloat64 []float64
	for _, v := range arrY {
		arrYfloat64 = append(arrYfloat64, float64(v))
	}
	for _, v := range arrX {
		arrXfloat64 = append(arrXfloat64, float64(v))
	}
	alpha, beta := stat.LinearRegression(arrXfloat64, arrYfloat64, nil, false)

	return alpha, beta
}

func main() {
	// Открываем CSV файл
	file, err := os.Open("Data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Читаем данные из файла
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Создаем новую регрессионную модель
	var r regression.Regression
	r.SetObserved("Score") // Зависимая переменная
	r.SetVar(0, "Grade")   // Независимая переменная

	var sizes []float64
	var values []float64

	// Обучаем модель на данных
	for _, record := range records[1:] { // Пропускаем заголовок
		size, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		value, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		sizes = append(sizes, size)
		values = append(values, value)

		r.Train(regression.DataPoint(value, []float64{size}))
	}

	// Запускаем регрессию
	r.Run()

	// Выводим формулу регрессии
	fmt.Printf("Regression Formula: %v\n", r.Formula)

	// Создаем график
	p := plot.New()

	p.Title.Text = "Линейная Регрессия"
	p.X.Label.Text = "Grade"
	p.Y.Label.Text = "Score"

	// Добавляем точки данных на график
	scatterData := make(plotter.XYs, len(sizes))
	for i := range sizes {
		scatterData[i].X = sizes[i]
		scatterData[i].Y = values[i]
	}

	scatter, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(scatter)

	// Добавляем регрессионную прямую
	lineData := make(plotter.XYs, 2)

	lineData[0].X = 1000 // Начальная оценка
	lineData[0].Y, err = r.Predict([]float64{lineData[0].X})
	if err != nil {
		fmt.Println(err)
	}

	lineData[1].X = 300000 // Конечная оценка
	lineData[1].Y, err = r.Predict([]float64{lineData[1].X})
	if err != nil {
		fmt.Println(err)
	}

	line, err := plotter.NewLine(lineData)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(line)

	// Сохраняем график в файл
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "regression_plot_test.png"); err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	// Точки для построения графика
//	points := []Point{
//		{1, 2},
//		{2, 4},
//		{3, 6},
//		{4, 8},
//		{5, 10},
//	}
//
//	// Степень полинома
//	degree := 2
//
//	// Расчет коэффициентов полинома
//	coefficients := calculateCoefficients(points, degree)
//
//	// Создание графика
//	p := plot.New()
//	p.Title.Text = "Полиномиальная регрессия"
//	p.X.Label.Text = "X"
//	p.Y.Label.Text = "Y"
//
//	// Построение точек данных
//	pts := make(plotter.XYs, len(points))
//	for i := 0; i < len(points); i++ {
//		pts[i].X = points[i].X
//		pts[i].Y = points[i].Y
//	}
//	s, err := plotter.NewScatter(pts)
//	if err != nil {
//		panic(err)
//	}
//	p.Add(s)
//
//	// Построение полиномиальной кривой
//	xValues := make([]float64, 100)
//	yValues := make([]float64, 100)
//	XYS := make([]plotter.XY, 100)
//	for i := 0; i < 100; i++ {
//		xValues[i] = float64(i) / 99 * 5 // Ось X от 0 до 5
//		yValues[i] = polynomial(xValues[i], coefficients)
//		XYS = append(XYS, plotter.XY{X: xValues[i], Y: yValues[i]})
//	}
//	l, err := plotter.NewLine(plotter.XYs(XYS))
//	if err != nil {
//		panic(err)
//	}
//	p.Add(l)
//
//	// Сохранение графика в файл
//	if err := p.Save(4*vg.Inch, 4*vg.Inch, "polynomial_regression.png"); err != nil {
//		panic(err)
//	}
//
//	fmt.Println("Коэффициенты полинома:", coefficients)
//}
