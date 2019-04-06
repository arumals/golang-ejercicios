package dimensiones

const KmPorMilla = 1.60934
const MetrosPorPie = 0.3048
const CentimegrosPorPulgada = 2.54

func MillasEnKm(s float64) float64 {
	return s * KmPorMilla
}

func PiesEnMetros(s float64) float64 {
	return s * MetrosPorPie
}

func PulgadasEnCentimentros(s float64) float64 {
	return s * CentimegrosPorPulgada
}
