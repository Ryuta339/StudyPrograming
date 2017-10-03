diff_e :: (Double -> Double) -> Double -> Double
diff_e f x = (f (x + h) - f (x - h)) / (2 * h)
	where h = 0.00005

diff :: (Double -> Double) -> Double -> Double
diff f = diff_e f

partial_diff_e :: ([Double] -> Double) -> Int -> [Double] -> Double
partial_diff_e f i x = (f (ladd x plh) - f (ladd x mlh)) / (2 * h)
	where h   = 0.00005;
		  plh = (replicate i 0) ++ [h] ++ (replicate (n-i-1) 0);
		  mlh = map (\x -> -x) plh;
		  n   = length x

partial_diff :: ([Double] -> Double) -> Int -> [Double] -> Double
partial_diff f i = partial_diff_e f i

ladd :: (Num a) => [a] -> [a] -> [a]
ladd x y = sum $ map (\(l1,l2) -> l1 + l2) (zip x y)
