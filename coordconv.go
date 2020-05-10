package swedishgrid

import "math"

var axis = 6378137.0
var flattening = 1.0 / 298.257222101
var centralMeridian = 15.806284529
var scale = 1.00000561024
var falseNorthing = -667.711
var falseEasting = 1500064.274
var e2 = flattening * (2.0 - flattening)
var n = flattening / (2.0 - flattening)
var aRoof = axis / (1.0 + n) * (1.0 + n*n/4.0 + n*n*n*n/64.0)
var delta1 = n/2.0 - 2.0*n*n/3.0 + 37.0*n*n*n/96.0 - n*n*n*n/360.0
var delta2 = n*n/48.0 + n*n*n/15.0 - 437.0*n*n*n*n/1440.0
var delta3 = 17.0*n*n*n/480.0 - 37*n*n*n*n/840.0
var delta4 = 4397.0 * n * n * n * n / 161280.0
var aStar = e2 + e2*e2 + e2*e2*e2 + e2*e2*e2*e2
var bStar = -(7.0*e2*e2 + 17.0*e2*e2*e2 + 30.0*e2*e2*e2*e2) / 6.0
var cStar = (224.0*e2*e2*e2 + 889.0*e2*e2*e2*e2) / 120.0
var dStar = -(4279.0 * e2 * e2 * e2 * e2) / 1260.0
var degToRad = math.Pi / 180
var lambdaZero = centralMeridian * degToRad

// RT90toWGS84 Converts coordinates from Swedish Grid RT90 to coordinates in WGS84
func RT90toWGS84(x float64, y float64) (float64, float64) {
	var xi = (x - falseNorthing) / (scale * aRoof)
	var eta = (y - falseEasting) / (scale * aRoof)
	var xiPrim = xi -
		delta1*math.Sin(2.0*xi)*math.Cosh(2.0*eta) -
		delta2*math.Sin(4.0*xi)*math.Cosh(4.0*eta) -
		delta3*math.Sin(6.0*xi)*math.Cosh(6.0*eta) -
		delta4*math.Sin(8.0*xi)*math.Cosh(8.0*eta)
	var etaPrim = eta -
		delta1*math.Cos(2.0*xi)*math.Sinh(2.0*eta) -
		delta2*math.Cos(4.0*xi)*math.Sinh(4.0*eta) -
		delta3*math.Cos(6.0*xi)*math.Sinh(6.0*eta) -
		delta4*math.Cos(8.0*xi)*math.Sinh(8.0*eta)
	var phiStar = math.Asin(math.Sin(xiPrim) / math.Cosh(etaPrim))
	var deltaLambda = math.Atan(math.Sinh(etaPrim) / math.Cos(xiPrim))
	var lngRadian = lambdaZero + deltaLambda
	var latRadian = phiStar + math.Sin(phiStar)*math.Cos(phiStar)*
		(aStar+
			bStar*(math.Pow(math.Sin(phiStar), 2))+
			cStar*(math.Pow(math.Sin(phiStar), 4))+
			dStar*(math.Pow(math.Sin(phiStar), 6)))
	return latRadian * 180.0 / math.Pi, lngRadian * 180.0 / math.Pi
}
