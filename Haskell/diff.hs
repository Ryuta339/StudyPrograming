diff_e :: (Double -> Double) -> Double -> Double
diff_e f x = (f (x + h) - f (x - h)) / (2 * h)
	where h = 0.00005

diff :: (Double -> Double) -> Double -> Double
diff f = diff_e f


